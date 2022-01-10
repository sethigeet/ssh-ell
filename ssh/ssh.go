package ssh

import (
	"fmt"
	"os"
	"os/user"
	"time"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

type Connection struct {
	Host             string
	Port             uint16
	User             string
	Timeout          time.Duration
	AuthMethod       ssh.AuthMethod
	IdentityFilePath string
	TerminalName     string

	AuthMethodCommonName string
	Connected            bool

	sshClient *ssh.Client
}

func (c *Connection) Connect() error {
	hostKeyCallback, err := GetHostKeyCallback(c)
	if err != nil {
		return err
	}
	sshClient, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", c.Host, c.Port), &ssh.ClientConfig{
		User:    c.User,
		Timeout: c.Timeout,
		Auth:    []ssh.AuthMethod{c.AuthMethod},

		HostKeyCallback: hostKeyCallback,
	})
	if err != nil {
		return err
	}

	c.sshClient = sshClient
	c.Connected = true

	go func(conn *Connection) {
		conn.sshClient.Wait()

		conn.Connected = false
	}(c)

	return nil
}

func (c *Connection) Disconnect() error {
	return c.sshClient.Close()
}

func (c *Connection) ApplyDefaults() error {
	curr, err := user.Current()
	if err != nil {
		return err
	}
	c.User = curr.Username

	c.Port = 22

	c.Timeout = time.Millisecond * 5000

	c.TerminalName = "xterm-256color"

	return nil
}

func (c *Connection) Shell() error {
	sess, err := c.sshClient.NewSession()
	if err != nil {
		return err
	}
	defer sess.Close()

	// set the pipes
	sess.Stdin = os.Stdin
	sess.Stdout = os.Stdout
	sess.Stderr = os.Stderr

	// request a psedo tty
	width, height, err := terminal.GetSize(0)
	if err != nil {
		return err
	}
	if err := sess.RequestPty(c.TerminalName, height, width, ssh.TerminalModes{
		ssh.ECHO:          0,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}); err != nil {
		return err
	}

	if err := sess.Shell(); err != nil {
		return err
	}

	if err := sess.Wait(); err != nil {
		return err
	}

	return nil
}
