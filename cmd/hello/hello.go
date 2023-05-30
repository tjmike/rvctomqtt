package main

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"runtime"
	"rvctomqtt/handler"
	"rvctomqtt/intf"
	"rvctomqtt/pform"
	"rvctomqtt/pool"
	"rvctomqtt/rvc"
	"rvctomqtt/rvcChangeListener"
	"rvctomqtt/rvcmqtt"

	//"rvctomqtt/rvc"
	"time"
)

func main() {
	print("MAX PROCS=")
	print(runtime.GOMAXPROCS(0))
	print("\n")

	// This channel sends change event  messages (not pointer)
	rvcChangeEvents := make(chan rvc.RvcItemIF, 32)

	// This channel send JSON message to MQTT
	mqttEvents := make(chan *rvcmqtt.MqttMessage, 32)

	// Listen on this to process the raw can message
	fromSocket := make(chan *intf.CanFrameIF, 32)

	snd := rvcmqtt.NewSender("tcp://192.168.50.12:1883", "goplay")
	go snd.AwaitMessages(mqttEvents)

	// seems like we must be explicit with the interface - we can't pass the item
	// that implements it
	//var frameFactoryInterface intf.CanFrameFactory = &can.CanFrameFactory{}
	var frameFactoryInterface intf.CanFrameFactory = &rvc.RVCFrameFactory{}
	var dataFramePool = pool.NewPool(&frameFactoryInterface)

	prodConfig := zap.NewProductionConfig()
	prodConfig.Encoding = "console"
	prodConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	prodConfig.EncoderConfig.EncodeDuration = zapcore.StringDurationEncoder
	log, _ := prodConfig.Build()

	//var log, _ = zap.NewProduction()

	// Assume this context is only used in a single goroutine at a time
	logFields := make([]zap.Field, 1)
	logFields[0] = zap.Field{
		Key:       "workID",
		Type:      zapcore.StringType,
		Integer:   0,
		String:    "RVCReader01",
		Interface: nil,
	}
	var ctx1 context.Context = context.Background()
	ctx1 = context.WithValue(ctx1, "logFields", logFields)

	go pform.GetRVCMessages(&ctx1, log, dataFramePool, fromSocket)

	logFields2 := make([]zap.Field, 1)
	logFields2[0] = zap.Field{
		Key:       "workID",
		Type:      zapcore.StringType,
		Integer:   0,
		String:    "RVCHandler01",
		Interface: nil,
	}
	var ctx2 context.Context = context.Background()
	ctx2 = context.WithValue(ctx2, "logFields", logFields2)

	go handler.RVCMessageHandler(&ctx2, log, fromSocket, dataFramePool, rvcChangeEvents)

	go rvcChangeListener.Listen(rvcChangeEvents, mqttEvents)

	for {
		fmt.Printf("Sleep # goRoutines = %d\n", runtime.NumGoroutine())

		time.Sleep(time.Second * 30)
	}

}
