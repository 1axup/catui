package catui

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func InitCat(isBot bool) Cat {
	kitty := Cat{
		Name:  "Kitty 1",
		Icon:  "cat1",
		IsBot: isBot,
	}

	return kitty
}

func InitScreen() {
	fmt.Println("Init")
}

func (cat Cat) Say(toSay string) {
	fmt.Printf("%s says: %s\n", cat.Name, toSay)
}

func Read() string {

	fmt.Print("> ")

	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	return line
}
