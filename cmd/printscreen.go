/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cli-devops/utils"

	"github.com/spf13/cobra"
)

// printscreenCmd represents the printscreen command
var printscreenCmd = &cobra.Command{
	Use:   "printscreen",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		site := cmd.Flag("url").Value.String()
		utils.GetChromeScreenShot(site, 100)
	},
}

func init() {
	rootCmd.AddCommand(printscreenCmd)
	printscreenCmd.PersistentFlags().StringP("url", "u", "", "Executa o printscreen da tela do site informado.")
}
