package main

import "fmt"

func imprimirMes(mes int) string {
	/* Escolhi realizar esta solução pois temos uma determinada ordem de "equivalência" entre os meses,
	o que é possível assim, fazer com que o valor dado corresponda com o index equivalente.*/
	meses := [12]string{"Janeiro", "Fevereiro", "Março", "Abril", "Maio", "Junho",
		"Julho", "Agosto", "Setembro", "Outubro", "Novembro", "Dezembro"}

	if mes < 1 || mes > 12 {
		return "Valor inválido!"
	} else {
		return meses[mes-1]
	}
}

func main() {
	fmt.Println(imprimirMes(3))
}
