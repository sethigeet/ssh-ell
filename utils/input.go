package utils

import (
	"github.com/c-bata/go-prompt"
	"github.com/c-bata/go-prompt/completer"
)

// InputOptions is the list of all the options that should be applied for the prompt
var InputOptions = []prompt.Option{
	prompt.OptionPrefixTextColor(prompt.Cyan),
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

func FilePathInput(prefix string, opts ...prompt.Option) string {
	fpCompleter := completer.FilePathCompleter{}

	opts = append(InputOptions, opts...)

	return prompt.Input(
		prefix,
		fpCompleter.Complete,
		opts...,
	)
}
