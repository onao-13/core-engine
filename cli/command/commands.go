package command

import (
	"core-engine/cli/action"
	"core-engine/cli/common"
	"fmt"
)

var Commands = map[string]*Command{
	common.CommandProject: CommandProject,
}

var (
	CommandProject = &Command{
		Name: common.CommandProject,
		Help: fmt.Sprintf("Created new project. Use flags %s, %s", common.FlagProjectName, common.FlagLocation),
		Subcommands: map[string]*Command{
			common.SubcommandProjectCreate: {
				Name: common.SubcommandProjectCreate,
				Run:  action.ProjectCreate,
			},
		},
	}
)
