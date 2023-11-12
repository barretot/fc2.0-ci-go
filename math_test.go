package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(15, 35)

	if total != 50 {
		t.Errorf("Resultado da soma é inválido: Resultado %d. Esperado: %d", total, 45)
	}
}
