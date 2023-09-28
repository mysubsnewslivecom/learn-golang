package main

import (
	"fmt"

	"github.com/mysubsnewslivecom/learn-golang/src/apis"
	"github.com/mysubsnewslivecom/learn-golang/src/greetings"
	"github.com/mysubsnewslivecom/learn-golang/src/menu"
)

func main() {
	fmt.Println("hello world")
	greetings.Greetings("Linux")
	apis.BoredApi()
	menu.MenuStart()
}
