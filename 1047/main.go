package main

import "fmt"

func main()  {
	var horaInicial, minutoInicial, horaFinal, minutoFinal int
	fmt.Scan(&horaInicial, &minutoInicial, &horaFinal, &minutoFinal)
	valorEmMinutosInicial := horaInicial * 60 + minutoInicial
	valorEmMinutosFinal := horaFinal * 60 + minutoFinal
	if valorEmMinutosFinal == valorEmMinutosInicial {
		fmt.Println("O JOGO DUROU 24 HORA(S) E 0 MINUTO(S)")
	} else if valorEmMinutosFinal > valorEmMinutosInicial {
		quantidadeHoras := (valorEmMinutosFinal - valorEmMinutosInicial) / 60
		quantidadeMinutos := (valorEmMinutosFinal - valorEmMinutosInicial) % 60
		fmt.Printf("O JOGO DUROU %d HORA(S) E %d MINUTO(S)\n", quantidadeHoras, quantidadeMinutos)
	} else {
		quantidadeHoras := (valorEmMinutosFinal - valorEmMinutosInicial + 1440 ) / 60
		quantidadeMinutos := (valorEmMinutosFinal - valorEmMinutosInicial + 1440 ) % 60
		fmt.Printf("O JOGO DUROU %d HORA(S) E %d MINUTO(S)\n", quantidadeHoras, quantidadeMinutos)
	}	
}