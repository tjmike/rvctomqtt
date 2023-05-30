package rvcmqtt

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type Mqttsender struct {
	server   string
	clientID string
	client   MQTT.Client
}

// AwaitMessages - wait for MQqttMesages and fire them upon receipt.
func (s *Mqttsender) AwaitMessages(msg chan *MqttMessage) {
	for {
		s.send(*<-msg)
	}

}

// send - Publish the supplied message to the mqtt server
func (s *Mqttsender) send(msg MqttMessage) {
	// NOTE: Seems like retained=true causes messages to not make it to the server
	s.client.Publish(msg.Topic, 0, false, msg.Payload)
}
func NewSender(broker string, clientID string) Mqttsender {

	opts := MQTT.NewClientOptions().AddBroker(broker)
	opts.SetClientID(clientID)
	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	return Mqttsender{
		server:   broker,
		clientID: clientID,
		client:   c,
	}
}

/*
func main() {
	snd := NewSender("tcp://192.168.50.12:1883", "goplay")
	var msg = MqttMessage{
		Topic:   "foo/bar",
		Payload: "{ \"Test\": \"testingaaa\" }",
	}
	snd.send(msg)
	fmt.Printf("DONE")

	msgChan := make(chan *MqttMessage, 32)
	go snd.AwaitMessages(msgChan)

	for i := 0; i < 10; i++ {
		var msg = MqttMessage{
			Topic:   "foo/bar",
			Payload: fmt.Sprintf("{ \"Test\": \"testing-%d\"  }", i),
		}
		msgChan <- &msg
		time.Sleep(time.Second * 5)
	}
}
*/
