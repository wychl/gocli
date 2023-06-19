/*
Copyright © 2023 Abner Wanyan <abner.wanyan@gmail.com>
*/
package flatcmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// flateCmd represents the flate command
var FlateCmd = &cobra.Command{
	Use:   "flate",
	Short: "flate解压缩",
	Long:  `flate解压缩`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("flate called")
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// flateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// flateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
