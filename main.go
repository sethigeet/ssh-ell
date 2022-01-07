package main

import (
	"os"

	"github.com/fatih/color"

	"github.com/sethigeet/ssh-ell/shell"
)

func main() {
	err := shell.Setup()
	if err != nil {
		color.New(color.FgRed, color.Bold).Fprintln(os.Stderr, "An error occurred!")
		color.New(color.FgRed).Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	shell.Run()
}
