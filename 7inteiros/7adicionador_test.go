package inteiros


/* Observe também que não estamos mais usando o pacote main, em vez disso, definimos um pacote chamado inteiros, pois o nome sugere que ele agrupará funções para trabalhar com números inteiros, como Adiciona. */

import "testing"
import "fmt"

/* Quando você tem mais de um argumento do mesmo tipo (no nosso caso dois inteiros) ao invés de ter (x int, y int) você pode encurtá-lo para (x, y int). */
func Adiciona(x, y int) int {
    return x + y
}

func ExampleAdiciona() {
    soma := Adiciona(1, 5)
    fmt.Println(soma)
    // Output: 6
}

func TestAdicionador(t *testing.T) {
    soma := Adiciona(2, 2)
    esperado := 4

    if soma != esperado {
        t.Errorf("esperado '%d', resultado '%d'", esperado, soma)
		/* Você deve ter notado que estamos usando %d como string de formatação, em vez de %s. Isso porque queremos que ele imprima um valor inteiro e não uma string */
    }
	
}

