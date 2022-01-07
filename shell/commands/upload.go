package commands

import "github.com/c-bata/go-prompt"

type Upload struct{}

func (Upload) Complete(d prompt.Document) []prompt.Suggest {
	return []prompt.Suggest{}
}

func (Upload) Execute(args ...string) {
}
