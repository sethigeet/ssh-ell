// Package shell provides the functionality to create a shell interface for the
// command with support for autocompletion.
package shell

import (
	"github.com/c-bata/go-prompt"
	"github.com/c-bata/go-prompt/completer"
	"github.com/fatih/color"

	"github.com/sethigeet/ssh-ell/shell/commands"
)

// InputOptions is the list of all the options that should be applied for the prompt
var InputOptions = []prompt.Option{
	prompt.OptionPrefix(config.PrefixString),
	prompt.OptionPrefixTextColor(config.PrefixTextColor),
	prompt.OptionTitle("ssh-ell"),
	prompt.OptionAddKeyBind(keybindings...),
	prompt.OptionCompletionWordSeparator(completer.FilePathCompletionSeparator), // need to set this for file path completion
	prompt.OptionSuggestionBGColor(prompt.LightGray),
	prompt.OptionSuggestionTextColor(prompt.Black),
	prompt.OptionDescriptionBGColor(prompt.DarkGray),
	prompt.OptionDescriptionTextColor(prompt.White),
	prompt.OptionSelectedSuggestionBGColor(prompt.Blue),
	prompt.OptionSelectedSuggestionTextColor(prompt.Black),
	prompt.OptionSelectedDescriptionBGColor(prompt.Blue),
	prompt.OptionSelectedDescriptionTextColor(prompt.Black),
	prompt.OptionScrollbarBGColor(prompt.LightGray),
	prompt.OptionScrollbarThumbColor(prompt.DarkGray),
}

var p *prompt.Prompt

func Setup() error {
	err := config.Parse()
	if err != nil {
		return err
	}

	p = prompt.New(
		commands.Execute,
		commands.Complete,
		InputOptions...,
	)

	return nil
}

func Run() {
	// Print a header to greet the user
	f := color.New(color.FgGreen, color.Bold)
	f.Printf(`

███████╗███████╗██╗  ██╗      ███████╗██╗     ██╗     
██╔════╝██╔════╝██║  ██║      ██╔════╝██║     ██║     
███████╗███████╗███████║█████╗█████╗  ██║     ██║     
╚════██║╚════██║██╔══██║╚════╝██╔══╝  ██║     ██║     
███████║███████║██║  ██║      ███████╗███████╗███████╗
╚══════╝╚══════╝╚═╝  ╚═╝      ╚══════╝╚══════╝╚══════╝

`)
	f = color.New(color.Bold)
	f.Println("Press Ctrl+D or type quit/exit to quit")
	f.Println("Type help to get help info")
	f.Println()

	// Start the prompt
	p.Run()
}
