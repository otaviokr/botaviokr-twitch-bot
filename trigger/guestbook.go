package trigger

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/otaviokr/botaviokr-twitch-bot/mqtt"
	hbot "github.com/otaviokr/hellivabot"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Guestbook struct {
	Joined int64 `json:"joined"`
	User string `json:"user"`
}

// Guestbook keeps an eye on who is in the chat.
func GuestBook(mqttClient *mqtt.MqttClient) hbot.Trigger {
	return hbot.Trigger{
		Condition: func (b *hbot.Bot, m *hbot.Message) bool {
			return strings.EqualFold(m.Command, "JOIN")
		},
		Action: func (b *hbot.Bot, m *hbot.Message) bool {
			topic := viper.GetString("triggers.guestbook.topic")
			timestamp := time.Now().Unix()
			if mqttClient == nil {
				log.WithFields(
					log.Fields{
						"username": m.Content,
						"command": m.Command,
						"timestamp": timestamp,
						"action": "GuestBook",
					}).Errorf("no MQTT client defined, event '%s joined' will not be published", m.Content)
				return false
			}

			log.WithFields(
				log.Fields{
					"username": m.Content,
					"command": m.Command,
					"timestamp": timestamp,
				}).Infof("user joined: %s", m.Content)

				guest := Guestbook {
					Joined: timestamp,
					User: m.Content,
				}

				payload, err := json.Marshal(guest)
				if err != nil {
					log.WithFields(
						log.Fields{
							"username": m.Content,
							"command": m.Command,
							"timestamp": timestamp,
						}).Error("failed to parsed payload to MQTT")
						return false
				}

				err = mqttClient.Publish(topic, string(payload))
				if err != nil {
					log.WithFields(
						log.Fields{
							"trigger": "Guestbook",
							"topic": topic,
							"username": m.Content,
						}).Error("failed to publish to MQTT")
				}

			return false
		},
	}
}