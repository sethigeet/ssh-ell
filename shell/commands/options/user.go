package options

import "github.com/sethigeet/ssh-ell/ssh"

var User = &Option{
	Name: "user",
	Desc: "The user on the host to which you want to connect to",

	Get: func(conn *ssh.Connection) string {
		if conn.User != "" {
			return conn.User
		}

		return "-"
	},

	Set: func(val string, conn *ssh.Connection) error {
		conn.User = val

		return nil
	},
}
