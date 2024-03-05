package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/koki-develop/mingo/internal/mingo"
	"github.com/spf13/cobra"
)

var (
	flagWrite bool
)

var rootCmd = &cobra.Command{
	Use:   "mingo [flags] [files]...",
	Short: "Go language also wants to be minified",
	Long:  "Go language also wants to be minified.",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		for i, file := range args {
			err := filepath.WalkDir(file, func(path string, d os.DirEntry, err error) error {
				if err != nil {
					return err
				}

				if d.IsDir() {
					return nil
				}

				if filepath.Ext(path) != ".go" {
					return nil
				}

				src, err := os.ReadFile(path)
				if err != nil {
					return err
				}

				var out io.Writer
				if flagWrite {
					f, err := os.Create(path)
					if err != nil {
						return err
					}
					defer f.Close()

					out = f
				} else {
					out = os.Stdout
					if i > 0 {
						if _, err := fmt.Fprintln(out); err != nil {
							return err
						}
					}
				}

				min, err := mingo.Minify(path, src)
				if err != nil {
					return err
				}

				if _, err := fmt.Fprint(out, min); err != nil {
					return err
				}

				return nil
			})
			if err != nil {
				return err
			}
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
