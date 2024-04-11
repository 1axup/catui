package catui

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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
	printCat()
}

func printCat() {
	fmt.Println("  \\ ")
	fmt.Println("   \\ ")
	fmt.Println("    /\\___/\\")
	fmt.Println("   /       \\")
	fmt.Println("   | * . * | ")
	fmt.Println("   \\_______/")
}

func (cat Cat) Say(toSay string) {
	//fmt.Printf("%s: %s\n", cat.Name, toSay)

	arr := makeBubbleArray(toSay)
	arr = catify(arr)

	for _, line := range arr {
		fmt.Println(line)
	}
}

func makeBubbleArray(text string) []string {

	textFormatted := []string{}

	unformatted := strings.Split(text, "\n")

	longestCharLen := findLongestLine(unformatted)

	catOffset := 11 // TODO: THIS IS BULLSHIT

	textFormatted = append(textFormatted, getCatOffset(catOffset)+"/"+getLinesToText(longestCharLen)+"\\")

	for _, line := range unformatted {
		textFormatted = append(textFormatted, getCatOffset(catOffset)+"| "+line+getSpacesMissing(line, longestCharLen)+" |")
	}

	textFormatted = append(textFormatted, getCatOffset(catOffset)+"\\"+getLinesToText(longestCharLen)+"/")

	return textFormatted
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
