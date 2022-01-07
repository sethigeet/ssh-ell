package commands

import "github.com/c-bata/go-prompt"

type Download struct{}

func (Download) Complete(d prompt.Document) []prompt.Suggest {
	return []prompt.Suggest{}
}

func (Download) Execute(args ...string) {
}
