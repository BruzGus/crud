package main

import (
	"database/sql" 
	_ "github.com/lib/pq" 
	"log"
)

//GetConnection ... obtiene una conexion a la BD
func GetConnection() *sql.DB{
	dsn := "postgres://postgres:13r4An@127.0.0.1:5432/crud?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil{
		log.Fatal(err)
	}
	return db
}