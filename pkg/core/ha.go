package core

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"sync"
	"time"
)

type HomeAssistant struct {
	conn *websocket.Conn
	events chan byte
	done chan struct{}
}

var once sync.Once
var ha HomeAssistant

func connect() (*websocket.Conn, error) {
	host := os.Getenv("HA_HOST")
	u := url.URL{Scheme: "wss", Host: host, Path: "/api/websocket"}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func authenticate(c *websocket.Conn) error {
	_, message, err := c.ReadMessage()
	if err != nil {
		return err
	}

	if strings.Contains(string(message), "auth_required") {
		token := os.Getenv("HA_TOKEN")
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

func(ha *HomeAssistant) getEventStream() chan byte {
	stream := make(chan byte, 1000)
	go func() {
		defer close(ha.done)
		for {
			_, message, err := ha.conn.ReadMessage()
			if err != nil {
				fmt.Println("read:", err)
				return
			}
			fmt.Printf("recv: %s\n", message)
		}
	}()
	return stream
}

func (ha *HomeAssistant) CloseConnection() error {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	for {
		select {
		case <-ha.done:
			fmt.Println("done")
			return nil
		case <-interrupt:
			fmt.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := ha.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
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
		connection, err := connect()
		if err != nil {
			log.Fatalln(err)
		}

		if err := authenticate(connection); err != nil {
			log.Fatalln(err)
		}

		if err := subscribeToEvents(connection); err != nil {
			log.Fatalln(err)
		}

		ha = HomeAssistant{
			conn: connection,
			done: make(chan struct{}),
		}

		ha.events = ha.getEventStream()
	})

	return ha
}

type Target struct {
	EntityId string `json:"entity_id"`
}

type Message struct {
	Id int `json:"id"`
	Type string `json:"type"`
	Domain string `json:"domain"`
	Service string `json:"service"`
	ServiceData map[string]string `json:"service_data"`
	Target Target `json:"target"`
}

func (ha HomeAssistant) SendMessage(t, domain, service, colorName, brightness, entityId string) error {
	serviceData := map[string]string{}
	if colorName != "" {
		serviceData["color_name"] = colorName
	}
	if brightness != "" {
		serviceData["brightness"] = brightness
	}

	m := Message{
		Id:          2,
		Type:        t,
		Domain:		 domain,
		Service:     service,
		ServiceData: serviceData,
		Target:      Target{
			EntityId: entityId,
		},
	}

	mBytes, err := json.Marshal(m)
	if err != nil {
		return err
	}

	if err := ha.conn.WriteMessage(1, mBytes); err != nil {
		return err
	}

	return nil
}
