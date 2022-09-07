package main

import "fmt"

func main() {
	var arrays [5]int = [5]int{1, 2, 3, 4, 5} //arrays
	var porcion []int = []int{1, 2, 3, 4, 5}  // porcion

	porcion = append(porcion, 5)
	//arrays = append(arrays, 5) Error de arrays

	fmt.Println("[arrays] Capacidad Arrays: ", cap(arrays), "Longitud: ", len(arrays))

	datoActual := 0
	for i := 0; i < 1000; i++ {
		porcion = append(porcion, i)

		datoProcesar := cap(porcion)
		if !(datoActual == datoProcesar) {
			fmt.Println("[porcion] Capacidad: ", cap(porcion), "Longitud: ", len(porcion))
			datoActual = datoProcesar
		}

	}

	/*
		Los Slice multiplica su capacidad por cada vez que la supere, cuando llega un momento sera mucha la asignacion de memoria y se utiliza otro formato

	*/
	capacidad := 128
	porcion2 := make([]int, capacidad)
	fmt.Println("[Slice]Capacidad: ", cap(porcion2), "Longitud: ", len(porcion2))
}

