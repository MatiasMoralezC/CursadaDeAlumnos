package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"os"
	"os/exec"
)

type Alumne struct {
	Id_alumne	int
	Nombre		string
	Apellido	string
	Dni			int
	Fecha_nacimiento	string
	Telefono	string
	Email		string
}

type Materia struct {
	Id_materia	int
	Nombre	string
}

type Comision struct {
	Id_materia 	int
	Id_comision int
	Cupo 		int
	}

type Correlatividad struct {
	Id_materia	int
	Id_mat_correlativa	int
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
	ejecutarPrograma()
}

func mostrarOpciones() int {
	fmt.Printf ("Elige una opcion para continuar:\n")
	fmt.Printf ("Para crear la DB, escriba el nùmero 1\n")
	fmt.Printf ("Para crear las tablas de l DB, escriba el nùmero 2\n")
	fmt.Printf ("Para cargar los datos de los archivos JSON, escriba el nùmero 3\n")
	fmt.Printf ("Para sarasa4, escriba el nùmero 4\n")
	fmt.Printf ("Para sarasa5, escriba el nùmero 5\n")
	fmt.Printf ("Para salir, escriba el nùmero 6\n")
	
	var opcion int
	fmt.Scanf("%d",&opcion)
	return opcion
}

func ejecutarPrograma() {
	fmt.Printf ("¡Bienvenido!\n")
	
	for {	
		clear()
		
		opcion := mostrarOpciones()
		
		clear()
		
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
			fmt.Printf("¡Hasta la proxima!\n")
			os.Exit(0)
		default:
			fmt.Printf("Opcion no valida\n")
		}
		
		if !preguntarContinuar() {
			break
		}
	}
}

func preguntarContinuar() bool {
	var continuar string
	fmt.Printf("¿Desea realizar otra accion? (s/n):\n")
	fmt.Scanf("%s",&continuar)
	return continuar == "s"
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	}

func createDatabase() {
	db,err := sql.Open("postgres", "user=postgres host=localhost dbname=postgres sslmode=disable")
	if err!= nil {
		log.Fatal(err)
	}
	defer db.Close()
	
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
	db,err := sql.Open("postgres", "user=postgres host=localhost dbname=prueba sslmode=disable")
	if err!= nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	_, err = db.Exec(`create table alumne(id_alumne int, nombre char(64), apellido char(64), dni int, fecha_nacimiento date, telefono char(64), email char(64))`)
	if err != nil {
		log.Fatal(err)
	}
		
	_, err = db.Exec(`create table materia(id_materia int, nombre char(64))`)
	if err != nil {
		log.Fatal(err)
	}
	
	_, err = db.Exec(`create table correlatividad(id_materia char(12), nombre char(64))`)
	if err != nil {
		log.Fatal(err)
	}
	
	_, err = db.Exec(`create table comision(id_materia int, id_comision int, cupo int)`)
	if err != nil {
		log.Fatal(err)
	}
	
	_, err = db.Exec(`create table cursada(id_materia int, id_alumne int, id_comision int, f_inscripcion timestamp, nota int, estado char(12))`)
	if err != nil {
		log.Fatal(err)
	}
	
	_, err = db.Exec(`create table periodo(semestre char(12), estado char(12))`)
	if err != nil {
		log.Fatal(err)
	}
	
	_, err = db.Exec(`create table historia_academica(ad_alumne int, semestre text, id_materia int, id_comision int, estado char(15), nota_regular int, nota_final int)`)
	if err != nil {
		log.Fatal(err)
	}
	
	_, err = db.Exec(`create table error(id_error int, operacion char(15), semestre text, id_alumne int, id_materia int, id_comision int, f_error timestamp, motivo char(64))`)
	if err != nil {
		log.Fatal(err)
	}
	
	_, err = db.Exec(`create table envio_mail(id_email int, f_generacion timestamp, email_alumne text, asunto text, cuerpo text, f_envio timestamp, estado char(10))`)
	if err != nil {
		log.Fatal(err)
	}
	
	_, err = db.Exec(`create table entrada_trx(id_orden int, operacion char(15), año int, nro_semestre int, id_alumne int, id_comision int)`)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Tablas cargadas.\n")
}	

func tablaExiste(tableName string) bool {
dbPrueba,err := sql.Open("postgres", "user=postgres host=localhost dbname=prueba sslmode=disable")
if err!= nil {
log.Fatal(err)
}
    var exists bool
    query := `SELECT EXISTS (
        SELECT FROM information_schema.tables
        WHERE table_schema = 'public' AND table_name = $1
    )`
    err := db.QueryRow(query, tableName).Scan(&exists)
    if err != nil {
        log.Fatal(err)
    }
    return exists
}

