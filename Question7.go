package main

import (
	"fmt"
	"strings"
)

func Dedup(input string) string {
	var b strings.Builder

	words := strings.Split(input, " ")
	for i, word := range words {
		// If we alredy have this word, skip.
		if strings.Contains(b.String(), word) {
			continue
		}

		b.WriteString(word)

		// If not last.
		if i < (len(words) - 1) {
			b.WriteString(" ")
		}
	}

	return b.String()
}

func main() {
	fmt.Println(Dedup("we develop sofware we craft software"))
}
