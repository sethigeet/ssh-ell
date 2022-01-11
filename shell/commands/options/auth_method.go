package options

import (
	"fmt"
	"syscall"

	"github.com/c-bata/go-prompt"
	"github.com/fatih/color"
	"github.com/sethigeet/ssh-ell/ssh"
	"github.com/sethigeet/ssh-ell/utils"
	"golang.org/x/term"
)

var AuthMethod = &Option{
	Name: "authMethod",
	Desc: "The auth method you want to use to connect to the host",

	Get: func(conn *ssh.Connection) string {
		if conn.AuthMethodCommonName != "" {
			return conn.AuthMethodCommonName
		}

		return "-"
	},

	Set: func(meth string, conn *ssh.Connection) error {
		var methArg string
		switch meth {
		case "password":
			color.New(color.FgCyan, color.Bold).Printf("Enter the passowrd: ")
			passwd, err := term.ReadPassword(int(syscall.Stdin))
			if err != nil {
				return err
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
			return fmt.Errorf("invalid `authMethod` specified")
		}

		authMeth, err := ssh.GetAuthMethod(meth, methArg)
		if err != nil {
			return err
		}

		conn.AuthMethod = authMeth

		return nil
	},

	Validate: func(conn *ssh.Connection) bool {
		return conn.AuthMethod != nil
	},
}
