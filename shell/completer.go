package shell

import "github.com/c-bata/go-prompt"

var commands = []prompt.Suggest{
	{Text: "help", Description: "Print out help info"},
	{Text: "set", Description: "Set the value of an option"},
	{Text: "get", Description: "Get the value of an option"},
	{Text: "connect", Description: "Connect using the set options"},
	{Text: "upload", Description: "Upload a file to the connected server"},
	{Text: "download", Description: "Download a file from the connected server"},
}

func completer(d prompt.Document) []prompt.Suggest {
	if d.TextBeforeCursor() == "" {
		return prompt.FilterHasPrefix(commands, d.GetWordBeforeCursor(), true)
	}
	return prompt.FilterFuzzy(commands, d.GetWordBeforeCursor(), true)
}
