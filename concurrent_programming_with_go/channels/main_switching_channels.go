package main

import (
	"fmt"
)

// Message comentado
type Message struct {
	To      []string
	From    string
	Content string
}

// FailedMessage comentado
type FailedMessage struct {
	ErrorMessage    string
	OriginalMessage Message
}

func main() {
	msgCh := make(chan Message, 1)
	errorCh := make(chan FailedMessage, 1)

	msgCh <- Message{
		To:      []string{"mike"},
		From:    "maikel",
		Content: "Contenido del mensaje",
	}

	errorCh <- FailedMessage{
		ErrorMessage:    "Contenido del mensaje de error",
		OriginalMessage: Message{},
	}

	// will read from first case first, if founds a message then following cases (after the matched cases) will be ignored.
	//  A default case may be used to avoid deadlock if no channels receive messages
	select {
	case msg := <-msgCh:
		fmt.Println("From: " + msg.From + " -> " + msg.Content)
	case errorMsg := <-errorCh:
		fmt.Println("Error!!! " + errorMsg.ErrorMessage)
	default:
		fmt.Println("Did not received messages in any channel")
	}
}
