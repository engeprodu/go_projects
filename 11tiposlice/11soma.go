package main

// Soma calcula o valor total dos números em um array

// func Soma(numeros [5]int) int {
// 	soma := 0
// 	for i := 0; i < 5; i++ {
// 		soma += numeros[i]
// 	}
// 	return soma
// }

// func Soma(numeros [5]int) int {
//     soma := 0
//     for _, numero := range numeros {
//         soma += numero
//     }
//     return soma
// }

func Soma(numeros []int) int {
    soma := 0
    for _, numero := range numeros {
        soma += numero
    }
    return soma
}