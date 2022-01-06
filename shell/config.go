package shell

import "github.com/c-bata/go-prompt"

type Config struct {
	PrefixString    string
	PrefixTextColor prompt.Color
}

var config = Config{
	PrefixString:    "❯ ",
	PrefixTextColor: prompt.Green,
}

// Parse parses the options provided on the command line while executing the
// program and sets them on the config variable appropriately
func (c Config) Parse() error {
	// TODO: Parse the options here

	return nil
}
