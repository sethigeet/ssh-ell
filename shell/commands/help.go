package commands

import "github.com/c-bata/go-prompt"

type Help struct{}

func (Help) Complete(d prompt.Document) []prompt.Suggest {
	return []prompt.Suggest{}
}

func (Help) Execute(args ...string) {
}
