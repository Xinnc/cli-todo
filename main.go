package main

import (
	"WorkWithFiles/cli"
	"os"
)

func main() {
	var command string
	if len(os.Args) > 1 {
		command = os.Args[1]
	} else {
		command = "help"
	}

	cli.RunCLI(command)
}
