package main

//import "fmt"

//Lista 02

/*1) Ler 3 valores (A, B e C) representando as medidas dos lados de um triângulo e escrever se
formam ou não um triângulo. OBS: para formar um triângulo, o valor de cada lado deve ser
menor que a soma dos outros 2 lados.

func main() {
	var a, b, c int

	fmt.Printf("Digite os valores dos lados do triângulo: \n")
	fmt.Printf("Lado A: ")
	fmt.Scan(&a)
	fmt.Printf("Lado B: ")
	fmt.Scan(&b)
	fmt.Printf("Lado C: ")
	fmt.Scan(&c)

	if a < (b+c) && b < (a+c) && c < (b+a) {
		fmt.Printf("Essas medias formam um triângulo")
	} else {
		fmt.Printf("Essas medias não formam um triângulo")
	}}*/

/*2) Faça um algoritmo para ler um número que é um código de usuário. Caso este código seja
diferente de um código armazenado internamente no algoritmo (igual a 1234) deve ser
apresentada a mensagem ‘Usuário inválido!’. Caso o Código seja correto, deve ser lido outro
valor que é a senha. Se esta senha estiver incorreta (a certa é 9999) deve ser mostrada a
mensagem ‘senha incorreta’. Caso a senha esteja correta, deve ser mostrada a mensagem
‘Acesso permitido’.

func main(){
	var user = 1234; 
	var senha = 9999;
	var teclado int;

	fmt.Printf("Digite o número do usuário: ")
	fmt.Scan(&teclado)

	if teclado == user {
		fmt.Printf("Digite a senha: ")
		fmt.Scan(&teclado)

		if senha == teclado {
			fmt.Printf("Acesso permitido")
		}else{
			fmt.Printf("Senha incorreta")
		}
	}else{
		fmt.Printf("Usuário inválido!")
	}}*/

/*3) Ler 3 valores (considere que não serão informados valores iguais) e escrever o maior deles.

func main(){
	var a, b, c, maior int

	fmt.Printf("Digite três números \n")
	fmt.Printf("A: ")
	fmt.Scan(&a)
	fmt.Printf("B: ")
	fmt.Scan(&b)
	fmt.Printf("C: ")
	fmt.Scan(&c)

	maior = a

	if b > maior {
		maior = b
	}
	if c > maior {
		maior = c
	}
	
    fmt.Printf("O maior número é: %d", maior)
}*/

/*4) Construir um algoritmo que leia dois números e efetue a adição. Caso o valor somado seja
maior que 20, este deverá ser apresentado somando-se a ele mais 8; caso o valor somado seja
menor ou igual a 20, este deverá ser apresentado subtraindo-se 5.

func main(){
	var num1, num2 int
	
	fmt.Printf("Digite dois números para serem somados: \n")
	fmt.Printf("Número 1: ")
	fmt.Scan(&num1)
    fmt.Printf("Número 2: ")
	fmt.Scan(&num2)

	soma := num1+num2

	if soma > 20 {
		fmt.Printf("O resultado da soma é: %d", soma+8)
	}else{
		fmt.Printf("O resultado da soma é: %d", soma-5)
	}
}*/

/*5) Faça um algoritmo que leia um número e mostre uma das mensagens: é múltiplo de 3 ou
não é múltiplo de 3.

func main171024() {
	var num int
	fmt.Print("Digite um número para saber se ele é ou não multiplo de 3: ")
	fmt.Scan(&num)

	if num % 3 == 0 {
		fmt.Println("Este número é multiplo de 3!")
	}else{
		fmt.Println("Este número não é multiplo de 3!")}}*/

/*6) Faça um algoritmo que leia um número e mostre se ele é ou não divisível por 3 e por 7.

func main171024(){
	var num int
	fmt.Print("Digite um número para saber se ele é ou não divisível por 3 e por 7: ")
	fmt.Scan(&num)

	if num % 3 == 0 && num % 7 == 0 {
		fmt.Println("Este número é divisível por 3 e por 7!")
	}else if num % 3 == 0{
		fmt.Println("Este número é divisível por 3 e não por 7!")
	}else if num % 7 == 0{
		fmt.Println("Este número é divisível por 7 e não por 3!")
	}else{
		fmt.Println("Este número não é divisível por 3 ou por 7!")}}*/