package rvcChangeListener

import (
	"fmt"
	"rvctomqtt/handler"
	"rvctomqtt/rvc"
	"rvctomqtt/rvcmqtt"
	"strings"
)

func Listen(evt chan rvc.RvcItemIF, jsonEvents chan *rvcmqtt.MqttMessage) {

	var topicRoot = "mycoach/events"

	var sb strings.Builder

	// forever
	for {
		sb.Reset()
		var item = <-evt
		var dgnName = item.GetName()
		sb.WriteString(topicRoot)
		sb.WriteString("/")
		sb.WriteString(dgnName)

		// There may or may not be an instance name
		// for now we assume if there's an instance value that's > 0 and <= 250 that this dgn supports
		// instances and we use the name if found and otherwise use the instance number
		var instance = item.GetInstance()
		if instance < 254 {

			tmp, found := rvc.GetInstanceName(rvc.DGNInstanceKey{
				DGN:      item.GetDGN(),
				Instance: instance,
			})
			sb.WriteString("/")
			if found {
				sb.WriteString(tmp)
			} else {

				// for address claimed the "instance" is really the source address
				// TODO : can we eliminate or centralize these special cases?????
				if (item.GetDGN() & 0xFF00) == rvc.DGN_ADDRESS_CLAIMED {
					instance = item.GetSourceAddress()
				}

				sb.WriteString(fmt.Sprintf("%d", instance))
			}
		} else {
			if (item.GetDGN() & 0xFF00) == rvc.DGN_ADDRESS_CLAIMED {
				instance = item.GetSourceAddress()
				sb.WriteString("/")
				sb.WriteString(fmt.Sprintf("SA-%d", instance))
			} else if item.GetDGN() == rvc.DGN_INITIAL_PACKET {
				instance = item.GetSourceAddress()
				sb.WriteString("/")
				sb.WriteString(fmt.Sprintf("SA-%d", instance))
				ip, ok := item.(*rvc.InitialPacket)
				if ok {
					sb.WriteString("/")
					sb.WriteString(fmt.Sprintf("RDGN-%x", ip.GetRequestedDGN()))
				} else {
					sb.WriteString("/NotInitialPacket")
				}

			} else if item.GetDGN() == rvc.DGN_DATA_PACKET {
				instance = item.GetSourceAddress()
				sb.WriteString("/")
				sb.WriteString(fmt.Sprintf("SA-%d", instance))
				dp, ok := item.(*rvc.DataPacket)
				if ok {
					sb.WriteString("/")
					sb.WriteString(fmt.Sprintf("PKTN-%d", dp.GetPacketNumber()))
				} else {
					sb.WriteString("/NotDataPacket")

				}
			}

		}

		//
		// package this up and send it to
		//fmt.Printf("EVENT: %s\n", item)
		fmt.Printf("TOPIC: %s JSON: %s\n", sb.String(), handler.DumpItemViaReflection(&item))

		// Send to MQTT
		jsonEvents <- &rvcmqtt.MqttMessage{sb.String(), handler.DumpItemViaReflection(&item)}
		// chan *rvcmqtt.MqttMessage
		// Send MQTT Packet
	}

}
