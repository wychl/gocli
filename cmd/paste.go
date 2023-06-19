/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

// pasteCmd represents the past command
var pasteCmd = &cobra.Command{
	Use:   "paste",
	Short: "粘贴",
	Long:  `粘贴`,
	Run: func(cmd *cobra.Command, args []string) {
		value, err := clipboard.ReadAll()
		if err != nil {
			log.Fatalln(err)
		}
		_, err = os.Stdout.Write([]byte(value))
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(pasteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pastCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pastCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
