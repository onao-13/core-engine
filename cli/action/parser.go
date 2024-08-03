package action

import (
	"core-engine/cli/common"
	"flag"
	"github.com/rs/zerolog/log"
)

func ParserParseFile(args []string) int {
	var (
		file string
	)

	flags := flag.NewFlagSet("ParserParseFile", flag.ExitOnError)
	flags.StringVar(&file, common.FlagParseFile, "", "parse file name")

	if err := flags.Parse(args); err != nil {
		log.Error().Err(err).Msg("Error parsing flags")
		return 1
	}

	return 0
}
