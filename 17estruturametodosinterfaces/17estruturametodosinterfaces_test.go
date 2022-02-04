// Gostaria de reforçar o quão grandioso o compilador é. É muito importante ter tempo para ler lentamente as mensagens de erro que você recebe, pois isso te ajudará a longo prazo.package main
package main

import "testing"

func TestPerimetro(t *testing.T) {
    retangulo := Retangulo{10.0, 10.0}
    resultado := Perimetro(retangulo)
    esperado := 40.0

    if resultado != esperado {
        t.Errorf("resultado %.2f, esperado %.2f", resultado, esperado)
    }
}

func TestArea(t *testing.T) {
    testesArea := []struct {
        forma    Forma
        esperado float64
    }{
        {Retangulo{12, 6}, 72.0},
        {Circulo{10}, 314.1592653589793},
        {Triangulo{12, 6}, 36.0},
    }

    for _, tt := range testesArea {
        resultado := tt.forma.Area()
        if resultado != tt.esperado {
            t.Errorf("resultado %.2f, esperado %.2f", resultado, tt.esperado)
        }
    }
}

//comentario de teste//comentario de teste

//comentario de teste//comentario de teste

