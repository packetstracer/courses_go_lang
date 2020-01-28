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

	close(ch)

	// listen to the channel until is closed (indicated by ok)
	// for {
	// 	if msg, ok := <-ch; ok {
	// 		fmt.Print(msg + "+")
	// 	} else {
	// 		break
	// 	}
	// }

	// range the values of the cannel
	for readValueFromChannel := range ch {
		fmt.Print(readValueFromChannel + "--")
	}
}
