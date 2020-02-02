package main

import (
	"errors"
	"time"
)

//Estudiante ... estructura del estudiante
type Estudiante struct {

	ID int
	Name string
	Age int16
	Active bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

// EstudianteCrear ..., registra un estudiante en la Base de Datos
func EstudianteCrear(e Estudiante) error{
	query := `INSERT INTO 
				estudiantes ( name, age, active)
				VALUES ($1, $2, $3)`

	db := GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)  // preparando la sentencia
	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err:= stmt.Exec(e.Name, e.Age, e.Active)
	if err != nil {
		return err
	}

	i, _ := r.RowsAffected() // controlamos las inserciones
	if i != 1 {
		return errors.New("Erro:Se esperaba una fila afectada")
	}
	return nil

}