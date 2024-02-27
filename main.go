package main

import (
	"fmt"
	"regexp"

	"gamer.com/go/numbers"
)

func main() {
	TestRegex()
}
func test() {
	r, _ := regexp.Compile("(^(foo|bar){1}$)")
	test_illegal_points := r.Match([]byte("foo"))
	test_illegal_commas := r.Match([]byte("bar"))
	test_natural := r.Match([]byte("foobar"))
	fmt.Println(test_illegal_points)
	fmt.Println(test_illegal_commas)
	fmt.Println(test_natural)
}
func TestRegex() {
	input := `13A Estos datos positivos también tienen su correlación con la salud. La tasa de mortalidad
entre los recién nacidos era de 64,8 por cada 1.000 en 1990. Pero en 2016, la realidad
era bien diferente, ya que solo 30,5B fallecían por cada 1.000 (en 26 años se ha reducido un
53%). Igual de positivo es la caída de la mortalidad entre los menores de 5A; en 26 años se
ha pasado de una tasa de 93,4 fallecidos por 1.0E2 a un 40,8 (una reducción del 56%).
De igual manera, el número de mujeres que fallece durante el parto también ha decrecido
(en 1990 el número de muertes fue de 530.000 y en 2015 de 303.00.0, una disminución del
43%).`
	input = `13J3121`
	// Este ya jala al 100% con el input de la tarea
	// r, _ := regexp.Compile("\\d+\\,?\\.?\\d+[a-zA-Z]?\\d?%?(\\.\\d+)*")
	r, _ := regexp.Compile("\\d+\\,?\\.?\\d+[a-zA-Z]?\\d?%?(\\.\\d+)*")
	asdf := r.FindAll([]byte(input), -1)
	for i, v := range asdf {
		fmt.Println(numbers.NewNumber(string(v), i+1).String())
		if i+1 == 2 || i+1 == 6 || i+1 == 11 || i+1 == 15 || i+1 == 17 {
		}
	}
}
