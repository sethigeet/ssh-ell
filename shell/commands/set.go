package commands

import "github.com/c-bata/go-prompt"

type Set struct{}

func (Set) Complete(d prompt.Document) []prompt.Suggest {
	return []prompt.Suggest{}
}

func (Set) Execute(args ...string) {
}
