// A função fmt.Println é um efeito colateral (que está imprimindo um valor no stdout [saída padrão do terminal]) e a string que estamos enviando para dentro dela é nosso domínio.

package main

import "fmt"

// Constantes devem melhorar a performance da nossa aplicação, assim como evitar que você crie uma string "Ola, " para cada vez que Ola é chamado.

const espanhol = "espanhol"
const frances = "frances"
const italiano = "italiano"
const prefixoOlaPortugues = "Olá, "
const prefixoOlaEspanhol = "Hola, "
const prefixoOlaFrances = "Bonjour, "
const prefixoOlaItaliano = "Salve, "

func Ola(nome string, idioma string) string {
    if nome == "" {
        nome = "Mundo"
    }

    prefixo := prefixoOlaPortugues

    switch idioma {
    case frances:
        prefixo = prefixoOlaFrances
    case espanhol:
        prefixo = prefixoOlaEspanhol
    case italiano:
        prefixo = prefixoOlaItaliano
}

return prefixo + nome + "!"

}

func main() {
    fmt.Println(Ola("Pedro Afonso","frances"))
}

// Criamos uma nova função usando func, mas dessa vez adicionamos outra palavra reservada string na sua definição. Isso significa que essa função terá como retorno uma string (cadeia de caracteres).

// Escrevendo testes
// Escrever um teste é como escrever uma função, com algumas regras:
// Precisa estar em um arquivo com um nome parecido com xxx_test.go
// A função de teste precisa começar com a palavra Test
// A função de teste recebe um único argumento, que é t *testing.T
// Por enquanto é o bastante para saber que o nosso t do tipo *testing.T é a nossa porta de entrada para a ferramenta de testes e assim você poderá utilizar o t.Fail() quando precisar relatar um erro.