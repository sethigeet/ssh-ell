package ssh

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"golang.org/x/crypto/ssh"
)

// GetAuthMethod gets the required auth method for Connection to use to connect
// to the host
func GetAuthMethod(methodType string, rest ...string) (ssh.AuthMethod, error) {
	switch methodType {
	case "password":
		return getPassowrdAuth(rest...)
	case "key":
		return getKeyAuth(rest...)
	default:
		panic(fmt.Sprintf("invalid method type provided: %s", methodType))
	}
}

func getPassowrdAuth(args ...string) (ssh.AuthMethod, error) {
	// Make sure the password is provided
	if len(args) != 1 {
		panic("invalid args provided, missing exactly one password arg")
	}

	return ssh.Password(string(args[0])), nil
}

func getKeyAuth(args ...string) (ssh.AuthMethod, error) {
	// Make sure the path to the identity file is passed
	if len(args) != 1 {
		panic("invalid args provided, missing exactly one identity file path arg")
	}

	// Get the absolute path of the file
	identityFilePath := args[0]
	if strings.HasPrefix(identityFilePath, "~") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}
		identityFilePath = strings.Replace(identityFilePath, "~", homeDir, 1)
	}
	identityFilePath, err := filepath.Abs(identityFilePath)
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

	color.New(color.FgGreen, color.Bold).Printf("Successfully loaded private key from %q\n", identityFilePath)
	return ssh.PublicKeys(signer), nil
}
