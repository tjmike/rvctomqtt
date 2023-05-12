package pform

import (
	"context"
	"go.uber.org/zap"
	"rvctomqtt/intf"
	"rvctomqtt/pool"
	"rvctomqtt/utils"
	"time"
)

func GetCANMessages(ctx *context.Context, log *zap.Logger, messagePool *pool.Pool, fromSocket chan *intf.CanFrameIF) {
	go GetCANMessagesXXX(ctx, log, messagePool, fromSocket)

}
func GetCANMessagesXXX(ctx *context.Context, log *zap.Logger, messagePool *pool.Pool, fromSocket chan *intf.CanFrameIF) {

	log = utils.ApplyContext(ctx, log)

	//f := map[string]zap.Field(ctx.Value("logFields"))
	//var f = zap.Field{
	//	Key:       "workID",
	//	Type:      zapcore.StringType,
	//	Integer:   0,
	//	String:    "CanReader1",
	//	Interface: nil,
	//}
	//log.With(f).Info("ZAP LOG")
	//log.With(f).Warn("ZAP LOG")
	log.Warn("ZAP LOG")
	log.Info("ZAP LOG")

	//fmt.Printf("IMPLEMENT ME\n")
	//XXX var m = &rvccan.Frame{}

	//var mm intf.CanFrameIF = &rvccan.Frame{}
	var pass = 0
	for {
		{
			pass++
			// TS:1681751409945040000 - canID=19FFFD42 DGNH = 1ff DGNL=fd ERR=false RTR=false EXT=true Data Len=8 Data=01640e0180553577 RAW=42fdff990800000001640e0180553577

			var f = messagePool.Get()
			(*f).SetTimeStamp(time.Now())

			if (pass % 2) == 0 {

				(*f).GetMessage()[0] = 0x42
				(*f).GetMessage()[1] = 0xfd
				(*f).GetMessage()[2] = 0xff
				(*f).GetMessage()[3] = 0x99
				(*f).GetMessage()[4] = 0x08
				(*f).GetMessage()[5] = 0
				(*f).GetMessage()[6] = 0
				(*f).GetMessage()[7] = 0
				(*f).GetMessage()[8] = 0x01
				(*f).GetMessage()[9] = 0x02
				(*f).GetMessage()[10] = 0
				(*f).GetMessage()[11] = 0
				(*f).GetMessage()[12] = 0
				(*f).GetMessage()[13] = 0
				(*f).GetMessage()[14] = 0
				(*f).GetMessage()[15] = 0
			} else {
				// RAW=8c da fe 99 08 00 00 00    3b 7c c8 fc 02 01 04 ff
				(*f).GetMessage()[0] = 0x8c
				(*f).GetMessage()[1] = 0xda
				(*f).GetMessage()[2] = 0xfe
				(*f).GetMessage()[3] = 0x99
				(*f).GetMessage()[4] = 0x08
				(*f).GetMessage()[5] = 0
				(*f).GetMessage()[6] = 0
				(*f).GetMessage()[7] = 0
				(*f).GetMessage()[8] = 0x01
				(*f).GetMessage()[9] = 0x7c
				(*f).GetMessage()[10] = 0xc8
				(*f).GetMessage()[11] = 0xfc
				(*f).GetMessage()[12] = 0x02
				(*f).GetMessage()[13] = 0x01
				(*f).GetMessage()[14] = 0x04
				(*f).GetMessage()[15] = 0xff
			}
			(*f).BuildCanFrameX()
			//m.timestamp = time.Now()
			//m.MessageBytes[0] = 1
			// we've read the frame and are now sending to the channel
			fromSocket <- f
			time.Sleep(time.Millisecond * 5)

		}
	}

}
