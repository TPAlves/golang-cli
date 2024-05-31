/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cli-devops/utils"
	"strings"

	"github.com/spf13/cobra"
)

// checkportCmd represents the checkport command
var checkportCmd = &cobra.Command{
	Use:   "checkport",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		host, _ := cmd.Flags().GetString("host")
		if host == "" {
			cmd.Println("Informe o host a ser validado.")
		}

		ports, _ := cmd.Flags().GetString("ports")
		if ports == "" {
			cmd.Println("Informe a porta/lista de porta a serem validadas.")
		}
		portsList := strings.Split(ports, ",")
		utils.CheckPort(host, portsList)
	},
}

func init() {
	rootCmd.AddCommand(checkportCmd)
	checkportCmd.PersistentFlags().StringP("host", "", "", "Host a ser validado")
	checkportCmd.PersistentFlags().StringP("ports", "p", "", "Lista de portas separadas por vírgula. Ex.: 8080,443,90")
}
