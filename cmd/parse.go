/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"go/parser"
	"go/token"
	"io/ioutil"

	"github.com/joelschutz/go-post/internal"
	"github.com/spf13/cobra"
)

// parseCmd represents the parse command
var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("parse called", args)
		fs := token.NewFileSet()
		f, err := parser.ParseFile(fs, args[0], nil, parser.ParseComments)
		if err != nil {
			fmt.Println(err)
		}
		buf, err := ioutil.ReadFile(args[0])
		if err != nil {
			fmt.Println(err)
		}

		p := internal.MDParser{File: buf}
		p.ParseComments(f.Comments)
		p.ParseDeclarations(f.Decls)
		fname := args[0] + ".md"
		buf, err = ioutil.ReadAll(p.Flush())
		if err != nil {
			fmt.Println(err)
		}

		ioutil.WriteFile(fname, buf, 777)
	},
}

func init() {
	rootCmd.AddCommand(parseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// parseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// parseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
