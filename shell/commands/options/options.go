package options

import (
	"github.com/c-bata/go-prompt"
	"github.com/sethigeet/ssh-ell/ssh"
)

type Option struct {
	Name, Desc string

	Set      func(val string, conn *ssh.Connection) error
	Get      func(conn *ssh.Connection) string
	Validate func(conn *ssh.Connection) bool
}

func (opt *Option) GetSuggestion() prompt.Suggest {
	return prompt.Suggest{Text: opt.Name, Description: opt.Desc}
}

// Options is the list of options supported by the `set` and `get` commands
// which set and get the options required by other commands
var Options = []*Option{
	AuthMethod,
	Host,
	Port,
	TerminalName,
	Timeout,
	User,
}

// AuthMethodOptions is the list of the options allowed to be used for the authMethod option
var AuthMethodOptions = []prompt.Suggest{
	{Text: "key", Description: "Use a key based authentication for ssh"},
	{Text: "password", Description: "Use password based authentication for ssh"},
}
