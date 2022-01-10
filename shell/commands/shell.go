package commands

import (
	"os"

	"github.com/c-bata/go-prompt"
	"github.com/fatih/color"
)

type Shell struct{}

func (Shell) Complete(d prompt.Document) []prompt.Suggest {
	return nil
}

func (Shell) Execute(args ...string) {
	if len(args) != 0 {
		color.New(color.FgRed, color.Bold).Fprintln(os.Stderr, "Invalid number of args specified")
		color.New(color.FgRed).Fprintln(os.Stderr, "The shell command does not take any args")
		return
	}

	if !conn.Connected {
		color.New(color.FgRed, color.Bold).Fprintf(os.Stderr, "You are not connected to any host!\n")
		return
	}

	err := conn.Shell()
	if err != nil {
		color.New(color.FgRed, color.Bold).Fprintln(os.Stderr, "An error occurred while trying to start a shell")
		color.New(color.FgRed).Fprintf(os.Stderr, "error: %s\n", err)
		return
	}
}
