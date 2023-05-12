package pform

// This is the public facing interface. It will call platform specific code.
// In this example doit (lower case is different code on darwin vs linux)

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"rvctomqtt/intf"
	"rvctomqtt/pool"
)

func init() {
	fmt.Printf("rvcp init\n")
}

func Doit() {
	fmt.Printf("rvcp doit\n")
	doit()
}

func GetRVCMessages(ctx *context.Context, log *zap.Logger, pool *pool.Pool, fromSocket chan *intf.CanFrameIF) {
	GetCANMessages(ctx, log, pool, fromSocket)
}

func ByteToUint() func([]byte) uint32 {
	return byteToUint()
}
