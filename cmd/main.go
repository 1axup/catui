package main

import (
	"github.com/1axup/catui/catui"
)

func main() {
	cat1 := catui.Cat{
		Name:  "Kitty",
		Icon:  "cat1",
		IsBot: false,
	}

	cat1.Say("Hello,\nKitty friend!\nLorem ipsum")
	//fmt.Println(catui.Read()) // find read multi line
	//cat1.Meow()
}
