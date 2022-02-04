package main

import "testing"
//nós criamos um arquivo _test
//nós criamos uma função que comece com Test
func TestSoma(t *testing.T) {
    teste := soma(3, 2, 1)
// roda a funçção
// estima um resultado    
    resultado := 6
    if teste != resultado {
        t.Error("Expected:", resultado, "Got:", teste)
        // se der 6, espero um resultado-resultado, e eu recebi, aquele valor
    }

}

func TestSoma2(t *testing.T) {
    teste := soma(3, 2, 2)
    resultado := 7
    if teste != resultado {
        t.Error("Expected:", resultado, "Got:", teste)
        // se der 6, espero um resultado-resultado, e eu recebi, aquele valor
    }

}