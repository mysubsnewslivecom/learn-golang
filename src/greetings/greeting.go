package greetings

import (
	"fmt"
)

// var in = bufio.NewReader(os.Stdin)

func Greetings(name string) {
	fmt.Printf("Greetings %s!!!\n", getName(name))
}

func getName(name string) string {
	return name
}

// func start() {
// 	fmt.Println("hello world")
// 	Greetings("Linux")

// 	choice, _ := in.ReadString('\n')

// 	switch strings.TrimSpace(choice) {
// 	case "1":
// 		fmt.Printf("Selected = %s ", choice)
// 	default:
// 		fmt.Printf("Default = %s ", choice)
// 	}
// }
