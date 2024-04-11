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
	fmt.Println(" /--------------\\")
	fmt.Println(" | Hello Friend |")
	fmt.Println(" \\--------------/")
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

func catify(input []string) []string {
	baseIndex := len(input)

	if baseIndex < 3 {
		return []string{}
	}

	input[(baseIndex - 3)] = swapIt(input[(baseIndex-3)], "/\\___/\\    ")
	input[(baseIndex - 2)] = swapIt(input[(baseIndex-2)], "|* . *|   /")
	input[(baseIndex - 1)] = swapIt(input[(baseIndex-1)], "\\_____/ \\/ ")

	return input
}

func swapIt(input string, swapText string) string {
	if len(swapText) > len(input) {
		return input
	}

	inputBytes := []byte(input)
	swapBytes := []byte(swapText)

	for i := 0; i < len(swapBytes); i++ {
		inputBytes[i] = swapBytes[i]
	}

	return string(inputBytes)
}

func getSpacesMissing(text string, chars int) string {
	missing := chars - len(text)
	out := ""
	if missing >= 1 {
		for i := 0; i < missing; i++ {
			out += " "
		}
	}
	return out
}

func findLongestLine(text []string) int {

	longest := 0

	for _, line := range text {
		if longest < len(line) {
			longest = len(line)
		}
	}

	return longest
}

func getLinesToText(len int) string {
	str := ""
	for i := 0; i < (len + 2); i++ {
		str += "-"
	}
	return str
}

func getCatOffset(offset int) string {

	str := ""
	// create header line
	for i := 0; i < offset; i++ {
		str += " "
	}

	return str
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
