package main

import "fmt"

func main101001() {
	var total int
	fmt.Println("Digite o valor para o saque: ")
	fmt.Scan(&total)

	qtdDuz := 0
	qtdCem := 0
	qtdCin := 0
	qtdVin := 0
	qtdDez := 0
	qtdCinc := 0
	qtdDois := 0

	if total >= 200 {
		qtdDuz = total / 200
		total = total % 200
	}
	if total >= 100 {
		qtdCem = total / 100
		total = total % 100
	}
	if total >= 50 {
		qtdCin = total / 50
		total = total % 50
	}
	if total >= 20 {
		qtdVin = total / 20
		total = total % 20
	}
	if total >= 10 {
		qtdDez = total / 10
		total = total % 10
	}
	if total >= 5 {
		qtdCinc = total / 5
		total = total % 5
	}
	if total >= 2 {
		qtdDois = total / 2
		total = total % 2
	}

	fmt.Printf("Notas de duzentos: %d\n", qtdDuz)
	fmt.Printf("Notas de cem: %d\n", qtdCem)
	fmt.Printf("Notas de cinquenta: %d\n", qtdCin)
	fmt.Printf("Notas de vinte: %d\n", qtdVin)
	fmt.Printf("Notas de dez: %d\n", qtdDez)
	fmt.Printf("Notas de cinco: %d\n", qtdCinc)
	fmt.Printf("Notas de dois: %d\n", qtdDois)
}
