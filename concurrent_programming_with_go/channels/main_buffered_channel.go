package main

import (
	"fmt"
	"strings"
)

func main() {
	phrase := "This is a test sentence to be processed by the receiver"
	words := strings.Split(phrase, " ")

	ch := make(chan string, len(words))

	for _, word := range words {
		ch <- word
	}

	for i := 0; i < len(words); i++ {
		fmt.Print(<-ch + "-")
	}
}
