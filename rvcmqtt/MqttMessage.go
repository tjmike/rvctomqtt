package rvcmqtt

import "fmt"

type MqttMessage struct {
	Topic   string // Topic
	Payload string // JSON payload
}

func (m *MqttMessage) String() string {
	return fmt.Sprint("Topic: %s Payload %s", m.Topic, m.Payload)
}
