package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 1)

	ch <- "channel message"

	fmt.Println(<-ch)
}
