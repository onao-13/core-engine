package action

import (
	"core-engine/cli/common"
	"core-engine/internal/core/language/novel-script/parser"
	"core-engine/internal/core/language/novel-script/scenario"
	"flag"
	"github.com/rs/zerolog/log"
)

func ParserParseFile(args []string) int {
	var (
		file        string
		projectPath string
	)

	flags := flag.NewFlagSet("ParserParseFile", flag.ExitOnError)
	flags.StringVar(&file, common.FlagParseFile, "", "parse file name")
	flags.StringVar(&projectPath, common.FlagProjectPath, "", "project path")

	if err := flags.Parse(args); err != nil {
		log.Error().Err(err).Msg("Error parsing flags")
		return 1
	}

	p := parser.Parser{}
	if err := p.Load(file); err != nil {
		return 1
	}

	ns, err := p.Parse()
	if err != nil {
		log.Error().Err(err).Msg("Error parsing file")
		return 1
	}

	scen := scenario.NewScenario(ns)

	err = scen.Save(projectPath)
	if err != nil {
		log.Error().Err(err).Msg("Error creating scenario")
		return 1
	}

	return 0
}
