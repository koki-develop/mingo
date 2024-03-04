package cmd

import (
	"fmt"
	"os"

	"github.com/koki-develop/mingo/internal/mingo"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "mingo",
	RunE: func(cmd *cobra.Command, args []string) error {
		for i, file := range args {
			if i > 0 {
				fmt.Println()
			}

			min, err := mingo.MinifyFile(file)
			if err != nil {
				return err
			}

			fmt.Print(min)
		}

		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
