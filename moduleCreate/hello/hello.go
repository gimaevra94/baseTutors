package main

import (
	"fmt"
	"home/greetings"
	"log"
)

func main() {
	log.SetPrefix(("greetings: "))
	log.SetFlags(0)

	names := []string{"хуй", "пизда", "давайдавай"}
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
