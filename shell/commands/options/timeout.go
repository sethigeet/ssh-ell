package options

import (
	"fmt"
	"strconv"
	"time"

	"github.com/sethigeet/ssh-ell/ssh"
)

var Timeout = &Option{
	Name: "timeout",
	Desc: "The timeout for the ssh connection",

	Get: func(conn *ssh.Connection) string {
		return conn.Timeout.String()
	},

	Set: func(val string, conn *ssh.Connection) error {
		timeout, err := strconv.Atoi(val)
		if err != nil {
			return fmt.Errorf("the value of timout must a valid positive interger")
		}

		conn.Timeout = time.Second * time.Duration(timeout)

		return nil
	},

	Validate: func(conn *ssh.Connection) bool {
		return conn.Timeout != 0
	},
}
