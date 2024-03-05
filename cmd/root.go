package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/koki-develop/mingo/internal/mingo"
	"github.com/spf13/cobra"
)

var (
	flagWrite bool
)

var rootCmd = &cobra.Command{
	Use: "mingo",
	RunE: func(cmd *cobra.Command, args []string) error {
		for i, file := range args {
			src, err := os.ReadFile(file)
			if err != nil {
				return err
			}

			var out io.Writer
			if flagWrite {
				f, err := os.Create(file)
				if err != nil {
					return err
				}
				defer f.Close()

				out = f
			} else {
				out = os.Stdout
				if i > 0 {
					fmt.Fprintln(out)
				}
			}

			min, err := mingo.Minify(file, src)
			if err != nil {
				return err
			}

			fmt.Fprint(out, min)
		}

		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&flagWrite, "write", "w", false, "write result to (source) file instead of stdout")
}
