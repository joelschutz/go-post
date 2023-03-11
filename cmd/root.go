/*
Copyright © 2023 JOEL SCHUTZ <JOELSSCHUTZ@YAHOO.COM.BR>

*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/joelschutz/go-post/internal"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-post",
	Short: "A brief application to extract selected comments as Markdown",
	Long: `A brief application to extract selected comments as Markdown.
	Check repository: https://github.com/joelschutz/go-post`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("parse called", args)

		p, err := internal.NewMDParserFromFile(args[0])
		if err != nil {
			log.Fatal(err)
		}
		fname := args[0] + ".md"

		ioutil.WriteFile(fname, []byte(p.Flush(args[0])), 0777)
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
}
