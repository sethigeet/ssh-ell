package shell

import (
	"fmt"
	"os"
	"strings"
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
}
