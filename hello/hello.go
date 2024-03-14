package main

import (
	"fmt"

	"yauhen.example.com/greetings"
)

func main() {
	message := greetings.Hello("Yauhen")
	fmt.Println(message)
}
