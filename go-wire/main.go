package main

import (
	"fmt"
	"log"
)

func main() {
	// m := newMessage("Hello google wire")
	// g := newGreeter(m)
	// e := newEvent(g)
	e, err := initializeEvent("Hello google wire[3]")
	if err != nil {
		log.Fatalln("ERROR", err)
	}
	e.start()
}

func newMessage(msg string) message {
	return message{value: msg}
}

type message struct {
	value string
}

func newGreeter(message message) greeter {
	return greeter{message: message}
}

type greeter struct {
	message message
}

func newEvent(g greeter) event {
	return event{greeter: g}
}

type event struct {
	greeter greeter
}

func (e event) start() {
	fmt.Println(e.greeter.message.value)
}
