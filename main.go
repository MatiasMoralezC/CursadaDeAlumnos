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
	Id_alumne int
	Nombre string
	Apellido string
	Dni int
	Fecha_nacimiento string
	Telefono string
	Email string
}

type Materia struct {
	Id_materia int
	Nombre string
}

type Comision struct {
	Id_materia int
	Id_comision int
	Cupo int
}

type Correlatividad struct {
	Id_materia int
	Id_mat_correlativa int
}

type Cursada struct {
	Id_materia int
	Id_alumne int
	Id_comision int
	F_inscripcion string
	Nota int
	Estado string
}

type Periodo struct {
	Semestre string
	Estado string
}

type Historia_academica struct {
	Id_alumne int
	Semestre string
	Id_materia int
	Id_comision int
	Estado string
	Nota_regular int
	Nota_final int
}

type Error struct {
	Id_error int
	Operacion string
	Semestre string
	Id_alumne int
	Id_materia int
	Id_comision int
	F_error string
	Motivo string
}

type Envio_mail struct {
	Id_email int
	F_generacion string
	Email_alumne string
	Asunto string
	Cuerpo string
	F_envio string
	Estado string
}

func main() {
	ejecutarPrograma()
}

func mostrarOpciones() int {
	fmt.Printf ("Elige una opcion para continuar:\n")
	fmt.Printf ("Para crear la DB, escriba el nùmero 1\n")
	fmt.Printf ("Para crear las tablas de la DB, escriba el nùmero 2\n")
	fmt.Printf ("Para cargar los datos de los archivos JSON, escriba el nùmero 3\n")
	fmt.Printf ("Para agregar las primary keys, escriba el nùmero 4\n")
	fmt.Printf ("Para agregar las foreign keys, escriba el nùmero 5\n")
	fmt.Printf ("Para borrar las Primary Keys y las Foreign Keys 6\n")
	fmt.Printf ("Para realizar la inscripciòn a una materia, escriba el nùmero 7\n")
	fmt.Printf ("Para salir, escriba el nùmero 8\n")

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
		agregarPrimaryKey ()

		case 5:
		agregarForeignKey ()

		case 6:
		borrarKeys ()

		case 7:
		inscripcionMateria()

		case 8:
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

	_, err = db.Exec(`drop database if exists garcia_montoro_moralez_rodriguez_db1;`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`create database garcia_montoro_moralez_rodriguez_db1;`)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Base de datos creada.\n")
}

func createDbTables() {
	db,err := sql.Open("postgres", "user=postgres host=localhost dbname=garcia_montoro_moralez_rodriguez_db1 sslmode=disable")
	if err!= nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`create table alumne(id_alumne int, nombre char(64), apellido char(64), dni int, fecha_nacimiento date, telefono char(64), email char(64));
					create table materia(id_materia int, nombre char(64));
					create table correlatividad(id_materia int, id_materia_correlativa int);
					create table comision(id_materia int, id_comision int, cupo int);
					create table cursada(id_materia int, id_alumne int, id_comision int, f_inscripcion timestamp, nota int, estado char(12));
					create table periodo(semestre char(12), estado char(12));
					create table historia_academica(id_alumne int, semestre text, id_materia int, id_comision int, estado char(15), nota_regular int, nota_final int);
					create table error(id_error int, operacion char(15), semestre text, id_alumne int, id_materia int, id_comision int, f_error timestamp, motivo char(64));
					create table envio_mail(id_email int, f_generacion timestamp, email_alumne text, asunto text, cuerpo text, f_envio timestamp, estado char(10));
					create table entrada_trx(id_orden int, operacion char(15), año int, nro_semestre int, id_alumne int, id_comision int);`)
	if err!= nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Tablas cargadas.\n")
}

func agregarPrimaryKey (){
	db,err := sql.Open("postgres", "user=postgres host=localhost dbname=garcia_montoro_moralez_rodriguez_db1 sslmode=disable")
	if err!= nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`alter table alumne add constraint pk_alumne primary key (id_alumne);
					alter table materia add constraint pk_materia primary key (id_materia);
					alter table correlatividad add constraint pk_correlatividad primary key (id_materia, id_materia_correlativa);
					alter table comision add constraint pk_comision primary key (id_materia, id_comision);
					alter table cursada add constraint pk_cursada primary key (id_materia, id_alumne);
					alter table periodo add constraint pk_periodo primary key (semestre);
					alter table historia_academica add constraint pk_academica primary key (id_alumne, semestre, id_materia);
					alter table error add constraint pk_error primary key (id_error);
					alter table envio_mail add constraint pk_envio_mail primary key (id_email);`)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Primary Keys cargadas.\n")
}

