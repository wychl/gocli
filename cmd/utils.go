package cmd

import (
	"io"
	"os"

	"github.com/spf13/cobra"
)

func GetCmdParam(cmd *cobra.Command, args []string) ([]byte, error) {
	if len(args) > 0 {
		return []byte(args[0]), nil
	}
	return io.ReadAll(os.Stdin)
}
