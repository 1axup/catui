package catui

type Cat struct {
	Name  string
	Icon  string // represents the name of an ASCII Art
	IsBot bool   // if its the bot or the person chatting
}

func (cat Cat) getSymbolArray() []string {

	symbol := []string{}

	if !cat.IsBot {
		symbol = append(symbol, "/\\___/\\    ")
		symbol = append(symbol, "|* . *|   /")
		symbol = append(symbol, "\\_____/ \\/ ")
	} else {
		symbol = append(symbol, "/-----\\    ")
		symbol = append(symbol, "| @ @ |   /")
		symbol = append(symbol, "\\-----/ \\/ \a")
	}

	return symbol
}
