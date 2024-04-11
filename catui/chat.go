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

func (cat Cat) Meow() {
	fmt.Println(" /--------\\")
	fmt.Println(" | Meow! |")
	fmt.Println(" \\-------/")
	PrintCat()
}

func PrintCat() {
	fmt.Println("  \\ ")
	fmt.Println("   \\ ")
	fmt.Println("    /\\___/\\")
	fmt.Println("   /       \\")
	fmt.Println("   | * . * | ")
	fmt.Println("   \\_______/")
}

func (cat Cat) Say(toSay string) {
	//fmt.Printf("%s: %s\n", cat.Name, toSay)

	arr := makeBubbleArray(toSay, cat.getSymbolSize())
	arr = catify(arr, cat)

	for _, line := range arr {
		fmt.Println(line)
	}
	fmt.Println()
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
