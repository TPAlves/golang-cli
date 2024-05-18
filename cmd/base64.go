/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cli-devops/utils"

	"github.com/spf13/cobra"
)

// base64Cmd represents the base64 command
var base64Cmd = &cobra.Command{
	Use:   "base64",
	Short: "Realiza o encode e decode de strings",
	Long: `Realiza o encode e decode de strings para base64. 
	Exemplo de uso:
		encode: devops-cli --encode "texto"
		decode: devops-cli --decode "texto"
	`,
	Run: func(cmd *cobra.Command, args []string) {
		encodeString, _ := cmd.Flags().GetString("encode")
		decodeString, _ := cmd.Flags().GetString("decode")

		if encodeString != "" {
			encode := utils.EncodeString(encodeString)
			cmd.Println(encode)
		} else if decodeString != "" {
			decode := utils.DecodeString(decodeString)
			cmd.Println(decode)
		}
	},
}

func init() {
	rootCmd.AddCommand(base64Cmd)
	base64Cmd.PersistentFlags().StringP("encode", "e", "", "Encode do texto")
	base64Cmd.PersistentFlags().StringP("decode", "d", "", "Decode do texto")
}
