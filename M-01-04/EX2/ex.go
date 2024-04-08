package main

import "fmt"

func imprimirClima() {
	var (
		temperatura        int8    = 29
		umidade            float32 = 39.2
		pressaoAtmosferica int16   = 1019
	)

	fmt.Printf("A temperatura hoje está é %v °C, com %v para umidade. A pressão atmosférica é %vhPa. \n", temperatura, umidade, pressaoAtmosferica)
}

func main() {
	imprimirClima()
	fmt.Printf("\n")
}
