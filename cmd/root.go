/*
Copyright © 2023 Abner Wanyan <abner.wanyan@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	flatcmd "github.com/wychl/gocli/cmd/flate"
	urlcmd "github.com/wychl/gocli/cmd/url"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gocli",
	Short: "gocli 是一个实用、简单和便捷的工具",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(flatcmd.FlateCmd)
	rootCmd.AddCommand(urlcmd.URLCmd)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
