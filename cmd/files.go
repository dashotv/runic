/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/dashotv/runic/static"
)

// filesCmd represents the files command
var filesCmd = &cobra.Command{
	Use:   "files",
	Short: "list embedded files",
	Long:  "list embedded files",
	Run: func(cmd *cobra.Command, args []string) {
		fs := static.FS
		entries, err := fs.ReadDir(".")
		if err != nil {
			log.Fatal("error reding fs dir: ", err)
		}
		for _, entry := range entries {
			fmt.Println(".", entry.Name())
		}
	},
}

func init() {
	rootCmd.AddCommand(filesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// filesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// filesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
