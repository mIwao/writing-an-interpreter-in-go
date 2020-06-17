package main

import (
	"fmt"
	"os"
	"os/user"
	"writing-an-interpreter-in-go/src/monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programing language!\n", user.Username)
	fmt.Printf("Feel free to type in commnads\n")
	repl.Start(os.Stdin, os.Stdout)
}
