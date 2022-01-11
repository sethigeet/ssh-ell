package options

import "github.com/sethigeet/ssh-ell/ssh"

var TerminalName = &Option{
	Name: "terminalName",
	Desc: "The name of the terminal to be used when starting a shell",

	Get: func(conn *ssh.Connection) string {
		if conn.TerminalName != "" {
			return conn.TerminalName
		}

		return "-"
	},

	Set: func(val string, conn *ssh.Connection) error {
		conn.TerminalName = val

		return nil
	},

	Validate: func(conn *ssh.Connection) bool {
		return conn.TerminalName != ""
	},
}
