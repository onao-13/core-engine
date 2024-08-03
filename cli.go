package main

type CLI struct {
}

func New() *CLI {
	return &CLI{}
}

func (cli *CLI) Run() {

}

func (cli *CLI) Usage() string {
	return ""
}
