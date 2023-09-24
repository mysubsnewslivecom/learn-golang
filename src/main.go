package main

import (
	"fmt"

	"github.com/mysubsnewslivecom/learn-golang/src/greetings"
	"github.com/mysubsnewslivecom/learn-golang/src/menu"
)

func main() {
	fmt.Println("hello world")
	greetings.Greetings("Linux")
	menu.MenuStart()
}
