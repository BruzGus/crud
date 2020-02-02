package main

import (
	"database/sql" 
	_ "github.com/lib/pq" 
	"log"
)

//getConnection obtiene una conexion a la BD

func getConnection() *sql.DB{
	dsn := "postgres://postgres:13r4An@127.0.0.1:5432/crud?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil{
		log.fatal(err)
	}
	return db
}