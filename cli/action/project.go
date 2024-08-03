package action

import (
	"core-engine/cli/common"
	"core-engine/internal/project"
	"flag"
	"github.com/rs/zerolog/log"
)

func ProjectCreate(args []string) int {
	var (
		projectName     string
		projectLocation string
	)

	flags := flag.NewFlagSet("ProjectCreate", flag.ExitOnError)
	flags.StringVar(&projectName, common.FlagProjectName, "", "Project name")
	flags.StringVar(&projectLocation, common.FlagLocation, "", "Project location")

	if err := flags.Parse(args); err != nil {
		log.Error().Err(err).Msg("Error parsing flags")
		return 1
	}

	newProject := project.NewProject(projectName, "", projectLocation, false)
	if newProject == nil {
		log.Error().Msg("Failed to create project")
		return 1
	}

	if err := newProject.Create(); err != nil {
		log.Error().Err(err).Msg("Failed to create project")
		return 1
	}

	return 0
}
