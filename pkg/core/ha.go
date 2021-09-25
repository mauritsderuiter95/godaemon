package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"time"
)

type HomeAssistant struct {
	host      string
	token     string
	sToken    string
	wsURL     string
	apiURL    string
	wsconn    *websocket.Conn
	events    chan HaEvent
	done      chan struct{}
	Callbacks map[string][]func(Event)
}

type HaEvent struct {
	Id    int    `json:"id"`
	Type  string `json:"type"`
	Event Event  `json:"event"`
}

var once sync.Once
var ha HomeAssistant

func connect(wsURL string) (*websocket.Conn, error) {
	c, res, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		fmt.Println(res)
		return nil, err
	}
	return c, nil
}

func authenticate(c *websocket.Conn, token string) error {
	_, message, err := c.ReadMessage()
	if err != nil {
		return err
	}

	if strings.Contains(string(message), "auth_required") {
		err := c.WriteMessage(1, []byte(fmt.Sprintf("{\"type\":\"auth\",\"access_token\":\"%s\"}", token)))
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("unexpected response: %s", message)
	}

	return nil
}

func subscribeToEvents(c *websocket.Conn) error {
	return c.WriteMessage(1, []byte("{\"id\":1,\"type\":\"subscribe_events\"}"))
}

func (ha *HomeAssistant) getEventStream() chan HaEvent {
	stream := make(chan HaEvent, 1000)
	go func() {
		defer close(ha.done)
		for {
			_, message, err := ha.wsconn.ReadMessage()
			if err != nil {
				fmt.Println("read:", err)
				return
			}

			e := HaEvent{}
			if err := json.Unmarshal(message, &e); err != nil {
				fmt.Println("parsing error:", err)
				return
			}

			stream <- e
			//fmt.Printf("%v\n", e)
		}
	}()
	return stream
}

func (ha *HomeAssistant) CloseConnection() error {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	signal.Notify(interrupt, os.Kill)

	for {
		select {
		case <-ha.done:
			fmt.Println("rebuild")

			err := ha.wsconn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				return err
			}
			time.Sleep(time.Second)
			return nil
		case <-interrupt:
			fmt.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := ha.wsconn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				return err
			}
			select {
			case <-ha.done:
			case <-time.After(time.Second):
			}
			return nil
		}
	}
}

func GetInstance() HomeAssistant {
	once.Do(func() {
		host := os.Getenv("HA_HOST")
		token := os.Getenv("HA_TOKEN")
		sToken := os.Getenv("SUPERVISOR_TOKEN")

		wsURL := fmt.Sprintf("wss://%s/api/websocket", host)
		apiURL := fmt.Sprintf("https://%s/api", host)

		if sToken != "" {
			host = "supervisor"
			token = sToken

			wsURL = fmt.Sprintf("ws://supervisor/core/websocket")
			apiURL = "http://supervisor/core/api"
		}
		connection, err := connect(wsURL)
		if err != nil {
			log.Fatalln(err)
		}
		if err := authenticate(connection, token); err != nil {
			log.Fatalln(err)
		}

		if err := subscribeToEvents(connection); err != nil {
			log.Fatalln(err)
		}

		ha = HomeAssistant{
			host:      host,
			token:     token,
			sToken:    sToken,
			wsURL:     wsURL,
			apiURL:    apiURL,
			wsconn:    connection,
			done:      make(chan struct{}),
			Callbacks: map[string][]func(Event){},
		}

		ha.events = ha.getEventStream()
	})

	return ha
}

func (ha *HomeAssistant) HandleEvents() {
	for m := range ha.events {
		id := m.Event.Data.EntityId
		if id == "" {
			id = m.Event.Data.Id
		}
		if val, ok := ha.Callbacks[id]; ok {
			for _, f := range val {
				go f(m.Event)
			}
		}
	}
}

func (ha HomeAssistant) GetState(entity string) State {
	client := http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/states/%s", ha.apiURL, entity), nil)
	if err != nil {
		fmt.Println(err)
		return State{}
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", ha.token))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return State{}
	}

	body, err := io.ReadAll(res.Body)

	if res.Status != "200 OK" {
		if err != nil {
			fmt.Println(err)
			return State{}
		}
	}

	s := State{}

	if err := json.Unmarshal(body, &s); err != nil {
		fmt.Println(err)
		return State{}
	}

	return s
}

func (ha HomeAssistant) CallService(domain, service, entityId string, attrs map[string]string) error {
	if attrs == nil {
		attrs = map[string]string{}
	}
	attrs["entity_id"] = entityId

	mBytes, err := json.Marshal(attrs)
	if err != nil {
		return err
	}

	client := http.Client{}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/services/%s/%s", ha.apiURL, domain, service), bytes.NewBuffer(mBytes))
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", ha.token))

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	if res.Status != "200 OK" {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		fmt.Println(string(body))
	}

	return nil
}
