package ssh

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"syscall"

	"github.com/fatih/color"
	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
)

// GetAuthMethod gets the required auth method for Connection to use to connect
// to the host
func GetAuthMethod(methodType string, rest ...string) (ssh.AuthMethod, error) {
	switch methodType {
	case "password":
		return getPassowrdAuth()
	case "keypair":
		return getKeypairAuth(rest...)
	default:
		return nil, fmt.Errorf("invalid method type provided: %s", methodType)
	}
}

func getPassowrdAuth() (ssh.AuthMethod, error) {
	color.New(color.FgCyan, color.Bold).Printf("\nEnter the passowrd: ")
	passwd, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return nil, fmt.Errorf("unable to read the password entered: %v", err)
	}
	fmt.Printf("\n")

	return ssh.Password(string(passwd)), nil
}

func getKeypairAuth(args ...string) (ssh.AuthMethod, error) {
	// Make sure the path to the identity file is passed
	if len(args) != 1 {
		panic("invalid args provided, missing exactly one identity file path arg")
	}

	// Get the absolute path of the file
	identityFilePath, err := filepath.Abs(args[0])
	if err != nil {
		return nil, fmt.Errorf("unable to read private key: %v", err)
	}
	// Read the file
	key, err := ioutil.ReadFile(identityFilePath)
	if err != nil {
		return nil, fmt.Errorf("unable to read private key: %v", err)
	}

	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, fmt.Errorf("unable to parse private key: %v", err)
	}

	return ssh.PublicKeys(signer), nil
}
