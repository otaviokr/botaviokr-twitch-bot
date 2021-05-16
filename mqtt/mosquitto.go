package mqtt

import (
	"fmt"

	mq "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
)

type MqttClient struct {
	clientId string
	Client mq.Client
}

func NewClient(clientId, broker string, port int) (*MqttClient, error) {
	if len(clientId) < 1 || len(broker) < 1 {
		log.WithFields(
			log.Fields{
				"clientId": clientId,
				"broker": broker,
			}).Error("property 'client ID' and/or 'broker host' not defined. No MQTT defined")

		return nil, fmt.Errorf("clientId and/or broker not defined")
	}

	// Connecting to MQTT.
	opts := mq.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID(clientId) // ClientID can be anything, just make sure it is unique
	// opts.SetUsername("username") // Not needed because we are using the public HiveMQ
	// opts.SetPassword("password") // Not needed because we are using the public HiveMQ
	// opts.SetDefaultPublishHandler(getMessageReceivedHandler(device))
	// opts.OnConnect = getOnConnectHandler()
	// opts.OnConnectionLost = getConnectionLostHandler()

	client := mq.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	return &MqttClient {
		clientId: clientId,
		Client:   client,
	}, nil
}

// Publish will send the message *text* to topic *topic*.
func (m *MqttClient) Publish(topic, text string) (error) {
	if !m.Client.IsConnected() {
		if token := m.Client.Connect(); token.Wait() && token.Error() != nil {
			return token.Error()
		}
	}
	if token := m.Client.Publish(topic, 0, false, text); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}
