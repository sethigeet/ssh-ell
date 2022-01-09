package commands

import (
	"os"

	"github.com/c-bata/go-prompt"
	"github.com/fatih/color"
)

type Upload struct{}

func (Upload) Complete(d prompt.Document) []prompt.Suggest {
	return []prompt.Suggest{}
}

func (Upload) Execute(args ...string) {
	if !conn.Connected {
		color.New(color.FgRed, color.Bold).Fprintf(os.Stderr, "You are not connected to any host!\n")
		return
	}
}
