package ssh

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/fatih/color"
	"github.com/sethigeet/ssh-ell/utils"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
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

// GetHostKeyCallback returns a new ssh.HostKeyCallback which will ask to user
// to add the key to `known_host` file or error out appropriately
func GetHostKeyCallback(conn *Connection) (func(string, net.Addr, ssh.PublicKey) error, error) {
	host := fmt.Sprintf("%s:%d", conn.Host, conn.Port)
	normHost := knownhosts.Normalize(host)
	return func(_ string, _ net.Addr, key ssh.PublicKey) error {
		khFilepath := filepath.Join(os.Getenv("HOME"), ".ssh", "known_hosts")
		file, err := os.Open(khFilepath)
		if err != nil {
			return err
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		var hostKey ssh.PublicKey
		for scanner.Scan() {
			fields := strings.Split(scanner.Text(), " ")
			if len(fields) != 3 {
				continue
			}
			if !strings.Contains(fields[0], normHost) {
				continue
			}

			var err error
			hostKey, _, _, _, err = ssh.ParseAuthorizedKey(scanner.Bytes())
			if err != nil {
				return fmt.Errorf("error parsing %q: %v", fields[2], err)
			}

			// make sure both the keys are of the same type
			if hostKey.Type() != key.Type() {
				hostKey = nil
				continue
			}

			// we have found the required key!
			break
		}

		if hostKey == nil {
			// key for host is not known
			color.New(color.FgRed, color.Bold).Println("### Warning ###")
			color.New(color.FgRed).Printf(`The authenticity of host %q can't be established.
%s key fingerprint is %s.
`,
				normHost,
				key.Type(),
				ssh.FingerprintSHA256(key),
			)
			if !utils.YesNoInput("Are you sure you want to continue connecting?", prompt.OptionPrefixTextColor(prompt.Cyan)) {
				return fmt.Errorf("aborted by user")
			}

			toWrite := fmt.Sprintf("%s %s", normHost, ssh.MarshalAuthorizedKey(key))
			err := utils.AppendToFile(khFilepath, []byte(toWrite), 0600)
			if err != nil {
				return err
			}
			color.New(color.FgYellow, color.Bold).Printf("Warning: Permanently added %q (%s) to the list of known hosts.\n", normHost, key.Type())
			return nil
		}

		if !bytes.Equal(hostKey.Marshal(), key.Marshal()) {
			// keys did not match
			color.New(color.FgRed, color.Bold).Fprintln(os.Stderr, "### WARNING: REMOTE HOST IDENTIFICATION HAS CHANGED! ###")
			color.New(color.FgRed).Fprintf(os.Stderr, `IT IS POSSIBLE THAT SOMEONE IS DOING SOMETHING NASTY!
Someone could be eavesdropping on you right now (man-in-the-middle attack)!
It is also possible that a host key has just been changed.
The fingerprint for the %s key sent by the remote host is %s.
Please contact your system administrator.
Add correct host key in ~/.ssh/known_hosts to get rid of this message.

`, key.Type(), ssh.FingerprintSHA256(key))
			return fmt.Errorf("hostkey for %q(%s) did not match", normHost, key.Type())
		}

		// keys matched
		return nil
	}, nil
}
