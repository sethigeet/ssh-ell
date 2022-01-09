package utils

import "strings"

var sep = " "
var quote = "\""

// ParseCmd parses a string(the cmd) appropriately by taking into account
// quotes, etc. and returns an array of strings containing each part of the
// command
func ParseCmd(cmd string) []string {
	var parsed []string

	var toConcat string
	for _, c := range strings.Split(cmd, sep) {
		cReplaced := strings.ReplaceAll(c, quote, "")
		cContainsQuote := strings.Contains(c, quote)
		toConcatIsNotEmpty := toConcat != ""
		if cContainsQuote || toConcatIsNotEmpty {
			if toConcatIsNotEmpty {
				toConcat += sep
			}
			toConcat += cReplaced
			if cContainsQuote && toConcatIsNotEmpty && toConcat != cReplaced {
				parsed = append(parsed, toConcat)
				toConcat = ""
			}

			continue
		}

		parsed = append(parsed, c)
	}

	if toConcat != "" {
		parsed = append(parsed, toConcat)
	}

	return parsed
}
