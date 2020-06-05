package main

import (
	"strings"
)

func reverseWords(s string) (reversed string) {
	words := strings.Fields(s)
	for i := len(words) - 1; i >= 0; i-- {
		reversed += words[i] + " "
	}
	reversed = strings.TrimSpace(reversed)
	return
}

