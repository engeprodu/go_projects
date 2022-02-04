package iteracao

import "testing"

func TestRepetir(t *testing.T) {
	repeticoes := Repetir("a")
	esperado := "aaaaaaaaaaaaaaaaaaaaa"

	if repeticoes != esperado {
		t.Errorf("era esperado que '%s' no entanto o que se conseguiu foi '%s'", esperado, repeticoes)
	}
}

func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetir("a")
	}
}

