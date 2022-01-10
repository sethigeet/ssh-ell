package options

import (
	"fmt"
	"strconv"

	"github.com/sethigeet/ssh-ell/ssh"
)

var Port = &Option{
	Name: "port",
	Desc: "The port of the host to which you want to connect to",

	Get: func(conn *ssh.Connection) string {
		return strconv.Itoa(int(conn.Port))
	},

	Set: func(val string, conn *ssh.Connection) error {
		port, err := strconv.Atoi(val)
		if err != nil {
			return fmt.Errorf("the value of port must a valid positive interger")
		}

		conn.Port = uint16(port)

		return nil
	},
}
