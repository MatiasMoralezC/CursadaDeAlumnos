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
	opcion := pedirOpcion()
	elegirOpcion(opcion)
}

func pedirOpcion() int {
	var opcion int
	fmt.Printf ("¡Bienvenido! Elige una opciòn para continuar:\n")
	fmt.Printf ("Para crear la DB, escriba el nùmero 1\n")
	fmt.Printf ("Para crear las tablas de l DB, escriba el nùmero 2\n")
	fmt.Printf ("Para cargar los datos de los archivos JSON, escriba el nùmero 3\n")
	fmt.Printf ("Para sarasa4, escriba el nùmero 4\n")
	fmt.Printf ("Para sarasa5, escriba el nùmero 5\n")
	fmt.Printf ("Para salir, escriba el nùmero 6\n")
	fmt.Scanf("%d",&opcion)
	return opcion
}

func elegirOpcion(opcion int) {
	switch opcion {
		case 1:
			createDatabase()
			
		case 2:
			createDbTables()
			
		case 3:
			levantarJSons()
			
		case 4:
			fmt.Printf("falta agregar funcion\n")
			//funcion1()
			
		case 5:
			fmt.Printf("falta agregar funcion\n")
			//funcion1()
			
		case 6:
			fmt.Printf("¡Hasta la pròxima!\n")
	}		
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
	
	fmt.Printf("Base de datos creada.\n")
}

func createDbTables() {
	dbPrueba,err := sql.Open("postgres", "user=postgres host=localhost dbname=prueba sslmode=disable")
	if err!= nil {
		log.Fatal(err)
	}
	defer dbPrueba.Close()
	
	_, err = dbPrueba.Exec(`create table alumne(id_alumne int, nombre char(64), apellido char(64), dni int, fecha_nacimiento date, telefono char(64), email char(64))`)
	if err != nil {
		log.Fatal(err)
	}
		
	_, err = dbPrueba.Exec(`create table materia(id_materia int, nombre char(64))`)
	if err != nil {
		log.Fatal(err)
	}
	
	_, err = dbPrueba.Exec(`create table correlatividad(id_materia int, nombre char(64))`)
	if err != nil {
		log.Fatal(err)
	}
	
	_, err = dbPrueba.Exec(`create table comision(id_materia int, id_comision int, cupo int)`)
	if err != nil {
		log.Fatal(err)
	}
	
	_, err = dbPrueba.Exec(`create table cursada(id_materia int, id_alumne int, id_comision int, f_inscripcion timestamp, nota int, estado char(12))`)
	if err != nil {
		log.Fatal(err)
	}
	
	_, err = dbPrueba.Exec(`create table periodo(semestre int, estado char(12))`)
	if err != nil {
		log.Fatal(err)
	}
	
	_, err = dbPrueba.Exec(`create table historia_academica(ad_alumne int, semestre text, id_materia int, id_comision int, estado char(15), nota_regular int, nota_final int)`)
	if err != nil {
		log.Fatal(err)
	}
	
	_, err = dbPrueba.Exec(`create table error(id_error int, operacion char(15), semestre text, id_alumne int, id_materia int, id_comision int, f_error timestamp, motivo char(64))`)
	if err != nil {
		log.Fatal(err)
	}
	
	_, err = dbPrueba.Exec(`create table envio_mail(id_email int, f_generacion timestamp, email_alumne text, asunto text, cuerpo text, f_envio timestamp, estado char(10))`)
	if err != nil {
		log.Fatal(err)
	}
	
	_, err = dbPrueba.Exec(`create table entrada_trx(id_orden int, operacion char(15), año int, nro_semestre int, id_alumne int, id_comision int)`)
	if err != nil {
		log.Fatal(err)
	}
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







