package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args

	cmdLine := strings.Join(args, " ")

	fmt.Printf("Command-line used to run the program: %s\n", cmdLine)
}
