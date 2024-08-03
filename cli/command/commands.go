package command

import (
	"core-engine/cli/action"
	"core-engine/cli/common"
	"fmt"
)

var Commands = map[string]*Command{
	common.CommandProject: CommandProject,
	common.CommandParser:  CommandNovelScriptParser,
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
	CommandNovelScriptParser = &Command{
		Name: common.CommandParser,
		Help: "Parse Novel-Script files",
		Subcommands: map[string]*Command{
			common.SubcommandParserParse: {
				Name: common.SubcommandParserParse,
				Run:  action.ParserParseFile,
			},
		},
	}
)
