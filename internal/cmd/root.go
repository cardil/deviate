package cmd

import (
	"os"

	"github.com/cardil/deviate/pkg/cli"
	"github.com/cardil/deviate/pkg/metadata"
	"github.com/spf13/cobra"
)

func root(customizers ...Option) *cobra.Command {
	cmd := &cobra.Command{
		Use:     metadata.Name,
		Short:   metadata.Description,
		Version: metadata.Version,
	}
	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)
	opts := &cli.Options{}
	subs := []subcommand{
		update{opts},
	}
	addFlags(cmd, opts)
	for _, sub := range subs {
		cmd.AddCommand(sub.command())
	}

	for _, opt := range customizers {
		opt(cmd)
	}
	return cmd
}

type subcommand interface {
	command() *cobra.Command
}
