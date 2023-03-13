// Monkey language interpreter
package main

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/freddiehaddad/monkey.interpreter/pkg/repl"
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
	fmt.Println("Press Ctrl+D to exit")

	repl.Start(os.Stdin, os.Stdout)

	fmt.Println("Goodbye", user.Username)
}
