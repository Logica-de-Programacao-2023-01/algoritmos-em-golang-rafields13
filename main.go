package main

import "fmt"

func main() {

	//Algoritmos de SELEÇÃO

	// 1. Faça um algoritmo que leia dois números inteiros e mostre o maior deles.

	var x float64
	fmt.Print("Escolha um número: ")
	fmt.Scan(&x)
	fmt.Println(x, "é um ótimo número!")

	var y float64
	fmt.Print("Escolha outra número: ")
	fmt.Scan(&y)
	fmt.Println(y, "é um outro número perfeito para operação!")

	if x > y {
		fmt.Println("x é maior do que y.")
	} else if x < y {
		fmt.Println("y é maior do que x.")
	} else {
		fmt.Println("x é igual a y.")
	}

	//2. Faça um algoritmo que leia três números inteiros e mostre o menor deles.

	var a float64
	fmt.Print("Escolha um número: ")
	fmt.Scan(&a)
	fmt.Println(a, "é um excelente número para começarmos a operação")

	var b float64
	fmt.Print("Escolha outro número: ")
	fmt.Scan(&b)
	fmt.Println(b, "é outra excelente escolha!")

	var g float64
	fmt.Print("Selecione o último número a ser escolhido:")
	fmt.Scan(&g)
	fmt.Println(g, "é perfeito, vamos começar a operação!")

	if a < b && a < g {
		fmt.Println("a é o menor número.")
	} else if b < a && b < g {
		fmt.Println("b é o menor número.")
	} else {
		fmt.Println("g é o menor número.")
	}

	// 3. Faça um algoritmo que leia um número inteiro e mostre se ele é par ou ímpar.

	h := 2

	var z int
	fmt.Print("Escolha um número: ")
	fmt.Scan(&z)
	fmt.Println(z, "é um ótimo número para operarmos!")

	if z%h == 0 {
		fmt.Println("z é numero par.")
	} else {
		fmt.Println("z é número ímpar.")
	}
}
