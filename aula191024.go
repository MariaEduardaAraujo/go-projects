package main

import (
	"fmt"
	"math"
)

//Lista 05

/*1) Leia 5 valores para uma variável A. A seguir mostre quantos valores digitados foram pares,
quantos valores digitados foram ímpares, quantos foram positivos e quantos foram negativos.

func main(){
	var A[] int
	var num int
	var par, impar, maior, menor int

	fmt.Printf("Digite 5 números: \n")
	for i := 0; i < 5; i++ {
		fmt.Printf("Número %d: ", i+1)
		fmt.Scan(&num)
		A = append(A,num)
	}

	for j := 0; j < len(A); j++ {
		if A[j] % 2 == 0 {
			par += 1
		}else{
			impar += 1
		}

		if A[j] > 0 {
			maior += 1
		}else{
			menor += 1
		}
	}

    fmt.Printf("\n")
	fmt.Printf("Par: %d\n", par)
	fmt.Printf("Ímpar: %d\n", impar)
	fmt.Printf("Maior que 0: %d\n", maior)
	fmt.Printf("Menor que 0: %d\n", menor)
}*/

/*2) Calcule e mostre a soma dos números pares entre 1 e 100, inclusive.

func aula191024(){
	var soma int

	for i := 1; i <= 100; i++ {
		if i % 2 == 0 {
			soma = soma + i
		}
	}

	fmt.Printf("A soma dos números pares de 1 até 100 é: %d", soma)
}*/

/*3) Calcule e mostre a média dos números pares entre 1 e 100, inclusive.

func aula191024(){
	var soma, media, c int

	for i := 1; i <= 100; i++ {
		if i % 2 == 0 {
			soma = soma + i
			c+=1
		}
	}

	media = soma/c

	fmt.Printf("A média dos números pares é: %d", media)
}*/

/*4) Leia 2 valores: X e Y. A seguir, calcule e mostre a soma dos números impares entre eles*/

func aula191024(){
	var num1, num2 int

	fmt.Printf("Digite dois números: \n")
	fmt.Printf("Número 1: ")
	fmt.Scan(&num1)
	fmt.Printf("Número 2: ")
	fmt.Scan(&num2)

	for i := 0; i < count; i++ {
		
	}
}

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
dentro do intervalo (10,20) e quantos estão fora do intervalo, mostrando essas informações.

func aula191024(){
    var c int
	var num[] int
	var dentro[] int
	var fora[] int

	fmt.Printf("Digite 10 números: \n")
	for i := 0; i < 10; i++ {
		fmt.Printf("Número %d: ", i+1)
		fmt.Scan(&c)
		num = append(num, c)
	}

	for j := 0; j < len(num); j++ {
	    if num[j] >= 10 && num[j] <= 20{
		    dentro = append(dentro, num[j])
		}else{
		    fora = append(fora, num[j])
		}
	}
	fmt.Printf("Números dentro do intervalo: %d\n", dentro)
	fmt.Printf("Números fora do intervalo: %d\n", fora)
}*/