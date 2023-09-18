package main

import (
	"fmt"

	"github.com/mysubsnewslivecom/learn-golang/greetings"
)

func main() {
	fmt.Println("Hello World")
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