func agregarForeignKey (){
	db,err := sql.Open("postgres", "user=postgres host=localhost dbname=garcia_montoro_moralez_rodriguez_db1 sslmode=disable")
	if err!= nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`alter table correlatividad add constraint fk_materia foreign key (id_materia) references materia (id_materia);
					alter table correlatividad add constraint fk_correlativa foreign key (id_materia_correlativa) references materia (id_materia);
					alter table comision add constraint fk_materia foreign key (id_materia) references materia (id_materia);
					alter table cursada add constraint fk_materia foreign key (id_materia) references materia (id_materia);
					alter table cursada add constraint fk_alumne foreign key (id_alumne) references alumne (id_alumne);
					alter table historia_academica add constraint fk_alumne foreign key (id_alumne) references alumne (id_alumne);
					alter table historia_academica add constraint fk_periodo foreign key (semestre) references periodo (semestre);
					alter table historia_academica add constraint fk_materia foreign key (id_materia) references materia (id_materia);
					alter table error add constraint fk_alumne foreign key (id_alumne) references alumne (id_alumne);
					alter table error add constraint fk_periodo foreign key (semestre) references periodo (semestre);
					alter table error add constraint fk_materia foreign key (id_materia) references materia (id_materia);`)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Foreign Keys cargadas.\n")
}

func borrarKeys (){
	db,err := sql.Open("postgres", "user=postgres host=localhost dbname=garcia_montoro_moralez_rodriguez_db1 sslmode=disable")
	if err!= nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`alter table cursada drop constraint fk_materia;
					alter table cursada drop constraint fk_alumne;
					alter table historia_academica drop constraint fk_materia;
					alter table historia_academica drop constraint fk_alumne;
					alter table historia_academica drop constraint fk_periodo;
					alter table error drop constraint fk_materia;
					alter table error drop constraint fk_alumne;
					alter table error drop constraint fk_periodo;
					alter table correlatividad drop constraint fk_materia;
					alter table correlatividad drop constraint fk_correlativa;
					alter table comision drop constraint fk_materia;
					alter table alumne drop constraint pk_alumne;
					alter table materia drop constraint pk_materia;
					alter table correlatividad drop constraint pk_correlatividad;
					alter table comision drop constraint pk_comision;
					alter table cursada drop constraint pk_cursada;
					alter table periodo drop constraint pk_periodo;
					alter table historia_academica drop constraint pk_academica;
					alter table error drop constraint pk_error;
					alter table envio_mail drop constraint pk_envio_mail;`)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Primary Keys y Foreign Keys borradas.\n")
}

func levantarJSons() {
	db,err := sql.Open("postgres", "user=postgres host=localhost dbname=garcia_montoro_moralez_rodriguez_db1 sslmode=disable")
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
	
	fmt.Printf("Tabla de historias académicas cargada.\n")
}

func inscripcionMateria() {
	db,err := sql.Open("postgres", "user=postgres host=localhost dbname=garcia_montoro_moralez_rodriguez_db1 sslmode=disable")
	if err!= nil{
		log.Fatal(err)
	}
	defer db.Close()
	
	//Hardcodeado: es para que haya un periodo abierto de inscripcion para que se pueda realizar bien.
	
	_, err = db.Exec(`insert into periodo values('2024-2','inscripcion');`)
	if err!= nil{
		log.Fatal(err)
	}
	
	_, err = db.Exec(`
		create function inscripcion_materia(id_alumne_buscado integer, id_materia_buscada integer, id_comision_buscada integer) returns void as $$
		declare
			resultado_periodo periodo%rowtype;
			resultado_alumne alumne%rowtype;
			resultado_materia materia%rowtype;
			resultado_comision comision%rowtype;
			resultado_cursada cursada%rowtype;
			materia_aprobada historia_academica%rowtype;
			correlativa correlatividad%rowtype;
			materia_encontrada boolean;
			correlativas_aprobadas boolean;
		begin
			select * into resultado_periodo from periodo where estado = 'inscripcion';
			
			if not found then
				raise 'periodo de inscripción cerrado';
			end if;
			
			select * into resultado_alumne from alumne where id_alumne = id_alumne_buscado;
			
			if not found then
				raise 'id de alumne no válido';
			end if;
			
			select * into resultado_materia from materia where id_materia = id_materia_buscada;
			
			if not found then
				raise 'id de materia no válido';
			end if;
			
			select * into resultado_comision from comision where id_comision = id_comision_buscada;
			
			if not found then
				raise 'id de comisión no válido';
			end if;
			
			select * into resultado_cursada from cursada where id_alumne = id_alumne_buscado and id_materia = id_materia_buscada and estado = 'aceptade';
			
			if found then
				raise 'alumne ya inscripte en la materia';
			end if;
			
			correlativas_aprobadas := true;
			for correlativa in select * from correlatividad where id_materia = id_materia_buscada loop
				materia_encontrada := false;
				for materia_aprobada in select * from historia_academica where id_alumne = id_alumne_buscado and (estado = 'regular' or estado = 'aprobada') loop
					if extract(id_materia from materia_aprobada) = extract(id_materia_correlativa from correlativa) then
						materia_encontrada = true;
					end if;
				end loop;
				
				if not materia_encontrada then
					correlativas_aprobadas = false;
				end if;
			end loop;
			
			if not correlativas_aprobadas then
				raise 'alumne no cumple requisitos de correlatividad';
			end if;
			
			insert into cursada values(id_materia_buscada, id_alumne_buscado, id_comision_buscada, current_timestamp, null, 'ingresade');
			
			end;
			$$ language plpgsql;
	`)
	if err != nil {
		log.Fatal(err)
	}
}








