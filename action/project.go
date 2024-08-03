package action

import (
	"core-engine/cli/common"
	"flag"
	"fmt"
)

func ProjectCreate(args []string) int {
	flagProjectName := flag.NewFlagSet(common.FlagProjectName, flag.ExitOnError)
	flagProjectLocation := flag.NewFlagSet(common.FlagLocation, flag.ExitOnError)

	flagProjectName.Parse(args)
	flagProjectLocation.Parse(args)

	fmt.Printf("Project name: %s; Location: %s", flagProjectName.Args(), flagProjectLocation.Args())

	return 0
}
