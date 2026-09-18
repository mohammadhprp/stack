package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"

	"github.com/mohammadhprp/stack/cmd"
	"github.com/mohammadhprp/stack/internal/install"
)

// go:embed patterns may not use "..", so the embed lives in the root package
// and framework/ stays the content root.
//
//go:embed all:framework
var frameworkFS embed.FS

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "stack:", err)
		os.Exit(1)
	}
}

func run() error {
	sub, err := fs.Sub(frameworkFS, "framework")
	if err != nil {
		return err
	}
	cat, err := install.Load(sub)
	if err != nil {
		return err
	}
	return cmd.NewRootCommand(cat).Execute()
}
