package catui

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
