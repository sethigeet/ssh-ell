package shell

import (
	"fmt"
	"os"
	"strings"

	"github.com/sethigeet/ssh-ell/shell/commands"
)

func executor(s string) {
	s = strings.TrimSpace(s)
	switch s {
	case "":
		return
	case "quit", "exit":
		fmt.Println("Bye!")
		os.Exit(0)
		return
	}

	commands.Execute(s)
}
