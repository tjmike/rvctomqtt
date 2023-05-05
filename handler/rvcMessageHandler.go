package handler

import (
	"fmt"
	"reflect"
	"rvctomqtt/intf"
	"rvctomqtt/pool"
	"rvctomqtt/rvc"
	"strings"
	"time"
	//"strings"
)

// RVCMessageHandler - we expect to see RvcFrames here. For now we leave the channel interface as can frames and just
// ignore data if it's not an RVC frame.
func RVCMessageHandler(fromSocket chan *intf.CanFrameIF, pool *pool.Pool) {
	fmt.Printf("############################### HANDLER #####################\n")
	var nmsg uint32 = 0

	//var m := treemap.NewWithIntComparator()

	var seen = make(map[uint32]uint64)

	var packets uint64 = 0
	for {
		data := <-fromSocket

		packets++
		rvcFrame, ok := (*data).(*rvc.RvcFrame)
		if ok {
			var dgn uint32 = uint32(rvcFrame.DGNHigh()) << 8
			dgn = dgn | uint32(rvcFrame.DGNLow())

			fmt.Printf("RAW FRAME: %s %x\n", rvcFrame.GetTimeStamp().Format("01-02-2006 15:04:05.000000"), rvcFrame.MessageBytes)

			// Get existing or create an ew RVCItem
			var rvcItem, ok = rvc.GetRVCItem(rvcFrame)

			if ok {
				// We MUST do this first - if Init isn't called then all bets are off
				// set the state of the tvc item from the data frame
				var rvcFrameDereferenced = *rvcItem
				rvcFrameDereferenced.Init(rvcFrame)

				var ts = rvcFrameDereferenced.GetTimestamp()
				var lastChanged = rvcFrameDereferenced.GetLastChanged()
				// if timestamps are equal then it must have changed
				if ts == lastChanged {
					fmt.Printf("%s\n", rvcFrameDereferenced)

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
				} else {
					fmt.Printf("XXXXX could not create  = %x\n", rvcFrame.MessageBytes)

				}
			} else {
				fmt.Printf("RVCMessageHandler: Could not create frame for dgn: %x\n", rvcFrame.DGN())
			}
			nseen := 1 + seen[dgn]
			seen[dgn] = nseen
		} else {
			fmt.Printf("NOT RVC FRAME???")

		}

		if (packets % 100) == 0 {
			fmt.Printf("############# STATS ########\n")
			for k, v := range seen {
				var name = rvc.DGNName(k)
				fmt.Printf("%10s %10x %10d  %20s\n", "STATS", k, v, name)
			}
		}

		// We're don with this message, put it back into the pool
		(*pool).ReturnToPool(data)

		nmsg++
	}
}
