package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/fatih/color"

	"github.com/sethigeet/ssh-ell/utils"
)

type Help struct{}

var helpMessages = map[string]string{
	"ssh-ell": `Write a generic help message for ssh-ell here`,

	"connect":    `Write a help message for the connect command here`,
	"disconnect": `Write a help message for the disconnect command here`,
	"download":   `Write a help message for the download command here`,
	"get":        `Write a help message for the get command here`,
	"help":       `Write a help message for the help command here`,
	"set":        `Write a help message for the set command here`,
	"status":     `Write a help message for the status command here`,
	"upload":     `Write a help message for the upload command here`,
}

func (Help) Complete(d prompt.Document) []prompt.Suggest {
	cmd := utils.ParseCmd(d.TextBeforeCursor())
	if len(cmd) > 2 {
		return nil
	}

	return prompt.FilterHasPrefix(Commands, d.GetWordBeforeCursor(), true)
}

func (Help) Execute(args ...string) {
	if len(args) > 1 {
		color.New(color.FgRed, color.Bold).Fprintln(os.Stderr, "Invalid args specified")
		color.New(color.FgRed).Fprintln(os.Stderr, "The help command can only take one arg at max!")
		return
	}

	var cmd string
	if len(args) == 0 {
		cmd = "ssh-ell"
	} else {
		cmd = args[0]
	}

	helpMessage, exists := helpMessages[cmd]

	if !exists {
		color.New(color.FgRed, color.Bold).Fprintln(os.Stderr, "Invalid command specified")
		color.New(color.FgRed).Fprintln(os.Stderr, "Please ask for help for a command that exists!")
		return
	}

	header := fmt.Sprintf("Help for %s:", cmd)
	color.New(color.FgGreen, color.Bold).Println(header)
	log := color.New(color.FgGreen)
	log.Println(strings.Repeat("-", len(header)))
	log.Println(helpMessage)
}
