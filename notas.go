package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Nota struct {
	Titulo    string `json:"titulo"`
	Contenido string `json:"contenido"`
}

func NewNota(titulo, contenido string) *Nota {
	return &Nota{Titulo: titulo, Contenido: contenido}
}

type Categoria struct {
	Nombre string  `json:"nombre"`
	Notas  []*Nota `json:"notas"`
}

func NewCategoria(nombre string) *Categoria {
	return &Categoria{Nombre: nombre}
}

func (c *Categoria) AgregarNota(nota *Nota) {
	c.Notas = append(c.Notas, nota)
}

func (c *Categoria) MostrarNotas() {
	fmt.Printf("Notas en la categoría %s:\n", c.Nombre)
	for _, nota := range c.Notas {
		fmt.Printf("- %s\n", nota.Titulo)
	}
}

func (c *Categoria) BorrarNota(titulo string) {
	nuevasNotas := make([]*Nota, 0, len(c.Notas))
	for _, nota := range c.Notas {
		if nota.Titulo != titulo {
			nuevasNotas = append(nuevasNotas, nota)
		}
	}
	c.Notas = nuevasNotas
}

type NotasApp struct {
	Categorias map[string]*Categoria
}

func NewNotasApp() *NotasApp {
	app := &NotasApp{
		Categorias: make(map[string]*Categoria),
	}
	app.CargarNotasDesdeArchivo()
	return app
}

func (app *NotasApp) Bienvenida() {
	fmt.Println("¡Bienvenido a la aplicación de notas!")
}

func (app *NotasApp) NuevaNota() {
	titulo := leerEntrada("Ingrese el título de la nota: ")
	contenido := leerEntradaMultilinea("Ingrese el contenido de la nota (presione Enter para terminar):\n")
	categoria := leerEntrada("En qué categoría desea guardar la nota? ")

	if _, existe := app.Categorias[categoria]; !existe {
		crearCategoria := leerEntrada("La categoría no existe, ¿desea crearla? (s/n) ")
		if crearCategoria == "s" {
			app.Categorias[categoria] = NewCategoria(categoria)
		} else {
			fmt.Println("Nota no guardada.")
			return
		}
	}

	app.Categorias[categoria].AgregarNota(NewNota(titulo, contenido))
	app.GuardarNotasEnArchivo()
	fmt.Printf("Nota guardada exitosamente en la categoría %s\n", categoria)
}

func (app *NotasApp) BorrarNota() {
	fmt.Println("Categorías disponibles:")
	for nombre := range app.Categorias {
		fmt.Printf("- %s\n", nombre)
	}

	nombreCategoria := leerEntrada("Ingrese el nombre de la categoría de la nota que desea borrar: ")
	categoria, existe := app.Categorias[nombreCategoria]
	if !existe {
		fmt.Println("Categoría no encontrada.")
		return
	}

	categoria.MostrarNotas()
	tituloNota := leerEntrada("Ingrese el título de la nota que desea borrar: ")
	categoria.BorrarNota(tituloNota)
	app.GuardarNotasEnArchivo()
	fmt.Println("Nota borrada exitosamente.")
}

func (app *NotasApp) CargarNotasDesdeArchivo() {
	file, err := os.Open("notas.json")
	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Println("Error al abrir el archivo de notas:", err)
		}
		return
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&app.Categorias)
	if err != nil {
		fmt.Println("Error al decodificar el archivo de notas:", err)
	}
}

func (app *NotasApp) GuardarNotasEnArchivo() {
	file, err := os.Create("notas.json")
	if err != nil {
		fmt.Println("Error al crear el archivo de notas:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(app.Categorias)
	if err != nil {
		fmt.Println("Error al codificar el archivo de notas:", err)
	}
}

func leerEntrada(mensaje string) string {
	var entrada string
	fmt.Print(mensaje)
	fmt.Scanln(&entrada)
	return entrada
}

func leerEntradaMultilinea(mensaje string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(mensaje)
	var lineas []string
	for scanner.Scan() {
		linea := scanner.Text()
		if linea == "" {
			break
		}
		lineas = append(lineas, linea)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error al leer la entrada:", err)
	}
	return strings.Join(lineas, "\n")
}

func main() {
	app := NewNotasApp()
	app.Bienvenida()

	for {
		fmt.Println("\n¿Qué desea hacer?")
		fmt.Println("1. Crear una nueva nota")
		fmt.Println("2. Borrar una nota")
		fmt.Println("3. Salir del programa")
		opcion := leerEntrada("Ingrese el número de la opción: ")

		switch opcion {
		case "1":
			app.NuevaNota()
		case "2":
			app.BorrarNota()
		case "3":
			fmt.Println("¡Hasta luego!")
			return
		default:
			fmt.Println("Opción no válida. Intente de nuevo.")
		}
	}
}
