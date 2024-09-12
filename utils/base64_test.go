package utils

import "testing"

func TestEncode(t *testing.T) {
	resultado := EncodeString("word")
	esperado := "d29yZA=="

	if esperado != resultado {
		t.Errorf("esperado %s, resultado %s", esperado, resultado)
	}
}

func TestDecode(t *testing.T) {
	resultado := DecodeString("d29yZA==")
	esperado := "word"

	if esperado != resultado {
		t.Errorf("esperado %s, resultado %s", esperado, resultado)
	}
}