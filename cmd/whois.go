/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"

	"cli-devops/utils"
)

// whoisCmd represents the whois command
var whoisCmd = &cobra.Command{
	Use:   "whois",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		domain := cmd.Flag("domain").Value.String()
		if domain == "" {
			cmd.Println("Informe o domínio a ser consultado no whois")
			return
		}
		cmd.Println(utils.GetWhois(domain))
	},
}

func init() {
	rootCmd.AddCommand(whoisCmd)
	whoisCmd.PersistentFlags().StringP("domain", "d", "", "Domínio a ser consultado no whois")
}
