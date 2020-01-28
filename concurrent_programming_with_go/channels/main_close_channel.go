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

	// closes the channel for the sender but not the receiver so the receiver can still read from it
	close(ch)

	for i := 0; i < len(words); i++ {
		fmt.Print(<-ch + "-")
	}

	// this line would case an error due to the channel is closed so no senders can write on it
	// ch <- "other value"
}
