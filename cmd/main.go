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

	bot1 := catui.Cat{
		Name:  "Botti",
		Icon:  "bot1",
		IsBot: true,
	}

	cat1.Say("Hello,\nKitty friend!\nLorem ipsum")
	bot1.Say("Hello,\nBotti friend!\nLorem ipsum")
}
