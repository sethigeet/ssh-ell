package options

import (
	"fmt"

	"github.com/c-bata/go-prompt"
	"github.com/sethigeet/ssh-ell/ssh"
)

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

func ValidateAllOpts(conn *ssh.Connection) error {
	for _, opt := range Options {
		isValid := opt.Validate(conn)
		if !isValid {
			return fmt.Errorf("the %s option must be set", opt.Name)
		}
	}

	return nil
}
