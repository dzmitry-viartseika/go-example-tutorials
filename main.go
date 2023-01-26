package main

// название пакета

import (
	"fmt"
	"reflect"
)

func main() {
	message := ""
	message = "HelloHello Wertey message"
	fmt.Println(reflect.TypeOf(message))
}
