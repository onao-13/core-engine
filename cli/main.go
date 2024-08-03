package main

import (
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) <= 0 {
		os.Exit(1)
	}

	cli := NewCLI(true)
	cli.Run(args)
}
