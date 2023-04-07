package hello

import (
	"fmt"
)

func init() {
	fmt.Println("INIT HELLO")
}

func Hello() string {
	return "Hello, big World!"
}
