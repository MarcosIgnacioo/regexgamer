package main

// # ([a-zA-Z]+[\d]*?#?\.?[a-zA-Z]*)*\*?<?>?=?((‘?’?'?)*([\d]*[a-zA-Z]*)*(‘?’?'?)*)?
import (
	"fmt"
	"regexp"
	"strings"

	"gamer.com/go/numbers"
)

func main() {
	TestRegexNewLines()
}

func TestRegexNewLines() {
	inputLn := `Estos datos positivos también tienen su correlación con la salud. La tasa de mortalidad \n
entre los recién nacidos era de 648 por cada 1.000 en 1990. Pero en 2016, la realidad \n
era bien diferente, ya que solo 30,5 fallecían por cada 1.000 (en 26 años se ha reducido un \n
53%). Igual de positivo es la caída de la mortalidad entre los menores de 5A; en 26 años se \n
ha pasado de una tasa de 93,4 fallecidos por 1.0E2 a un 40,8 (una reducción del 56%). \n
De igual manera, el número de mujeres que fallece durante el parto también ha decrecido \n
(en 1990 el número de muertes fue de 530.000 y en 2015 de 303.00 una disminución del \n
43%).`
	lines := strings.Split(inputLn, `\n`)
	fmt.Println(inputLn)
	r, _ := regexp.Compile(`\d+\,?\.?\d+[a-zA-Z]?\d*%?(\.?\,?\d+)*`)
	n := 1
	for i, line := range lines {
		nums := r.FindAll([]byte(line), -1)
		for _, number := range nums {
			fmt.Println(numbers.NewNumber(string(number), n, i+1).String())
			n++
		}
	}
}

// func TestRegex() {

// 	input := `Estos datos positivos también tienen su correlación con la salud. La tasa de mortalidad
// entre los recién nacidos era de 64,8 por cada 1.000 en 1990. Pero en 2016, la realidad
// era bien diferente, ya que solo 30,5B fallecían por cada 1.000 (en 26 años se ha reducido un
// 53%). Igual de positivo es la caída de la mortalidad entre los menores de 5A; en 26 años se
// ha pasado de una tasa de 93,4 fallecidos por 1.0E2 a un 40,8 (una reducción del 56%).
// De igual manera, el número de mujeres que fallece durante el parto también ha decrecido
// (en 1990 el número de muertes fue de 530.000 y en 2015 de 303.00.0, una disminución del
// 43%).`
// 	r, _ := regexp.Compile(`\d*\,?\.?\d+[a-zA-Z]?\d*%?(\.\d+)*`)
// 	nums := r.FindAll([]byte(input), -1)
// 	for i, v := range nums {
// 		fmt.Println(numbers.NewNumber(string(v), i+1, 0).String())
// 	}
// }

// func test() {
// 	//\\d*\\.?\\d+%?

// 	// Tambien podria intentar hacer un select mas generico aqui y agarrar numeros aunque tengan letras dentro, ya luego en el package numbers los filtro
// 	// r, _ := regexp.Compile(`\d*\d?\.*\,*\d+[a-zA-Z]*\d*\,?\.?\d*%?`)
// 	r, _ := regexp.Compile(`a`)
// 	test_natural := r.FindAll([]byte("abc"), -1)

// 	for _, v := range test_natural {
// 		fmt.Println(string(v))
// 	}
// }

// // test masomenos \d*[E0-9]?[a-zA-Z]*\d*[[[\.\,]*\d*]*\%?.?\d*]?
// // hasta agregamos uno mas que no viene pero deberia de salir tmb la neta
// // version optimizada del que si jala
// // \d*\,?\.?\d+[a-zA-Z]?\d*%?(\.\d+)*
// // Este ya jala al 100% con el input de la tarea
// // r, _ := regexp.Compile("\\d+\\,?\\.?\\d+[a-zA-Z]?\\d?%?(\\.\\d+)*")
