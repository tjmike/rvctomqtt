package svc

import (
	"fmt"
	"sync"
)

func init() {
	fmt.Println("init svc")
}

var lock = &sync.Mutex{}

type Single struct {
	val string
}

var singleInstance *Single

func GetInstance() *Single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &Single{val: "unset"}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}
