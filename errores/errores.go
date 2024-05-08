package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	texto, err := os.Open("ejemplo1_textolargo.txt")
	if err != nil {
		fmt.Println("Error de apertura de ejemplo1.txt", err)
		return
	}
	defer texto.Close()

	datos := make([]byte, 1024)
	totalBytes := 0

	for {
		n, err := texto.Read(datos)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Fin del archivo")
				break
			} else {
				fmt.Println("Error al leer el archivo", err)
				return
			}
		}
		totalBytes += n
		fmt.Println(string(datos[:n]))

	}

	fmt.Println("Bytes totales leidos: ", totalBytes)
}

/*
func main() {
	texto, err := os.Open("ejemplo1.txt")
	if err != nil {
		fmt.Println("Error de apertura de ejemplo1.txt", err)
		return
	}
	defer texto.Close()

	datos := make([]byte, 1024)
	n, err := texto.Read(datos)
	if err != nil {
		if err == io.EOF {
			fmt.Println("Fin del archivo")
		} else {
			fmt.Println("Error al leer el archivo", err)
			return
		}
	}

	fmt.Println("-------------------------------")
	fmt.Println(texto) // Esto probablemente no imprima nada útil
	fmt.Println("-------------------------------")
	fmt.Println(string(datos[:n]))
}
*/

/*
func main() {
	texto, err := os.Open("ejemplo1 .txt")
	if err != nil {
		fmt.Println("error al tratar de leer el archivo", err)
		defer texto.Close()
		return
	}
	datos := make([]byte, 1024)
	n, err := texto.Read(datos)
	if err != nil && err != io.EOF {
		fmt.Println("Error al leer el archivo", err)
		return
	}
	defer texto.Close()
	//fmt.Println(string(texto))
	fmt.Println("-------------------------------")
	fmt.Println(texto)
	fmt.Println("-------------------------------")
	fmt.Println(string(datos[:n]))
}
*/
