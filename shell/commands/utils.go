package commands

import (
	"github.com/c-bata/go-prompt"
)

func GetCmdByName(cmds []*Command, name string) *Command {
	for _, cmd := range cmds {
		if cmd.Name == name {
			return cmd
		}
	}

	return nil
}

func GetCmdCompletionItems(cmds []*Command) []prompt.Suggest {
	suggestions := make([]prompt.Suggest, len(cmds))

	for i, cmd := range cmds {
		suggestions[i] = cmd.GetSuggestion()
	}

	return suggestions
}
