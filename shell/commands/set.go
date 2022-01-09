package commands

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/c-bata/go-prompt"
	"github.com/fatih/color"
	"golang.org/x/term"

	"github.com/sethigeet/ssh-ell/ssh"
	"github.com/sethigeet/ssh-ell/utils"
)

type Set struct{}

var AuthMethodOptions = []prompt.Suggest{
	{Text: "key", Description: "Use a key based authentication for ssh"},
	{Text: "password", Description: "Use password based authentication for ssh"},
}

func (Set) Complete(d prompt.Document) []prompt.Suggest {
	cmd := utils.ParseCmd(d.TextBeforeCursor())
	if len(cmd) > 2 {
		if cmd[1] == "authMethod" {
			return prompt.FilterHasPrefix(AuthMethodOptions, d.GetWordBeforeCursor(), true)
		}
		return nil
	}

	return prompt.FilterHasPrefix(Options, d.GetWordBeforeCursor(), true)
}

func (Set) Execute(args ...string) {
	if len(args) == 1 {
		color.New(color.FgRed, color.Bold).Fprintln(os.Stderr, "Value not specified")
		color.New(color.FgRed).Fprintln(os.Stderr, "The set command take two arguments, the first being the option and the second being the value.")
		return
	}

	if conn.Connected {
		color.New(color.FgRed, color.Bold).Fprintf(os.Stderr, "You cannot set options while you are connected to a host!\n")
		return
	}

	val := strings.Join(args[1:], " ")
	switch args[0] {
	case "authMethod":
		var methArg string
		switch val {
		case "password":
			color.New(color.FgCyan, color.Bold).Printf("Enter the passowrd: ")
			passwd, err := term.ReadPassword(int(syscall.Stdin))
			if err != nil {
				color.New(color.FgRed, color.Bold).Fprintln(os.Stderr, "An error occurred while retreiving the passowrd")
				color.New(color.FgRed).Fprintf(os.Stderr, "error: %s\n", err)
			}
			fmt.Printf("\n")
			methArg = string(passwd)
			conn.AuthMethodCommonName = "password"
		case "key":
			methArg = utils.FilePathInput(
				"Enter the path to the private key file: ",
				prompt.OptionPrefixTextColor(prompt.Cyan),
			)
			conn.AuthMethodCommonName = "key"
		default:
			color.New(color.FgRed, color.Bold).Fprintln(os.Stderr, "Invalid `authMethod` specified")
			color.New(color.FgRed).Fprintln(os.Stderr, "Run `help set` to see what values can be set for all options!")
			return
		}

		meth, err := ssh.GetAuthMethod(val, methArg)
		if err != nil {
			color.New(color.FgRed, color.Bold).Fprintln(os.Stderr, "An error occurred while setting the `authMethod`")
			color.New(color.FgRed).Fprintf(os.Stderr, "error: %s\n", err)
			return
		}
		conn.AuthMethod = meth
	case "host":
		conn.Host = val
	case "port":
		port, err := strconv.Atoi(val)
		if err != nil {
			color.New(color.FgRed, color.Bold).Fprintln(os.Stderr, "Invalid value for port specified")
			color.New(color.FgRed).Fprintln(os.Stderr, "The value of port must a valid positive interger")
		}
		conn.Port = uint16(port)
	case "timeout":
		timeout, err := strconv.Atoi(val)
		if err != nil {
			color.New(color.FgRed, color.Bold).Fprintln(os.Stderr, "Invalid value for timeout specified")
			color.New(color.FgRed).Fprintln(os.Stderr, "The value of timout must a valid positive interger")
		}
		conn.Timeout = time.Second * time.Duration(timeout)
	case "user":
		conn.User = val
	default:
		color.New(color.FgRed, color.Bold).Fprintln(os.Stderr, "Invalid option specified")
		color.New(color.FgRed).Fprintln(os.Stderr, "Run `help set` to see the options that you can set!")
	}

}
