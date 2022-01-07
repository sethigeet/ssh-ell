package shell

import (
	"github.com/c-bata/go-prompt"

	"github.com/sethigeet/ssh-ell/shell/commands"
)

func completer(d prompt.Document) []prompt.Suggest {
	if d.TextBeforeCursor() == "" {
		return prompt.FilterHasPrefix(commands.Commands, d.GetWordBeforeCursor(), true)
	}

	return commands.Complete(d)
}
