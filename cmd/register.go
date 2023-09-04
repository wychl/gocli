package cmd

import "github.com/spf13/cobra"

func Register(cmds ...*cobra.Command) {
	rootCmd.AddCommand(cmds...)
}
