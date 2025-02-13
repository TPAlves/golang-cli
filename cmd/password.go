/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"

	"cli-devops/utils"
)

// passwordCmd represents the password command
var passwordCmd = &cobra.Command{
	Use:   "password",
	Short: "Gera senha aleatória",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		size, _ := cmd.Flags().GetInt("total")
		cmd.Println(utils.GeneretePassword(size))
	},
}

func init() {
	rootCmd.AddCommand(passwordCmd)
	passwordCmd.PersistentFlags().IntP("total", "t", 16, "Tamanho da senha a ser gerada.")
}
