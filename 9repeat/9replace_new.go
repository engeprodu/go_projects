package main

import (
	"fmt"
	"strings"
)

const a = "Todos os dias quando acordo, não tenho mais o tempo que passou mas temos muito tempo, temos todo tempo do mundo..."

func main() {
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 1))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", 1))
	fmt.Println(strings.Replace(a, "a", "i", -1))
	fmt.Println(a)
}