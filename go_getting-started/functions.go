package main

import (
	"errors"
	"fmt"
)

func main() {
	port := 3000
	// _ is a write only var, so is like ignoring the output and avoid compiling errors of declared var but not used
	_, error := startWebserver(port)
	fmt.Println(port, error)
}

func startWebserver(port int) (int, error) {
	fmt.Println("Starting...")
	fmt.Println("Started on port", port)

	return port, errors.New("Something went wrong!")
}
