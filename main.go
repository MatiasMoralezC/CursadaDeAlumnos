package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type Alumne struct {
	Id_alumne	int	`json:id_alumne`
	Nombre		string	`json:nombre`
	Apellido	string	`json:apellido`
	Dni			int	`json:id_alumne`
	Fecha_nacimiento	string	`json:fecha_nacimiento`
	Telefono	string	`json:telefono`
	Email		string	`json:email`
}

type Materia struct {
	Id_materia	int
	Nombre	string
}

type Correlatividad struct {
	Id_materia	int
	Id_mat_correlatva	int
}

type Cursada struct {
	Id_materia	int
	Id_alumne	int
	Id_comision	int
	F_inscripcion	string
	Nota		int
	Estado		string
}

type Periodo struct {
	Semestre	string
	Estado	string
}

type Historia_academica struct {
	Id_alumne	int
	Semestre	string
	Id_materia	int
	Id_comision	int
	Estado	string
	Nota_regular	int
	Nota_final	int
}

type Error struct {
	Id_error	int
	Operacion	string
	Semestre	string
	Id_alumne	int
	Id_materia	int
	Id_comision	int
	F_error		string
	Motivo		string
}

type Envio_mail struct {
	Id_email	int
	F_generacion	string
	Email_alumne	string
	Asunto		string
	Cuerpo		string
	F_envio		string
	Estado		string
}

func main() {
	createDatabase()
	levantarJSons()
}

func createDatabase() {
	db,err := sql.Open("postgres", "user=postgres host=localhost dbname=postgres sslmode=disable")
	if err!= nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	_, err = db.Exec(`SELECT pg_terminate_backend(pid) FROM pg_stat_activity WHERE datname = 'mi_base_de_datos';`)
	if err != nil {
		log.Fatal(err)
	}
	
	_, err = db.Exec(`drop database if exists prueba;`)
	if err != nil {
		log.Fatal(err)
	}
	
	_, err = db.Exec(`create database prueba;`)
	if err != nil {
		log.Fatal(err)
	}
	
	dbPrueba,err := sql.Open("postgres", "user=postgres host=localhost dbname=prueba sslmode=disable")
	if err!= nil {
		log.Fatal(err)
	}
	defer dbPrueba.Close()
	
	_, err = dbPrueba.Exec(`create table alumne(id_alumne int, nombre char(64), apellido char(64), dni int, fecha_nacimiento date, telefono char(64), email char(64))`)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Base de datos creada.\n")
}

func levantarJSons() {
	dataAlumnes, err := ioutil.ReadFile("alumnes.json")
	if err != nil{
		log.Fatal(err)
	}
	
	var alumnes []Alumne
	err = json.Unmarshal(dataAlumnes, &alumnes)
	if err != nil {
		log.Fatal(err)
	}
	
	db,err := sql.Open("postgres", "user=postgres host=localhost dbname=prueba sslmode=disable")
	if err!= nil{
		log.Fatal(err)
	}
	defer db.Close()
	
	for _, alumne := range alumnes {
		_, err := db.Exec("insert into alumne values ($1, $2, $3, $4, $5, $6, $7)", alumne.Id_alumne, alumne.Nombre, alumne.Apellido, alumne.Dni, alumne.Fecha_nacimiento, alumne.Telefono, alumne.Email)
		if err != nil{
			log.Fatal(err)
		}
	}
	
	fmt.Printf("Tabla de alumnes cargada.\n")
}
