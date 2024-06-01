package utils

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"
)

func GeneretePassword(size int) string {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	if size < 8 {
		fmt.Println("O tamanho mínimo é de 8 caracteres")
		return ""
	}
	b := make([]byte, size)
	if _, err := r.Read(b); err != nil {
		fmt.Println("Erro ao gerar senha")
		return ""
	}
	return base64.StdEncoding.EncodeToString(b)
}