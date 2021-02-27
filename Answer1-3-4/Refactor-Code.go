package main

import (
	"fmt"
	"strings"
)

func findFirstStringInBracket(str string) string {
	if len(str) > 0 {
		indexFirstBracket := strings.Index(str, "(")
		if indexFirstBracket >= 0 {
			wordsAfterFirstBracket := str[indexFirstBracket+1:]
			indexClosingBracketFound := strings.Index(wordsAfterFirstBracket, ")")
			if indexClosingBracketFound >= 0 {
				return wordsAfterFirstBracket[:indexClosingBracketFound]
			}
			return ""
		}
	}
	return ""
}

func main() {
	s := "Muhammad Nur )))(Fadli)"
	fmt.Println(findFirstStringInBracket(s))
}
