package options

import "github.com/c-bata/go-prompt"

func GetOptByName(name string) *Option {
	for _, cmd := range Options {
		if cmd.Name == name {
			return cmd
		}
	}

	return nil
}

func GetOptCompletionItems() []prompt.Suggest {
	suggestions := make([]prompt.Suggest, len(Options))

	for i, cmd := range Options {
		suggestions[i] = cmd.GetSuggestion()
	}

	return suggestions
}
