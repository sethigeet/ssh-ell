package commands

import "github.com/c-bata/go-prompt"

type Disconnect struct{}

func (Disconnect) Complete(d prompt.Document) []prompt.Suggest {
	return []prompt.Suggest{}
}

func (Disconnect) Execute(args ...string) {
}
