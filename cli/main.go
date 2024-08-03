package main

import (
	"core-engine/cli/command"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	//fmt.Println(args)

	baseCommand := command.Select(args)
	fmt.Println(baseCommand)
}
