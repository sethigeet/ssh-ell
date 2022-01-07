package commands

import "github.com/c-bata/go-prompt"

type Get struct{}

func (Get) Complete(d prompt.Document) []prompt.Suggest {
	return []prompt.Suggest{}
}

func (Get) Execute(args ...string) {
}
