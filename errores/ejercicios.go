package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
)

func main() {
	fmt.Println("Ejercicio 1")
	r, err := raiz(-5.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}
	fmt.Println("-------------------------------------")
	fmt.Println("Ejercicio 2")
	lineas, err := cuentaLineas("ejemplo 1.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("el archivo ejemplo1.txt tiene ", lineas, " lineas")
	}
	fmt.Println("-------------------------------------")
	fmt.Println("Ejercicio 3")

	/*
	   numero1 := 2
	   numero2 := "2"

	   return numero1 == numero2
	*/
}

/*
Esta funcion recibe un numero n para obtener su raiz cuadrada. Si n es negativo
esta funcion regresara el error "No se puede sacar raiz a un numero negativo"
*/
func raiz(n float64) (float64, error) {
	if n < 0 {
		return 0, errors.New("No se puede sacar raiz a un numero negativo")
	}
	r := math.Sqrt(n)
	return r, nil

}
func cuentaLineas(nombreArchivo string) (int, error) {
	archivo, err := os.Open(nombreArchivo)
	if err != nil {
		return 0, errors.New("error al leer el archivo")
	}
	defer archivo.Close()

	lineas := 0
	scanner := bufio.NewScanner(archivo)

	for scanner.Scan() {
		lineas++
	}
	if err := scanner.Err(); err != nil {
		return 0, errors.New("error al leer le archivo")
	}
	return lineas, nil
}
