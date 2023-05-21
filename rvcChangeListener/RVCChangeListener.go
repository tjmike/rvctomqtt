package rvcChangeListener

import (
	"fmt"
	"rvctomqtt/handler"
	"rvctomqtt/rvc"
)

func Listen(evt chan rvc.RvcItemIF) {
	// forever
	for {
		var item = <-evt

		// package this up and send it to
		fmt.Printf("EVENT: %s\n", item)
		handler.DumpItemViaReflection(&item)

	}

}
