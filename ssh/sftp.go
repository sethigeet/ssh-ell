package ssh

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/pkg/sftp"
	"github.com/schollz/progressbar/v3"
)

func (c *Connection) Upload(srcPath, dstPath string) error {
	// open an SFTP session over an existing ssh connection.
	sftp, err := sftp.NewClient(c.sshClient)
	if err != nil {
		return err
	}
	defer sftp.Close()

	// Open the source file
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// Create the destination file
	dstFile, err := sftp.Create(dstPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	info, err := srcFile.Stat()
	if err != nil {
		return err
	}

	// create a progressbar to show the progress
	bar := progressbar.NewOptions64(
		info.Size(),
		progressbar.OptionSetDescription("Uploading"),
		progressbar.OptionSetWriter(os.Stderr),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetWidth(10),
		progressbar.OptionThrottle(65*time.Millisecond),
		progressbar.OptionShowCount(),
		progressbar.OptionOnCompletion(func() {
			fmt.Printf("\n")
		}),
		progressbar.OptionSpinnerType(14),
		progressbar.OptionFullWidth(),
	)
	bar.RenderBlank()

	// write to file
	_, err = io.Copy(io.MultiWriter(dstFile, bar), srcFile)
	if err != nil {
		return err
	}

	return nil
}

func (c *Connection) Download(srcPath, dstPath string) error {
	// open an SFTP session over an existing ssh connection.
	sftp, err := sftp.NewClient(c.sshClient)
	if err != nil {
		return err
	}
	defer sftp.Close()

	// Open the source file
	srcFile, err := sftp.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// Create the destination file
	dstFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	info, err := srcFile.Stat()
	if err != nil {
		return err
	}

	// create a progressbar to show the progress
	bar := progressbar.NewOptions64(
		info.Size(),
		progressbar.OptionSetDescription("Downloading"),
		progressbar.OptionSetWriter(os.Stderr),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetWidth(10),
		progressbar.OptionThrottle(65*time.Millisecond),
		progressbar.OptionShowCount(),
		progressbar.OptionOnCompletion(func() {
			fmt.Printf("\n")
		}),
		progressbar.OptionSpinnerType(14),
		progressbar.OptionFullWidth(),
	)
	bar.RenderBlank()

	// write to file
	_, err = io.Copy(io.MultiWriter(dstFile, bar), srcFile)
	if err != nil {
		return err
	}

	return nil
}
