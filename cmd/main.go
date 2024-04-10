package main

import (
	"fmt"

	"github.com/1axup/catui/catui"
)

func main() {
	fmt.Println("Hello, Go!")
	cat1 := catui.Cat{
		Name:  "Kitty",
		Icon:  "ca1",
		IsBot: false,
	}
	catui.Say("Hello, Kitty!", &cat1)
	fmt.Println(catui.Read())
}
