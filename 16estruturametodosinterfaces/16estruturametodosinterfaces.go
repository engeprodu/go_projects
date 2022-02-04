package main

// // Perimetro retorna o perímetro de um retângulo
// func Perimetro(largura float64, altura float64) float64 {
// 	return 2* (largura + altura)
// }

// func Area(largura float64, altura float64) float64 {
// 	return largura * altura
// }
type Retangulo struct {
    Largura float64
    Altura  float64
}

func Perimetro(retangulo Retangulo) float64 {
    return 2 * (retangulo.Largura + retangulo.Altura)
}

func Area(retangulo Retangulo) float64 {
    return retangulo.Largura * retangulo.Altura
}

type Circulo struct {
    Raio float64
}

// Area retorna a área de um círculo
func (c Circulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}