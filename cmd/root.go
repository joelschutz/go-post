/*
Copyright © 2023 JOEL SCHUTZ <JOELSSCHUTZ@YAHOO.COM.BR>

*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/joelschutz/go-post/internal"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-post [path to file]",
	Short: "A brief application to extract selected comments as Markdown",
	Long: `A brief application to extract selected comments as Markdown.
	Check repository: https://github.com/joelschutz/go-post`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("parse called", args)

		tag, _ := cmd.Flags().GetString("tag")

		for _, path := range args {
			if strings.HasSuffix(path, "...") {
				return internal.ParseDir(strings.TrimSuffix(path, "..."), tag)
			} else {
				return internal.ParseFile(path, tag)
			}
		}
		return nil
	},
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
	rootCmd.PersistentFlags().StringP("tag", "t", "all", "Filters out comments tags different from the specified")
}
