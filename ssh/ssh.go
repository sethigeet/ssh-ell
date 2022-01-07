package ssh

import (
	"fmt"
	"os/user"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type Connection struct {
	Host             string
	Port             uint16
	User             string
	Timeout          time.Duration
	AuthMethod       ssh.AuthMethod
	IdentityFilePath string

	AuthMethodCommonName string

	sshClient  *ssh.Client
	sftpClient *sftp.Client
}

func (c *Connection) Connect() error {
	sshClient, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", c.Host, c.Port), &ssh.ClientConfig{
		User:    c.User,
		Timeout: c.Timeout,
		Auth:    []ssh.AuthMethod{c.AuthMethod},
	})
	if err != nil {
		return err
	}

	c.sshClient = sshClient

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

	return nil
}
