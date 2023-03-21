package main

import "fmt"

func main() {

	//Algoritmos de SEQUÊNCIA

	// 1. Faça um algoritmo que leia três números inteiros e mostre a soma entre eles.

	var m float64
	fmt.Print("Escolha um número: ")
	fmt.Scan(&m)
	fmt.Println(m, "é um número interessante.")

	var n float64
	fmt.Print("Escolha outro número: ")
	fmt.Scan(&n)
	fmt.Println(n, "é outra ótima escolha.")

	var o float64
	fmt.Print("Escolha um último número: ")
	fmt.Scan(&o)
	fmt.Println(o, "é uma excelente escolha! Vamos começar a operação!")

	var soma = m + n + o
	fmt.Println(soma, "é a soma dos números escolhidos por você! :)")

	// 2. Faça um algoritmo que leia um número inteiro e mostre o seu dobro, triplo e quadruplo.

	var p float64
	fmt.Print("Escolha um número: ")
	fmt.Scan(&p)
	fmt.Println(p, "é uma ótima escolha!")

	var dobro = p * 2
	fmt.Println(dobro, "é", p, "multiplicado por 2.")

	var triplo = p * 3
	fmt.Println(triplo, "é", p, "multiplicado por 3.")

	var quádruplo = p * 4
	fmt.Println(quádruplo, "é", p, "multiplicado por 4.")

	// 3. Faça um algoritmo que leia o peso e a altura de uma pessoa e calcule o seu IMC (Índice de Massa Corporal).

	var altura float64
	fmt.Print("Indique sua altura aqui (em metros) -> ")
	fmt.Scan(&altura)
	fmt.Println("Sua altura é", altura, "metros.")

	var peso float64
	fmt.Print("Indique seu peso aqui (em kgs.) -> ")
	fmt.Scan(&peso)
	fmt.Println("Você pesa", peso, "kgs..")

	var imc float64 = peso / (altura * altura)
	fmt.Println(imc, "é seu índice de massa corporal.")

	// 4. Faça um algoritmo que leia três números reais e mostre a média ponderada entre eles, com pesos 2, 3 e 5, respectivamente.

	var q float64
	fmt.Print("Escolha um número: ")
	fmt.Scan(&q)
	fmt.Println(q, "é uma excelente escolha!")

	var r float64
	fmt.Print("Escolha mais um número: ")
	fmt.Scan(&r)
	fmt.Println(r, "é outra boa opção!")

	var s float64
	fmt.Print("Escolha um último número: ")
	fmt.Scan(&s)
	fmt.Println(s, "não poderia ser uma escolha ruim!")

	var média_ponderada float64 = ((2 * q) + (3 * r) + (5 * s)) / 10
	fmt.Println(média_ponderada, "é a média ponderada dos números escolidos por você com pesos 2, 3 e 5, respectivamente.")

	// 5. Faça um algoritmo que leia a idade de uma pessoa em anos e mostre a idade em dias.

	var idade int
	fmt.Print("Digite aqui a sua idade: ")
	fmt.Scan(&idade)
	fmt.Println("Você tem", idade, "anos.")

	var idade_em_dias int = idade * 365
	fmt.Println("Você já viveu", idade_em_dias, "dias.")

	// 6. Faça um algoritmo que leia o salário de um funcionário e calcule o seu novo salário com um aumento de 15%.

	var salário_atual float64
	fmt.Print("Digite aqui o salário atual: ")
	fmt.Scan(&salário_atual)
	fmt.Println("O salário atual é de", salário_atual, "reais.")

	var salário_com_aumento float64 = (salário_atual * 15 / 100) + salário_atual
	fmt.Println("O salário com um aumento de 15% seria de", salário_com_aumento, "reais.")

	// 7. Faça um algoritmo que leia um número inteiro e mostre o seu sucessor e antecessor.

	var t float64
	fmt.Print("Escolha um número: ")
	fmt.Scan(&t)
	fmt.Println(t, "não é uma má escolha!")

	var antecessor float64 = t - 1
	fmt.Println(antecessor, "é o antecessor do número escolhido por você.")

	var sucessor float64 = t + 1
	fmt.Println(sucessor, "é o sucessor do número escolhido por você.")

	// 8. Faça um algoritmo que leia o número de dias trabalhados por um funcionário e o valor da sua diária e calcule o seu salário.

	var diária float64
	fmt.Print("Digite aqui o valor da sua diária: ")
	fmt.Scan(&diária)
	fmt.Println(diária, "reais é o valor da sua diária.")

	var dias_trabalhados float64
	fmt.Print("Coloque aqui quantos dias você trabalhou nesse mês: ")
	fmt.Scan(&dias_trabalhados)
	fmt.Println("Você trabalhou", dias_trabalhados, "dias nesse mês.")

	var calc_sal float64 = dias_trabalhados * diária
	fmt.Println(calc_sal, "reais é o salário desse mês (com base na quantidade de dias trabalhados e no valor da diária).")

	// 9. Faça um algoritmo que leia o preço de um produto e mostre o seu valor com desconto de 10%.

	var preço_produto float64
	fmt.Print("Digite aqui o preço do produto sem desconto: ")
	fmt.Scan(&preço_produto)
	fmt.Println("O preço do produto sem desconto é de", preço_produto, "reais.")

	var produto_com_desconto float64 = preço_produto - (preço_produto * 10 / 100)
	fmt.Println("O preço do produto com um desconto de 10% seria de", produto_com_desconto, "reais.")

	// 10. Faça um algoritmo que leia o peso de uma pessoa em quilos e converta para libras.

	var quilos float64
	fmt.Print("Indique seu peso aqui (em kgs.) -> ")
	fmt.Scan(&quilos)
	fmt.Println("Você pesa", quilos, "kgs..")

	var libras float64 = quilos * 2.205
	fmt.Println("Você pesa", libras, "libras.")

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

	// 2. Faça um algoritmo que leia três números inteiros e mostre o menor deles.

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

	// 4. Faça um algoritmo que leia a altura e o sexo de uma pessoa e mostre se ela está abaixo, dentro ou acima do peso ideal, utilizando as fórmulas do exercício 3 da lista de algoritmos sequenciais.

	var altura1 float64
	fmt.Print("Indique sua altura aqui (em metros) -> ")
	fmt.Scan(&altura1)
	fmt.Println("Sua altura é", altura1, "metros.")

	var peso1 float64
	fmt.Print("Indique seu peso aqui (em kgs.) -> ")
	fmt.Scan(&peso1)
	fmt.Println("Você pesa", peso1, "kgs..")

	var imc1 float64 = peso1 / (altura1 * altura1)
	fmt.Println(imc1, "é seu índice de massa corporal.")

	var sexo string
	fmt.Print("Indique seu sexo aqui (masculino ou feminino) -> ")
	fmt.Scan(&sexo)
	fmt.Println(sexo, "é o seu sexo.")

	if imc < 18.5 {
		fmt.Println("Você está abaixo do peso normal.")
	} else if imc < 24.9 {
		fmt.Println("Você está dentro do peso normal.")
	} else {
		fmt.Print("Você está acima do peso normal.")
	}

	// 5. Faça um algoritmo que leia um número inteiro e mostre se ele é múltiplo de 3 e de 5 ao mesmo tempo.

	var u int
	fmt.Print("Escolha um número: ")
	fmt.Scan(&u)
	fmt.Println(u, "não é nada mau!")

	if u%3 == 0 && u%5 == 0 {
		fmt.Println(u, "é múltiplo de 3 e de 5 ao mesmo tempo.")
	} else {
		fmt.Println(u, "não é múltiplo de 3 e de 5 ao mesmo tempo.")
	}

	// 6. Faça um algoritmo que leia dois números inteiros e mostre o resultado da multiplicação entre eles, se ambos forem positivos; ou a soma, se pelo menos um deles for negativo.

	var v float64
	fmt.Print("Escolha um número: ")
	fmt.Scan(&v)
	fmt.Println(v, "é um excelente número!")

	var w float64
	fmt.Print("Escolha outro número: ")
	fmt.Scan(&w)
	fmt.Println(w, "também não é nada mau!")

	if v > 0 && w > 0 {
		fmt.Println(v*w, "é a multiplicação entre as variáveis.")
	} else {
		fmt.Println(v+w, "é a soma entre as varáveis.")
	}

	// 7. Faça um algoritmo que leia o salário de um funcionário e calcule o seu novo salário com um aumento de 10% se o salário for menor que R$ 1000,00; ou de 5% se o salário for maior ou igual a R$ 1000,00.

	var sal float64
	fmt.Print("Digite aqui o valor do seu salário: ")
	fmt.Scan(&sal)
	fmt.Println("Seu salário é de", sal, "reais.")

	if sal < 1000 {
		fmt.Println("O salário com um aumento de 10% é de", (sal*10/100)+sal, "reais.")
	} else if sal >= 1000 {
		fmt.Println("O salário com um aumento de 5% é de", (sal*5/100)+sal, "reais.")
	}

	// 8. Faça um algoritmo que leia um número inteiro entre 1 e 7 e mostre o dia da semana correspondente (1 = domingo, 2 = segunda-feira, etc.).

	var número_entre_1_e_7 int
	fmt.Print("Escolha um número entre 1 e 7: ")
	fmt.Scan(&número_entre_1_e_7)

	if número_entre_1_e_7 == 1 {
		fmt.Println("Hoje é Domingo!")
	} else if número_entre_1_e_7 == 2 {
		fmt.Println("Hoje é Segunda-feira!")
	} else if número_entre_1_e_7 == 3 {
		fmt.Println("Hoje é Terça-feira!")
	} else if número_entre_1_e_7 == 4 {
		fmt.Println("Hoje é Quarta-feira!")
	} else if número_entre_1_e_7 == 5 {
		fmt.Println("Hoje é Quinta-feira!")
	} else if número_entre_1_e_7 == 6 {
		fmt.Println("Hoje é Sexta-feira!")
	} else if número_entre_1_e_7 == 7 {
		fmt.Println("Hoje é Sábado!")
	} else {
		fmt.Print("Selecione apenas números inteiros entre 1 e 7.")
	}

	// 9. Faça um algoritmo que leia três números reais e mostre-os em ordem crescente.

	var n1 float64
	fmt.Print("Escolha um número: ")
	fmt.Scan(&n1)
	fmt.Println(n1, "é um excelente número!")

	var n2 float64
	fmt.Print("Escolha um número: ")
	fmt.Scan(&n2)
	fmt.Println(n2, "não é um número ruim!")

	var n3 float64
	fmt.Print("Escolha um número: ")
	fmt.Scan(&n3)
	fmt.Println(n3, "é um outro ótimo número para operarmos!")

	if n1 < n2 && n1 < n3 && n2 < n3 {
		fmt.Println(n1, n2, n3)
	} else if n1 < n2 && n1 < n3 && n2 > n3 {
		fmt.Println(n1, n3, n2)
	} else if n2 < n1 && n2 < n3 && n1 < n3 {
		fmt.Println(n2, n1, n3)
	} else if n2 < n1 && n2 > n3 && n1 > n3 {
		fmt.Println(n2, n3, n1)
	} else if n3 < n1 && n3 < n2 && n1 < n2 {
		fmt.Println(n3, n1, n2)
	} else if n3 < n1 && n3 < n2 && n1 > n2 {
		fmt.Println(n3, n2, n1)
	}

	// 10. Faça um algoritmo que leia a idade de uma pessoa e mostre a sua classificação, de acordo com a seguinte tabela:
	//até 9 anos: mirim
	//10 a 13 anos: infantil
	//14 a 17 anos: juvenil
	//maiores de 18 anos: adulto

	var idade1 int
	fmt.Print("Coloque aqui sua idade: ")
	fmt.Scan(&idade1)
	fmt.Println("Você tem", idade1, "anos.")

	if idade1 <= 9 {
		fmt.Println("Você está na categoria mirim.")
	} else if idade1 >= 10 && idade1 <= 13 {
		fmt.Println("Você está na categoria infantil.")
	} else if idade1 >= 14 && idade1 <= 17 {
		fmt.Println("Você está na categoria juvenil.")
	} else {
		fmt.Println("Você está na categoria adulto.")
	}
}
