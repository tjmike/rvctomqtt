package main

import (
	"fmt"
	svc "rvctomqtt/reg"

	//"github.com/tjmike/rvctomqtt/internal/hello"
	//"github.com/tjmike/rvctomqtt/pform"
	//"github.com/tjmike/rvctomqtt/reg"
	"rvctomqtt/internal/hello"
	"rvctomqtt/pform"
)

func main() {

	fmt.Println(hello.Hello())

	//svc.single ss = svc.GetInstance()
	var ss *svc.Single = svc.GetInstance()
	//svc.Single foo  = &Single{};

	fmt.Printf("%#v\n", ss)

	pform.Doit()

}
