package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/tkircsi/monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Println("Type commands....")

	repl.Start(os.Stdin, os.Stdout)
}
