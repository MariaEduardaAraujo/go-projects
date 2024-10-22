package main

import ("fmt")//,"math")

//Lista 05

/*5) Apresente o quadrado de cada um dos números pares entre 1 e 1000, inclusive.

func aula191024(){
	for i := 0; i <= 1000; i++ {
		if i % 2 == 0 {
			quadrado := math.Pow(float64(i),2)
			fmt.Printf("O quadrado de %d é: %.1f\n", i, quadrado)}}}*/

/*6) Apresente todos os números divisíveis por 5 que sejam maiores do que 0 e menores ou
iguais a 200.

func aula191024()  {
    fmt.Println("Os números divisíveis por 5, entre 1 e 200 são: ")
	for i := 1; i <= 200; i++ {
		if i % 5 == 0 {
			fmt.Println(i)}}}*/

/*7) Escreva um algoritmo que leia 10 valores quaisquer. A seguir, mostre quantos deles estão
dentro do intervalo (10,20) e quantos estão fora do intervalo, mostrando essas informações.*/

func main(){
	var num[10]int
	fmt.Printf("Digite 10 números: ")
	for i := 0; i < 10; i++ {
		fmt.Scan(&num[i])
	}
}