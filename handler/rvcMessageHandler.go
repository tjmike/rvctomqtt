package handler

import (
	"fmt"
	"reflect"
	"rvctomqtt/intf"
	"rvctomqtt/rvc"
	"strings"
	"time"
	//"strings"
)

func CanMessageHandler(fromSocket, toSocket chan *intf.CanFrameIF) {
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

			fmt.Printf("RAW FRAME: %x\n", rvcFrame.MessageBytes)

			var rvcItem, ok = rvc.GetRVCItem(rvcFrame)

			if ok {
				// We MUST do this first - if Init isn't called then all bets are off
				var rvcFrameDereferenced = *rvcItem
				rvcFrameDereferenced.Init(rvcFrame)

				var ts = rvcFrameDereferenced.GetTimestamp()
				var lastChanged = rvcFrameDereferenced.GetLastChanged()
				// if timestamps are equal then it must have changed
				if ts == lastChanged {
					fmt.Printf("DCSS1 VALUE = %s\n", rvcFrameDereferenced)

					// let's see what we have
					var xtype = reflect.TypeOf(*rvcItem)
					//var xval = reflect.ValueOf(rvcItem)
					var xval2 = reflect.ValueOf(*rvcItem)
					//var xkind = xtype.Kind()

					//fmt.Printf("reflected type =  %s kind = %s val %s\n", xtype, xkind, xval)
					var _, ok = xtype.MethodByName("GetName")
					if ok {
						inputs := make([]reflect.Value, 0)

						var dgnNameTmp1 = xval2.MethodByName("GetName").Call(inputs)
						var dgnName = dgnNameTmp1[0].Interface()

						var instanceName interface{} = "N/A"
						var mbn = xval2.MethodByName("GetInstanceName")

						//fmt.Printf("\tMBN: %s\n", mbn.Kind())
						if mbn.Kind() != reflect.Invalid {
							instanceName = xval2.MethodByName("GetInstanceName").Call(inputs)[0].Interface()
						}
						fmt.Printf("\t%x %s(%s)\n", dgn, dgnName, instanceName)

						var nmethods = xtype.NumMethod()
						if ok {
							//switch dgnName.Kind() {
							//case reflect.String:
							//	fmt.Printf("ZZZ IS REFLECT STRING%s\n", dgnName.String())
							//}
							//fmt.Printf("*************** METHOD: **GetName** dgn; %x name= %s\n", dgn, dgnName)
							//fmt.Printf("ZZZ %s\n", dgnName.String())
							//fmt.Printf("ZZZ %s\n", dgnName.Kind())
							//fmt.Printf("ZZZ %s\n", dgnName.Type())
						}

						for i := 0; i < nmethods; i++ {
							var xmethod = xtype.Method(i)
							var xmtype = xmethod.Type
							var xmname = xmethod.Name
							var nout = xmtype.NumOut()
							var nin = xmtype.NumIn()

							if nout == 1 && nin == 1 {
								if strings.HasPrefix(xmname, "Get") {
									var methodOutputDataType = xmtype.Out(0).Name()

									if methodOutputDataType == "uint8" || methodOutputDataType == "uint2" || methodOutputDataType == "uint16" || methodOutputDataType == "uint32" {

										//fmt.Printf("xval2 = %s\n", xval2)
										//fmt.Printf("xval2  nmethods= %d\n", xval2.NumMethod())
										//fmt.Printf("xval2  m1= %s\n", xval2.Method(i))

										var XXX = xval2.Method(i).Call(inputs)
										var yyy = XXX[0]
										//fmt.Printf("xval2  RET= %x\n", XXX)
										//fmt.Printf("xval2  RET= %x\n", yyy)

										//fmt.Printf("METHOD: dgn: %x  %s/%s = %d type: %s\n", dgn, dgnName, xmname, yyy, methodOutputDataType)

										//mi := xval.Method(i)
										//fmt.Printf("NIN = %s\n", mi)

										//     reflect.ValueOf(&t).MethodByName("Foo").Call([]reflect.Value{})
										//     reflect.ValueOf(&t).MethodByName("Geeks").Call([]reflect.Value{})
										//mbn := xval.MethodByName(xmname)

										//fmt.Printf("NIN = %s\n", xval.Method(i).Call(inputs))
										//fmt.Printf("EXP = %s\n", xmethod.IsExported())
										//fmt.Printf("FUNC = %s\n", xmethod.Func.Call(inputs))
										//reflecVal := mbn.Call(inputs)
										//refVal := mbn.Call(nil)
										// refVal := mbn.Call([]reflect.Value{})

										//fmt.Printf("\tGETTER: dgn: %x method: %s/%s name = %s ret = %s val = %d\n", dgn, dgnName, xmethod, xmname, methodOutputDataType, yyy)
										//fmt.Printf("\tGETTER: dgn: %x method: %s/%s ret = %s val = %d\n", dgn, dgnName, xmname, methodOutputDataType, yyy)
										fmt.Printf("\t\t%s(%s)=%d\n", xmname, methodOutputDataType, yyy)

									} else if methodOutputDataType == "Time" {
										inputs := make([]reflect.Value, 0)
										var XXX = xval2.Method(i).Call(inputs)
										var yyy = XXX[0]
										var zzz = yyy.Interface()
										var zzz2 = zzz.(time.Time)

										//fmt.Printf("\t\t%s(%s)=%s %s\n", xmname, methodOutputDataType, yyy.Type(), zzz2.Format("01-02-2006 15:04:05.000000"))

										//fmt.Printf("\tDGN: %x  method: %s/%s name = %s ret = %s val = %s\n", dgn, dgnName, xmethod, xmname, methodOutputDataType, yyy)
										fmt.Printf("\t\t%s(%s)=%s\n", xmname, methodOutputDataType, zzz2.Format("01-02-2006 15:04:05.000000"))

									} else if methodOutputDataType == "string" {
										inputs := make([]reflect.Value, 0)
										var XXX = xval2.Method(i).Call(inputs)
										var yyy = XXX[0]

										//fmt.Printf("\tGETTER: dgn: %x method: %s name = %s/%s ret = %s val(STR) = %s\n", dgn, xmethod, dgnName, xmname, methodOutputDataType, yyy)
										//fmt.Printf("\tGETTER: dgn: %d name = %s/%s ret = %s val(STR) = %s\n", dgn, dgnName, xmname, methodOutputDataType, yyy)
										fmt.Printf("\t\t%s(%s)=%s\n", xmname, methodOutputDataType, yyy)
									} else if methodOutputDataType == "float64" {
										inputs := make([]reflect.Value, 0)
										var XXX = xval2.Method(i).Call(inputs)
										var yyy = XXX[0]

										//fmt.Printf("\tGETTER: dgn: %x  method: %s name = %s/%s ret = %s val = %f\n", dgn, xmethod, dgnName, xmname, methodOutputDataType, yyy)
										//fmt.Printf("\tGETTER: dgn: %x  name = %s/%s ret = %s val = %f\n", dgn, dgnName, xmname, methodOutputDataType, yyy)
										fmt.Printf("\t\t%s(%s)=%f\n", xmname, methodOutputDataType, yyy)

									} else {
										//fmt.Printf("\tGETTER: UNSUPPORTED TYPE: dgn: %x  method: %s name = %s/%s ret = %s \n", dgn, xmethod, dgnName, xmname, methodOutputDataType)
										fmt.Printf("\t\t(UNSUPPORTED TYPE)   name = %s = %s \n", xmname, methodOutputDataType)
										//fmt.Printf("METHOD: dgn: %x UNSUPPORTED TYPE: method = %s/%s %s\n", dgn, dgnName, xmname, methodOutputDataType)
									}

								}
							}

							//for j := 0; j < nout; j++ {
							//	var nn = xmtype.Out(j)
							//
							//	var kk = nn.Kind()
							//	var nnn = nn.Name()
							//
							//	if strings.HasPrefix("Get", "Get") {
							//
							//		fmt.Printf("\t\t out[%d]=%s kind = %s name = %s\n", j, nn, kk, nnn)
							//	}
							//
							//}

						}

						//fmt.Printf("reflected type =  %s kind = %s val %s\n", xtype, xkind, xval)
					}
					//var nFields = t.NumField()
					//for i;=0 i<nFields;i++ {
					//	t.
					//}

					//
					//if (dgn == rvc.DGN_DC_SOURCE_STATUS_1) || (dgn == rvc.DGN_DC_SOURCE_STATUS_1_SPYDER) || (dgn == rvc.DGN_DC_DIMMER_STATUS_3) {
					//
					//	var rvcFrameDereferenced rvc.RvcItemIF
					//
					//	if (dgn == rvc.DGN_DC_SOURCE_STATUS_1) || (dgn == rvc.DGN_DC_SOURCE_STATUS_1_SPYDER) {
					//		rvcFrameDereferenced = &rvc.DCSourceStatus1{}
					//	} else {
					//		rvcFrameDereferenced = &rvc.DCDimmerStatus3{}
					//	}

					/*
						var flds = rvcFrameDereferenced.Fields()
						var len = len((*flds))
						for x := 0; x < len; x++ {
							var f = (*flds)[x]
							var tt = f.Gettype()
							var nn = f.GetName()
							var nnS = string(nn)
							var zz string
							switch tt {
							case rvc.F64:
								zz = fmt.Sprintf("%f", rvcFrameDereferenced.FieldFloat64(f))
								break
							case rvc.U8:
								zz = fmt.Sprintf("%d", rvcFrameDereferenced.FieldUint8(f))
								break
							default:
								zz = "UNKNOWN TYPE"
							}

							if nn == rvc.instance {
								fmt.Printf(" DCSS1 MSG instance FOUND\n")
								var k = rvc.DGNInstanceKey{DGN: rvcFrame.DGN(), Instance: byte(rvcFrameDereferenced.FieldUint8(f))}
								var iname, ok = rvc.DGNInstanceNames[k]
								if ok {
									fmt.Printf(" DCSS1 MSG %x instance FOUND = FOUND NAME %s\n", dgn, iname)

									nnS = nnS + fmt.Sprintf("(%s)", iname)
								} else {
									fmt.Printf(" DCSS1 MSG %x instance FOUND = NOT FOUND!! instance =  %s\n", dgn, zz)

								}
							}

							fmt.Printf("DCSS1 MSGaaa name=%s type=%s value=%s dgn = %x \n", nnS, tt, zz, dgn)

						}
						//for _, fld := range *rvcFrameDereferenced.Fields() {
						//
						//}
						//var f = rvcItem(rvcFrame)

						nseen := 1 + seen[dgn]
						seen[dgn] = nseen
						var VVV = lookup(rvcFrame, dgn)
						var sa = rvcFrame.GetSourceAddress()
						fmt.Printf("\n\nZZZ RVC FRAME. DGNHI = %x sa = %x #seen = %d name = %s \n", rvcFrame.DGNHigh(), sa, nseen, VVV)
					*/

				}
			} else {
				fmt.Printf("Could not create frame for dgn: %x\n", rvcFrame.DGN())
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
		// This is all we do with messages for now
		//fmt.Println(*data)
		toSocket <- data
		nmsg++
	}
}
