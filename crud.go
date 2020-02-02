package main

import (
	"fmt"
	"log"
)

func main() {
	//instanciamos un estudiante
	e := Estudiante{
			Name: "Bryan Augusto Cruz Castro",
			Age: 29,
			Active : false,
	}

	err := EstudianteCrear(e)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Creado exitosamente")

}