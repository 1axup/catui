package main

import (
	"fmt"

	"github.com/1axup/catui/catui"
)

func main() {
	fmt.Println("Hello, Go!")

	cat1 := catui.Cat{
		Name:  "Kitty",
		Icon:  "cat1",
		IsBot: false,
	}

	cat1.Say("Hello, Kitty!")
	//fmt.Println(catui.Read())
}
