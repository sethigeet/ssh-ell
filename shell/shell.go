// Package shell provides the functionality to create a shell interface for the
// command with support for autocompletion.
package shell

import (
	"github.com/c-bata/go-prompt"
	gpc "github.com/c-bata/go-prompt/completer"
	"github.com/fatih/color"
)

var p *prompt.Prompt

func Setup() error {
	err := config.Parse()
	if err != nil {
		return err
	}

	p = prompt.New(
		executor,
		completer,
		// Options
		prompt.OptionPrefix(config.PrefixString),
		prompt.OptionPrefixTextColor(config.PrefixTextColor),
		prompt.OptionTitle("ssh-ell"),
		prompt.OptionAddKeyBind(keybindings...),
		prompt.OptionCompletionWordSeparator(gpc.FilePathCompletionSeparator), // need to set this for file path completion
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
	)

	return nil
}

func Run() {
	f := color.New(color.FgGreen, color.Bold)
	f.Println(`

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
	p.Run()
}
