package commands

import "github.com/c-bata/go-prompt"

// Commands is the list of the commands supported by the shell package
var Commands = []prompt.Suggest{
	{Text: "connect", Description: "Connect using the set options"},
	{Text: "disconnect", Description: "Disconnect from the ssh connection"},
	{Text: "download", Description: "Download a file from the connected server"},
	{Text: "get", Description: "Get the value of an option"},
	{Text: "help", Description: "Print out help info"},
	{Text: "set", Description: "Set the value of an option"},
	{Text: "shell", Description: "Start an interactive shell on the connected host"},
	{Text: "status", Description: "Get the status of the connection"},
	{Text: "upload", Description: "Upload a file to the connected server"},
}

// Options is the list of options supported by the `set` and `get` commands
// which set and get the options required by other commands
var Options = []prompt.Suggest{
	{Text: "authMethod", Description: "The auth method you want to use to connect to the host"},
	{Text: "host", Description: "The host you want to connect to"},
	{Text: "port", Description: "The port of the host to which you want to connect to"},
	{Text: "terminalName", Description: "The name of the terminal to be used when starting a shell"},
	{Text: "timeout", Description: "The timeout for the ssh connection"},
	{Text: "user", Description: "The user on the host to which you want to connect to"},
}

// AuthMethodOptions is the list of the options allowed to be used for the authMethod option
var AuthMethodOptions = []prompt.Suggest{
	{Text: "key", Description: "Use a key based authentication for ssh"},
	{Text: "password", Description: "Use password based authentication for ssh"},
}
