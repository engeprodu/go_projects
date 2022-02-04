package main

// // Perimetro retorna o perímetro de um retângulo
// func Perimetro(largura float64, altura float64) float64 {
// 	return 2* (largura + altura)
// }

// func Area(largura float64, altura float64) float64 {
// 	return largura * altura
// }
type Retangulo struct {
    Largura  float64
    Altura float64
}

func (r Retangulo) Area() float64  {
    return 0
}

type Circulo struct {
    Raio float64
}

func (c Circulo) Area() float64  {
    return 0
}