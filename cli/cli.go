package main

import (
	"core-engine/cli/command"
	"github.com/rs/zerolog/log"
)

type CLI struct {
	IsDebug bool
}

func NewCLI(isDebug bool) *CLI {
	return &CLI{
		IsDebug: isDebug,
	}
}

func (cli *CLI) Run(args []string) int {
	argLevel, curCommand := command.Select(args)
	if curCommand == nil {
		return 1
	}

	if cli.IsDebug {
		log.Info().Msgf("Current Command: %s", curCommand.Name)
	}

	args = args[argLevel:]

	return curCommand.Run(args)
}

func (cli *CLI) Usage() string {
	return ""
}
