package main

// Soma calcula o valor total dos números em um slice
func Soma(numeros []int) int {
	soma := 0
	for _, numero := range numeros {
		soma += numero
	}
	return soma
}


// func SomaTudo(numerosParaSomar ...[]int) (somas []int) {
//     return
// }

// SomaTudo calcula as respectivas somas de cada slice recebido
// func SomaTudo(numerosParaSomar ...[]int) []int {
// 	quantidadeDeNumeros := len(numerosParaSomar)
// 	somas := make([]int, quantidadeDeNumeros)

// 	for i, numeros := range numerosParaSomar {
// 		somas[i] = Soma(numeros)
// 	}

// 	return somas
//     // return somas
// }


func SomaTudo(numerosParaSomar ...[]int) []int {
    var somas []int
    for _, numeros := range numerosParaSomar {
        somas = append(somas, Soma(numeros))
    }

    return somas
	
}

func SomaTodoOResto(numerosParaSomar ...[]int) []int {
    var somas []int
    for _, numeros := range numerosParaSomar {
        final := numeros[1:]
        somas = append(somas, Soma(final))
    }

    return somas
}

