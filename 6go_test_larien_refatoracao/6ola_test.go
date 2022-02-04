package main

import "testing"

// t representa a estrutura que mapeia o teste

func TestOla(t *testing.T) {
    verificaMensagemCorreta := func(t *testing.T, resultado, esperado string) {
        t.Helper()
        // t.Helper () é necessário para dizermos ao conjunto de testes que este é um método auxiliar. Ao fazer isso, quando o teste falhar, o número da linha relatada estará em nossa chamada de função, e não dentro do nosso auxiliar de teste. Isso ajudará outros desenvolvedores a rastrear os problemas com maior facilidade. Se você ainda não entendeu, comente, faça um teste falhar e observe a saída do teste.
        if resultado != esperado {
            t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
        }
    }
// cada t.run trata um cenário
    t.Run("quando diz olá para as pessoas", func(t *testing.T) {
        resultado := Ola("Chris", "")
        esperado := "Olá, Chris" + "!"
        verificaMensagemCorreta(t, resultado, esperado)
    })

    t.Run("quando nao diz nada e isso retorna mundo como opção", func(t *testing.T) {
        resultado := Ola("", "")
        esperado := "Olá, Mundo" + "!"
        // na linha de cima, é correspondente ao que é determinado na linha 13 do código, enquanto variável de string
        verificaMensagemCorreta(t, resultado, esperado)
    })
    
    t.Run("em espanhol", func(t *testing.T) {
        resultado := Ola("Elodie", "espanhol")
        esperado := "Hola, Elodie" + "!"
        verificaMensagemCorreta(t, resultado, esperado)
    })

    t.Run("em francês", func(t *testing.T) {
        resultado := Ola("Elodie", "frances")
        esperado := "Bonjour, Elodie" + "!"
        verificaMensagemCorreta(t, resultado, esperado)
    })

    t.Run("em italiano", func(t *testing.T) {
        resultado := Ola("Elodie", "italiano")
        esperado := "Salve, Elodie" + "!"
        verificaMensagemCorreta(t, resultado, esperado)
    })

    t.Run("em ingles", func(t *testing.T) {
        resultado := Ola("Elodie", "ingles")
        esperado := "Hello, Elodie" + "!"
        verificaMensagemCorreta(t, resultado, esperado)
    })

}

// Escrever um teste
// Compilar o código sem erros
// Rodar o teste, ver o teste falhar e certificar que a mensagem de erro faz sentido
// Escrever a quantidade mínima de código para o teste passar
// Refatorar