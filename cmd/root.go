/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/i582/cfmt/cmd/cfmt"
	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cli-devops",
	Short: "Aplicação de DevOps 🚀",
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli-devops.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	cfmt.RegisterStyle("code", func(s string) string {
		return cfmt.Sprintf("{{%s}}::red|underline", s)
	})

	cfmt.Printf(`
{{+-------------------------------------------+}}::bold
{{|       🚀🚀🚀 DevOps CLI 🚀🚀🚀            |}}::bold
{{+-------------------------------------------+}}::bold
{{100}}::#ffffff myStyle := color.{{New(color.FgWhite, color.BgBlack, color.OpBold)}}::code|bold
{{[100, 17]}}::blue Undefined function New at {{~/projects/test}}::underline:100

{{101}}::#ffffff {{myStyle}}::code.Print("t")
{{[101, 0]}}::blue Undefined variable myStyle at {{~/projects/test}}::underline:101

Para mais informações visite: https://www.google.com.br/
	`)
}
