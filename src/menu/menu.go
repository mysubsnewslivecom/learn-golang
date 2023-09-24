package menu

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type menuItem struct {
	name   string
	prices map[string]float64
}

var menu = []menuItem{
	{name: "Coffee", prices: map[string]float64{"small": 1.65, "medium": 1.80, "large": 2.00}},
	{name: "Espresso", prices: map[string]float64{"small": 1.90, "medium": 2.25, "large": 2.75}},
}

var in = bufio.NewReader(os.Stdin)

func MenuStart() {
loop:
	for {
		fmt.Println("Please select an option")
		fmt.Println("1) Print Menu")
		fmt.Println("2) Add Item")
		fmt.Println("q) Quit")
		choice, _ := in.ReadString('\n')

		fmt.Printf("choice: %s", choice)

		switch strings.TrimSpace(choice) {
		case "1":
			Print()
		case "2":
			Add()
		case "q":
			break loop
		default:
			fmt.Printf("Unknown option")
		}
	}
}

func Add() {
	fmt.Println("Please enter the name of the item")
	name, _ := in.ReadString('\n')
	menu = append(menu, menuItem{name: strings.TrimSpace(name)})
}

func Print() {
	for _, item := range menu {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))
		for size, cost := range item.prices {
			fmt.Printf("\t%10s%10.2f\n", size, cost)
		}
	}
}
