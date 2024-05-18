package utils

import (
	"encoding/base64"
	"fmt"
	"os"
)

func EncodeString(s string) string {
	encode := base64.StdEncoding.EncodeToString([]byte(s))
	return encode 
}

func DecodeString(s string) string {
	decode, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		fmt.Println("Ocorreu um erro ao decodificar:", err)
		os.Exit(1)
	}
	return string(decode)
}

