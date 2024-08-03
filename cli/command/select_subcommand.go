package command

import (
	"fmt"
	"github.com/samber/lo"
)

func Select(args []string) (int, *Command) {
	baseCommand := selectCommand(args[0])
	if baseCommand == nil {
		return 0, nil
	}

	if baseCommand.Subcommands == nil {
		if baseCommand.Run == nil {
			return 0, nil
		}
		return 1, baseCommand
	}

	curCommand, argLevel := findSubcommandIntoArgs(baseCommand, args)

	if curCommand == nil || curCommand.Run == nil {
		return 1, nil
	}

	// clear subcommand
	args = args[argLevel:]

	fmt.Println(curCommand.Name)

	return argLevel, curCommand
}

func findSubcommandIntoArgs(com *Command, args []string) (*Command, int) {
	var (
		curCommand *Command = com
		argLevel            = 1
	)

	for _, arg := range args[1:] {
		destSubcommand := selectSubcommandCommand(curCommand, arg)
		if destSubcommand == nil {
			break
		}
		curCommand = destSubcommand
		argLevel += 1
	}

	return curCommand, argLevel
}

func selectSubcommandCommand(com *Command, name string) *Command {
	if com == nil || com.Subcommands == nil {
		return nil
	}
	baseSubcommand, ok := com.Subcommands[name]
	return lo.Ternary(!ok || baseSubcommand == nil, nil, baseSubcommand)
}

func selectCommand(name string) *Command {
	baseCommand, ok := Commands[name]
	return lo.Ternary(!ok || baseCommand == nil, nil, baseCommand)
}
