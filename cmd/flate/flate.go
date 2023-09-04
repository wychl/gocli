/*
Copyright © 2023 Abner Wanyan <abner.wanyan@gmail.com>
*/
package flatcmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wychl/gocli/cmd"
)

// flateCmd represents the flate command
var flateCmd = &cobra.Command{
	Use:   "flate",
	Short: "flate解压缩",
	Long:  `flate解压缩`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("flate called")
	},
}

func init() {
	cmd.Register(flateCmd)
}
