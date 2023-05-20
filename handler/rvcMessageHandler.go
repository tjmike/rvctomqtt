package handler

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"reflect"
	"rvctomqtt/intf"
	"rvctomqtt/pool"
	"rvctomqtt/rvc"
	"rvctomqtt/utils"
	"strings"
	"time"
)

// RVCMessageHandler - we expect to see RvcFrames here. For now, we leave the channel interface as can frames and just
// ignore data if it's not an RVC frame.
func RVCMessageHandler(ctx *context.Context, log *zap.Logger, fromSocket chan *intf.CanFrameIF, pool *pool.Pool) {

	// interesting issue - when
	log = utils.ApplyContext(ctx, log)

	log.Info("############################### HANDLER #####################")
	var nmsg uint32 = 0

	//var m := treemap.NewWithIntComparator()

	// # of times the key (dgn) has been seen
	var seen = make(map[uint32]uint64)

	var goodPackets uint64 = 0
	var badpackets uint64 = 0
	for {
		data := <-fromSocket //get the RVC packet

		rvcFrame, ok := (*data).(*rvc.RvcFrame) // all frames **should** be RvcFrams
		if ok {
			goodPackets++
			var dgn uint32 = rvcFrame.DGN()

			//logRawFrame(log, rvcFrame, dgn)

			// Get existing or create a new RVCItem
			// Note that ok will be false if we don't recognize the DGN because we then won't be able to create a DGN
			// specific instance.
			var rvcItemPtr, ok = rvc.GetRVCItem(rvcFrame)

			if ok {
				// We MUST do this first - if Init isn't called then all bets are off
				// set the state of the rvc item from the data frame
				var rvcItem = *rvcItemPtr
				rvcItem.Init(rvcFrame)

				// if timestamps are equal then it must have changed or is new
				if rvcItem.GetTimestamp() == rvcItem.GetLastChanged() {
					if log.Level() >= zapcore.InfoLevel {
						// just to prevent noise....
						if !rvc.Ignore(rvcItemPtr) {
							log.Info(fmt.Sprintf("CHANGED: %s", rvcItem))
						}
					}
					dumpItemViaReflection(rvcItemPtr, dgn)
				} else {
					if log.Level() >= zapcore.DebugLevel {
						log.Debug(fmt.Sprintf("no change detected  = %x existing = %s", rvcFrame.MessageBytes, rvcItem))
					}
				}
			} else {
				if log.Level() >= zapcore.InfoLevel {
					log.Info(fmt.Sprintf("RVCMessageHandler: Could not create frame for dgn: %x", rvcFrame.DGN()))
				}
			}
			nseen := 1 + seen[dgn]
			seen[dgn] = nseen
		} else {
			badpackets++
			// If we get here it's likely a bug
			log.Warn("NOT RVC FRAME???")
		}

		if (goodPackets % 10000) == 0 {
			logStats(log, seen, goodPackets, badpackets)
		}

		// We're don with this message, put it back into the pool
		(*pool).ReturnToPool(data)

		nmsg++
	}
}

func logStats(log *zap.Logger, seen map[uint32]uint64, good uint64, bad uint64) {
	if log.Level() >= zapcore.InfoLevel {
		log.Info(fmt.Sprintf("############# STATS ######## #goodParckets=%d #badPackets=%d", good, bad))
		for k, v := range seen {
			var name = rvc.DGNName(k)
			if log.Level() >= zapcore.InfoLevel {
				log.Info(fmt.Sprintf("%10s %10x %10d  %20s", "STATS", k, v, name))
			}
		}
	}
}

func logRawFrame(log *zap.Logger, rvcFrame *rvc.RvcFrame, dgn uint32) {
	if log.Level() >= zapcore.InfoLevel {
		log.Info(fmt.Sprintf("RAW FRAME: %s dgn: %x sa: %x , details: %s raw: %x",
			rvcFrame.GetTimeStamp().Format("01-02-2006 15:04:05.000000"), dgn, rvcFrame.GetSourceAddress(),
			rvcFrame,
			rvcFrame.MessageBytes))
	}
}

func dumpItemViaReflection(rvcItem *rvc.RvcItemIF, dgn uint32) {

	// let's see what we have
	var reflectedType = reflect.TypeOf(*rvcItem)
	var reflectedValue = reflect.ValueOf(*rvcItem)

	// special case - all should have a getName method. We want to establish that the
	// method exists with the type method by name so we can later use the value methodByName
	// TODO make string a constant
	var _, ok = reflectedType.MethodByName("GetName")
	if ok {

		inputs := make([]reflect.Value, 0)
		var getNameResults = reflectedValue.MethodByName("GetName").Call(inputs)
		var dgnName = getNameResults[0].Interface()
		var instanceName interface{} = "N/A"
		// We assume if the item has GetName that it also has GetInstanceName
		var mbn = reflectedValue.MethodByName("GetInstanceName")
		//fmt.Printf("\tMBN: %s\n", mbn.Kind())
		if mbn.Kind() != reflect.Invalid {
			instanceName = mbn.Call(inputs)[0].Interface()
			//instanceName = reflectedValue.MethodByName("GetInstanceName").Call(inputs)[0].Interface()
		}
		fmt.Printf("\t%x %s(%s)\n", dgn, dgnName, instanceName)
		var nmethods = reflectedType.NumMethod()
		for i := 0; i < nmethods; i++ {
			var xmethod = reflectedType.Method(i)
			var xmtype = xmethod.Type
			var xmname = xmethod.Name
			var nout = xmtype.NumOut()
			var nin = xmtype.NumIn()
			if nout == 1 && nin == 1 {
				if strings.HasPrefix(xmname, "Get") {
					var methodOutputDataType = xmtype.Out(0).Name()

					if methodOutputDataType == "uint8" || methodOutputDataType == "Uint2" || methodOutputDataType == "uint16" || methodOutputDataType == "uint32" {
						var XXX = reflectedValue.Method(i).Call(inputs)
						var yyy = XXX[0].Uint()
						fmt.Printf("\t\t%s(%s)=%d\n", xmname, methodOutputDataType, yyy)

					} else if methodOutputDataType == "Time" {
						inputs := make([]reflect.Value, 0)
						var XXX = reflectedValue.Method(i).Call(inputs)
						var yyy = XXX[0]
						var zzz = yyy.Interface()
						var zzz2 = zzz.(time.Time)
						fmt.Printf("\t\t%s(%s)=%s\n", xmname, methodOutputDataType, zzz2.Format("01-02-2006 15:04:05.000000"))

					} else if methodOutputDataType == "string" {
						inputs := make([]reflect.Value, 0)
						var XXX = reflectedValue.Method(i).Call(inputs)
						var yyy = XXX[0]
						fmt.Printf("\t\t%s(%s)=%s\n", xmname, methodOutputDataType, yyy)
					} else if methodOutputDataType == "float64" {
						inputs := make([]reflect.Value, 0)
						var XXX = reflectedValue.Method(i).Call(inputs)
						var yyy = XXX[0].Float()
						fmt.Printf("\t\t%s(%s)=%f\n", xmname, methodOutputDataType, yyy)
					} else {
						fmt.Printf("\t\t(UNSUPPORTED TYPE)   name = %s type = %s \n", xmname, methodOutputDataType)
					}
				}
			}
		}

	}

}
