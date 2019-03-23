package main

import (
	"Writing-Interpreter-In-Go/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey Programming Language!\n", user.Name)
	fmt.Println("Feel free to type any command")
	repl.Start(os.Stdin, os.Stdout)
}