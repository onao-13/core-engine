package command

type Command struct {
	Name        string
	Help        string
	Run         func(args []string) int
	Subcommands map[string]*Command
}