func ingresarPrimaryKey (){
dbPrueba,err := sql.Open("postgres", "user=postgres host=localhost dbname=prueba sslmode=disable")
if err!= nil {
log.Fatal(err)
}
defer dbPrueba.Close()

fmt.Printf("Ingrese el nombre de la tabla: ")
    var tableName string
    _, err = fmt.Scanln(&tableName)
    if err != nil {
        log.Fatal(err)
    }
    if !tablaExiste(tableName) {
        fmt.Printf("La tabla '%s' no existe en la base de datos.\n", tableName)
        return
    }
 
   
fmt.Printf("ingrese las primary keys:")
var primaryKeys string
_, err := fmt.Scanf(&primaryKeys)
if err != nil {
log.Fatal(err)
}
keys := strings.Split(primaryKeys, ",")

for _, key := range keys {
        key = strings.TrimSpace(key)
        var exists bool
        query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE id = $1)", tableName)
        if err != nil {
            log.Fatal(err)
        }
        if exists {
            fmt.Printf("La clave primaria '%s' existe en la tabla.\n", key)
        } else {
            fmt.Printf("La clave primaria '%s' NO existe en la tabla.\n", key)
        }
    }
}

func levantarJSons() {
	db,err := sql.Open("postgres", "user=postgres host=localhost dbname=prueba sslmode=disable")
	if err!= nil{
		log.Fatal(err)
	}
	defer db.Close()
	
	dataAlumnes, err := ioutil.ReadFile("alumnes.json")
	if err != nil{
		log.Fatal(err)
	}
	
	var alumnes []Alumne
	err = json.Unmarshal(dataAlumnes, &alumnes)
	if err != nil {
		log.Fatal(err)
	}
	
	for _, alumne := range alumnes {
		_, err := db.Exec("insert into alumne values ($1, $2, $3, $4, $5, $6, $7)", alumne.Id_alumne, alumne.Nombre, alumne.Apellido, alumne.Dni, alumne.Fecha_nacimiento, alumne.Telefono, alumne.Email)
		if err != nil{
			log.Fatal(err)
		}
	}
	
	fmt.Printf("Tabla de alumnes cargada.\n")
	
	dataMaterias, err := ioutil.ReadFile("materias.json")
	if err != nil{
		log.Fatal(err)
	}
	
	var materias []Materia
	err = json.Unmarshal(dataMaterias, &materias)
	if err != nil {
		log.Fatal(err)
	}
	
	for _, materia := range materias {
		_, err := db.Exec("insert into materia values ($1, $2)", materia.Id_materia, materia.Nombre)
		if err != nil{
			log.Fatal(err)
		}
	}
	
	fmt.Printf("Tabla de materias cargada.\n")
	
	dataComisiones, err := ioutil.ReadFile("comisiones.json")
	if err != nil{
		log.Fatal(err)
	}
	
	var comisiones []Comision
	err = json.Unmarshal(dataComisiones, &comisiones)
	if err != nil {
		log.Fatal(err)
	}
	
	for _, comision := range comisiones {
		_, err := db.Exec("insert into comision values ($1, $2, $3)", comision.Id_materia, comision.Id_comision, comision.Cupo)
		if err != nil{
			log.Fatal(err)
		}
	}
	
	fmt.Printf("Tabla de comisiones cargada.\n")
	
	dataPeriodos, err := ioutil.ReadFile("periodos.json")
	if err != nil{
		log.Fatal(err)
	}
	
	var periodos []Periodo
	err = json.Unmarshal(dataPeriodos, &periodos)
	if err != nil {
		log.Fatal(err)
	}
	
	for _, periodo := range periodos {
		_, err := db.Exec("insert into periodo values ($1, $2)", periodo.Semestre, periodo.Estado)
		if err != nil{
			log.Fatal(err)
		}
	}
	
	fmt.Printf("Tabla de periodos cargada.\n")
	
	dataCorrelativas, err := ioutil.ReadFile("correlatividades.json")
	if err != nil{
		log.Fatal(err)
	}
	
	var correlatividades []Correlatividad
	err = json.Unmarshal(dataCorrelativas, &correlatividades)
	if err != nil {
		log.Fatal(err)
	}
	
	for _, correlativa := range correlatividades {
		_, err := db.Exec("insert into correlatividad values ($1, $2)", correlativa.Id_materia, correlativa.Id_mat_correlativa)
		if err != nil{
			log.Fatal(err)
		}
	}
	
	fmt.Printf("Tabla de correlatividades cargada.\n")
	
	dataHistorias, err := ioutil.ReadFile("historia_academica.json")
	if err != nil{
		log.Fatal(err)
	}
	
	var historias_academicas []Historia_academica
	err = json.Unmarshal(dataHistorias, &historias_academicas)
	if err != nil {
		log.Fatal(err)
	}
	
	for _, historia_academica := range historias_academicas {
		_, err := db.Exec("insert into historia_academica values ($1, $2, $3, $4, $5, $6, $7)", historia_academica.Id_alumne, historia_academica.Semestre, historia_academica.Id_materia,
							historia_academica.Id_comision, historia_academica.Estado, historia_academica.Nota_regular, historia_academica.Nota_final)
		if err != nil{
			log.Fatal(err)
		}
	}
	
	fmt.Printf("Tabla de historias acadèmicas cargada.\n")
}








