package main

import (
	"Interpreter/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s!\n", user.Username)
	fmt.Printf("Input 'Ctrl+z' to exit\n")
	repl.Start(os.Stdin, os.Stdout)
}
