package commands

import "github.com/c-bata/go-prompt"

type Connect struct{}

func (Connect) Complete(d prompt.Document) []prompt.Suggest {
	return []prompt.Suggest{}
}

func (Connect) Execute(args ...string) {
}
