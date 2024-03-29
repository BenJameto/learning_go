import json
import os

class Nota:
    def __init__(self, titulo, contenido):
        self.titulo = titulo
        self.contenido = contenido

    def to_dict(self):
        return {"titulo": self.titulo, "contenido": self.contenido}

    @staticmethod
    def from_dict(nota_dict):
        return Nota(nota_dict["titulo"], nota_dict["contenido"])

class Categoria:
    def __init__(self, nombre):
        self.nombre = nombre
        self.notas = []

    def agregar_nota(self, nota):
        self.notas.append(nota)

    def mostrar_notas(self):
        print(f"Notas en la categoría {self.nombre}:")
        for nota in self.notas:
            print(f"- {nota.titulo}")

    def borrar_nota(self, titulo):
        self.notas = [nota for nota in self.notas if nota.titulo != titulo]

    def to_dict(self):
        return {"nombre": self.nombre, "notas": [nota.to_dict() for nota in self.notas]}

    @staticmethod
    def from_dict(categoria_dict):
        categoria = Categoria(categoria_dict["nombre"])
        categoria.notas = [Nota.from_dict(nota_dict) for nota_dict in categoria_dict["notas"]]
        return categoria

class NotasApp:
    def __init__(self):
        self.categorias = {}
        self.cargar_notas_desde_archivo()

    def bienvenida(self):
        print("¡Bienvenido a la aplicación de notas!")

    def nueva_nota(self):
        titulo = input("Ingrese el título de la nota: ")
        contenido = input("Ingrese el contenido de la nota: ")

        categoria = input("En qué categoría desea guardar la nota? ")
        if categoria not in self.categorias:
            crear_categoria = input("La categoría no existe, ¿desea crearla? (s/n) ")
            if crear_categoria.lower() == "s":
                self.categorias[categoria] = Categoria(categoria)
            else:
                print("Nota no guardada.")
                return

        self.categorias[categoria].agregar_nota(Nota(titulo, contenido))
        self.guardar_notas_en_archivo()
        print(f"Nota guardada exitosamente en la categoría {categoria}")

    def borrar_nota(self):
        print("Categorías disponibles:")
        for nombre_categoria in self.categorias:
            print(f"- {nombre_categoria}")

        nombre_categoria = input("Ingrese el nombre de la categoría de la nota que desea borrar: ")
        if nombre_categoria not in self.categorias:
            print("Categoría no encontrada.")
            return

        categoria = self.categorias[nombre_categoria]
        categoria.mostrar_notas()
        titulo_nota = input("Ingrese el título de la nota que desea borrar: ")
        categoria.borrar_nota(titulo_nota)
        self.guardar_notas_en_archivo()
        print("Nota borrada exitosamente.")

    def cargar_notas_desde_archivo(self):
        if os.path.exists("notas.json"):
            with open("notas.json", "r") as file:
                data = json.load(file)
                for nombre_categoria, categoria_data in data.items():
                    self.categorias[nombre_categoria] = Categoria.from_dict(categoria_data)

    def guardar_notas_en_archivo(self):
        with open("notas.json", "w") as file:
            data = {nombre_categoria: categoria.to_dict() for nombre_categoria, categoria in self.categorias.items()}
            json.dump(data, file, indent=4)

    def iniciar(self):
        self.bienvenida()

        while True:
            print("\n¿Qué desea hacer?")
            print("1. Crear una nueva nota")
            print("2. Borrar una nota")
            print("3. Salir del programa")
            opcion = input("Ingrese el número de la opción: ")

            if opcion == "1":
                self.nueva_nota()
            elif opcion == "2":
                self.borrar_nota()
            elif opcion == "3":
                print("¡Hasta luego!")
                break
            else:
                print("Opción no válida. Intente de nuevo.")

if __name__ == "__main__":
    app = NotasApp()
    app.iniciar()
