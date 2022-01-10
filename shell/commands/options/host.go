package options

import (
	"github.com/sethigeet/ssh-ell/ssh"
)

var Host = &Option{
	Name: "host",
	Desc: "The host you want to connect to",

	Get: func(conn *ssh.Connection) string {
		if conn.Host != "" {
			return conn.Host
		}

		return "-"
	},

	Set: func(val string, conn *ssh.Connection) error {
		conn.Host = val
		return nil
	},
}
