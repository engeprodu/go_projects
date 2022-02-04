package main

import "testing"

func TestOla(t *testing.T) {
//     Declarando variáveis
// Estamos declarando algumas variáveis com a sintaxe nomeDaVariavel := valor, que nos permite reutilizar alguns valores nos nossos testes de maneira legível.
    resultado := Ola()
    esperado := "Olá, mundooo"

    if resultado != esperado {
        // t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
        // t.Error("Expected:", resultado, "Got:", teste)
        t.Error("Expected:", resultado, "Got:", esperado)

//         t.Errorf
// Estamos chamando o método Errorf em nosso t que irá imprimir uma mensagem e falhar o teste. O sufixo f no final de Errorf representa que podemos formatar e montar uma string com valores inseridos dentro de valores de preenchimentos %s. Quando fazemos um teste falhar, devemos ser bastante claros com o que aconteceu.
// Iremos explorar a diferença entre métodos e funções depois.
    }
}