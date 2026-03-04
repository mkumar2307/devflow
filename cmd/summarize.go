/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/mkumar2307/devflow/internal/parser"
	"github.com/mkumar2307/devflow/internal/scanner"

	"github.com/spf13/cobra"
)

// summarizeCmd represents the summarize command
var summarizeCmd = &cobra.Command{
	Use:   "summarize",
	Short: "Summarize Go project structure",

	RunE: func(cmd *cobra.Command, args []string) error {

		files, err := scanner.ScanRepo(".")
		if err != nil {
			return err
		}

		for _, file := range files {
			summary, err := parser.ParseFile(file)
			if err != nil {
				continue
			}

			fmt.Printf("\n📄 %s\n", summary.Path)
			fmt.Println("  Functions:", summary.Functions)
			fmt.Println("  Structs:", summary.Structs)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(summarizeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// summarizeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// summarizeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
