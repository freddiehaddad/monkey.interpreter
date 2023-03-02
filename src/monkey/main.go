// Monkey language interpreter
package main

import (
	"fmt"
	"log"
	"os/user"
)

func init() {
	// Initialize the logger
	flags := log.Ldate | log.Ltime | log.LUTC | log.Lshortfile
	log.SetFlags(flags)
}

func main() {
	user, err := user.Current()
	if err != nil {
		log.Fatalf("Failed to get username: %s", err)
	}
	fmt.Printf("Hello %s! Welcome to the Monkey Interpreter.\n", user.Username)
	fmt.Println("This project is still in development and the REPL is not ready yet :(")
}
