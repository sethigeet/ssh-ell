package commands

import (
	"os"

	"github.com/c-bata/go-prompt"
	"github.com/fatih/color"
	"github.com/sethigeet/ssh-ell/utils"
)

type Download struct{}

func (Download) Complete(d prompt.Document) []prompt.Suggest {
	completer := utils.FilePathCompleter{IgnoreCase: true}

	return completer.Complete(d)
}

func (Download) Execute(args ...string) {
	if len(args) != 2 {
		color.New(color.FgRed, color.Bold).Fprintln(os.Stderr, "Invalid number of args specified")
		color.New(color.FgRed).Fprintln(os.Stderr, "The connect command does not take any args")
		return
	}

	if !conn.Connected {
		color.New(color.FgRed, color.Bold).Fprintf(os.Stderr, "You are not connected to any host!\n")
		return
	}

	srcPath, err := utils.GetAbsPath(args[0])
	if err != nil {
		color.New(color.FgRed, color.Bold).Fprintln(os.Stderr, "Unable to open the source file!")
		color.New(color.FgRed).Fprintf(os.Stderr, "error: %s\n", err)
		return
	}
	dstPath, err := utils.GetAbsPath(args[1])
	if err != nil {
		color.New(color.FgRed, color.Bold).Fprintln(os.Stderr, "Unable to open the destination file!")
		color.New(color.FgRed).Fprintf(os.Stderr, "error: %s\n", err)
		return
	}

	err = conn.Download(srcPath, dstPath)
	if err != nil {
		color.New(color.FgRed, color.Bold).Fprintln(os.Stderr, "Unable to upload the file!")
		color.New(color.FgRed).Fprintf(os.Stderr, "error: %s\n", err)
		return
	}

	color.New(color.FgGreen, color.Bold).Printf("Successfully downloaded %q to %q\n", srcPath, dstPath)
}
