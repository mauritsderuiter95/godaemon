package core

import "time"

type Event struct {
	Id    int    `json:"id"`
	Type  string `json:"type"`
	Event struct {
		EventType string `json:"event_type"`
		Data      struct {
			EntityId string `json:"entity_id"`
			OldState struct {
				EntityId   string `json:"entity_id"`
				State      string `json:"state"`
				Attributes struct {
					MinMireds           int       `json:"min_mireds"`
					MaxMireds           int       `json:"max_mireds"`
					SupportedColorModes []string  `json:"supported_color_modes"`
					ColorMode           string    `json:"color_mode"`
					Brightness          int       `json:"brightness"`
					ColorTemp           int       `json:"color_temp"`
					HsColor             []float64 `json:"hs_color"`
					RgbColor            []int     `json:"rgb_color"`
					XyColor             []float64 `json:"xy_color"`
					IsDeconzGroup       bool      `json:"is_deconz_group"`
					AllOn               bool      `json:"all_on"`
					FriendlyName        string    `json:"friendly_name"`
					SupportedFeatures   int       `json:"supported_features"`
				} `json:"attributes"`
				LastChanged time.Time `json:"last_changed"`
				LastUpdated time.Time `json:"last_updated"`
				Context     struct {
					Id       string      `json:"id"`
					ParentId interface{} `json:"parent_id"`
					UserId   string      `json:"user_id"`
				} `json:"context"`
			} `json:"old_state"`
			NewState struct {
				EntityId   string `json:"entity_id"`
				State      string `json:"state"`
				Attributes struct {
					MinMireds           int       `json:"min_mireds"`
					MaxMireds           int       `json:"max_mireds"`
					SupportedColorModes []string  `json:"supported_color_modes"`
					ColorMode           string    `json:"color_mode"`
					Brightness          int       `json:"brightness"`
					ColorTemp           int       `json:"color_temp"`
					HsColor             []float64 `json:"hs_color"`
					RgbColor            []int     `json:"rgb_color"`
					XyColor             []float64 `json:"xy_color"`
					IsDeconzGroup       bool      `json:"is_deconz_group"`
					AllOn               bool      `json:"all_on"`
					FriendlyName        string    `json:"friendly_name"`
					SupportedFeatures   int       `json:"supported_features"`
				} `json:"attributes"`
				LastChanged time.Time `json:"last_changed"`
				LastUpdated time.Time `json:"last_updated"`
				Context     struct {
					Id       string      `json:"id"`
					ParentId interface{} `json:"parent_id"`
					UserId   string      `json:"user_id"`
				} `json:"context"`
			} `json:"new_state"`
			Domain      string `json:"domain"`
			Service     string `json:"service"`
			ServiceData struct {
				EntityId interface{} `json:"entity_id"`
			} `json:"service_data"`
		} `json:"data"`
		Origin    string    `json:"origin"`
		TimeFired time.Time `json:"time_fired"`
		Context   struct {
			Id       string      `json:"id"`
			ParentId interface{} `json:"parent_id"`
			UserId   string      `json:"user_id"`
		} `json:"context"`
	} `json:"event"`
}
