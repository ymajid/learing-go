package main

import (
	"fmt"
	"os"
	"os/user"

	"monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("hello %s\nthis is the monkey programming language\n", user.Username)
	fmt.Printf("type in commands into the repl")
	repl.Start(os.Stdin, os.Stdout)
}
