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

    /*// Configurar la conexión a la base de datos
    connStr := "user=postgres host=localhost dbname=garcia_montoro_moralez_rodriguez_db1 sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Cargar las funciones SQL en la base de datos
    err = loadAllStoredProcedures(db)
    if err != nil {
        log.Fatalf("Error al cargar las funciones SQL: %v\n", err)
    }

    fmt.Println("Funciones SQL cargadas exitosamente.")

    // llamar a la función apertura_inscripcion
    query := `SELECT * FROM ingreso_nota($1, $2, $3, $4)`
    err = db.QueryRow(query, id_alumne_buscado, id_materia_buscada, id_comision_buscada, nota_ingresada).Scan(&p_result, &p_error_message)
    if err != nil {
        log.Fatal(err)
    }

    // mostrar el resultado
    fmt.Printf("Result: %v\n", p_result)
    if p_error_message != "" {
        fmt.Printf("Error: %s\n", p_error_message)
    } else {
        fmt.Println("Nota ingresada o actualizada exitosamente.")
    }*/
}

func mostrarOpciones() int {
	fmt.Printf ("Elige una opcion para continuar:\n")
	fmt.Printf ("Para crear la DB, escriba el nùmero 1\n")
	fmt.Printf ("Para crear las tablas de la DB, escriba el nùmero 2\n")
	fmt.Printf ("Para cargar los datos de los archivos JSON, escriba el nùmero 3\n")
	fmt.Printf ("Para agregar las primary keys, escriba el nùmero 4\n")
	fmt.Printf ("Para agregar las foreign keys, escriba el nùmero 5\n")
	fmt.Printf ("Para borrar las Primary Keys y las Foreign Keys 6\n")
	fmt.Printf ("Para cargar todos los Stored Procedures, escriba el nùmero 7\n")
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
		loadAllStoredProcedures()
		
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
					create table periodo(semestre char(6), estado char(15));
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

func loadAllStoredProcedures() {
	loadInscripcionMateria()
	loadAperturaInscripcion()
	loadAplicacionDeCupos()
	loadIngresoNota()
	loadBajaDeInscripcion()
	loadCierreDeInscripcion()
	
	loadEmailAltaInscripcion()
	loadEmailBajaInscripcion()
	loadEmailAplicacionCupos()
	loadEmailInscripcionEnEspera()
	loadEmailCierreCursada()
}

func loadInscripcionMateria() {
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
			
			select * into resultado_comision from comision where id_materia = id_materia_buscada and id_comision = id_comision_buscada;
			
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
					if materia_aprobada.id_materia = correlativa.id_materia_correlativa then
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

// CARGA EL SP EN LA DB NO LO EJECUTA, DSPS LO EJECUTAMOS EN EL MAIN
func loadAperturaInscripcion() {
    db,err := sql.Open("postgres", "user=postgres host=localhost dbname=garcia_montoro_moralez_rodriguez_db1 sslmode=disable")
	if err!= nil{
		log.Fatal(err)
	}
	defer db.Close()
	
	_, err = db.Exec(`
    create or replace function apertura_inscripcion(p_semestre varchar(6), out p_result boolean, out p_error_message text) as $$
    declare
        v_estado_actual varchar(6);
        v_anio_actual int;
        v_count int;
        v_semestre char(1);
    begin
        p_error_message := '';

        v_anio_actual := substring(p_semestre from 1 for 4)::int;
        v_semestre := substring(p_semestre from 6 for 1);

        if v_semestre not in ('1', '2') then
            p_result := false;
            p_error_message := 'número de semestre no válido';
            return;
        end if;

        if v_anio_actual < extract(year from current_date) then
            p_result := false;
            p_error_message := 'no se permiten inscripciones para un período anterior';
            return;
        end if;

        select estado into v_estado_actual from periodo where semestre = p_semestre;

        if v_estado_actual is not null and v_estado_actual != 'cerrado' then
            p_result := false;
            p_error_message := format('no es posible reabrir la inscripción del período, estado actual: %s', v_estado_actual);
            return;
        end if;

        select count(*) into v_count from periodo where estado in ('inscripcion', 'cierre inscrip') and semestre != p_semestre;

        if v_count > 0 then
            p_result := false;
            p_error_message := 'no es posible abrir otro período de inscripción, ya existe otro período en estado inscripción o cierre inscripción';
            return;
        end if;

        insert into periodo (semestre, estado) values (p_semestre, 'inscripcion')
        on conflict (semestre) do update set estado = excluded.estado;

        p_result := true;
    end;
    $$ language plpgsql;
    `)
	if err!= nil{
		log.Fatal(err)
	}
}


func loadBajaDeInscripcion (){
	db,err := sql.Open("postgres", "user=postgres host=localhost dbname=garcia_montoro_moralez_rodriguez_db1 sslmode=disable")
	if err!= nil{
	log.Fatal(err)
	}
	defer db.Close()
	//hardcodeo inscripcion en cursada
	_, err = db.Exec(`insert into periodo values('2026-1','cursada');`)
	if err!= nil{
	log.Fatal(err)
	}
	_, err = db.Exec(`
	create function bajaDeInscripcion(id_alumne_buscade integer, id_materia_buscada integer) returns void as $$
	declare
		resultado_periodo periodo%rowtype;
		resultado_alumne alumne%rowtype;
		resultado_materia materia%rowtype;
		resultado_comision comision%rowtype;
		resultado_cursada cursada%rowtype;


		cursada record;
		alumne_enespera record;

	begin
		select * into resultado_periodo from periodo where estado = 'inscripcion' or estado = 'cursada';

		if not found then
		raise 'no se permiten bajas en este periodo';
		end if;

		select * into resultado_alumne from alumne where id_alumne = id_alumne_buscade;

		if not found then
		raise 'id de alumne no válido';
		end if;

		select * into resultado_materia from materia where id_materia = id_materia_buscada;

		if not found then
		raise 'id de materia no válido';
		end if;

		select * into resultado_cursada from cursada where id_alumne = id_alumne_buscade and id_materia = id_materia_buscada and estado = 'ingresade';

		if found then
		raise 'alumne no inscripte en la materia';
		end if;

		update cursada set estado = 'dade de baja' where cursada.id_alumne = id_alumne_buscade;


		if resultado_periodo.estado = 'cursada' then
			select * from cursada where id_materia = id_materia_buscada and id_alumne = alumne_enespera.id and estado = 'en espera'
			order by f_inscripcion asc limit 1;
			if not found then
				raise 'alumne no cumple requisitos de correlatividad';
			end if;

			update cursada set estado = 'aceptade' where id_alumne = alumne_enespera.id_alumne ;
		end if;

	end;
	$$ language plpgsql;
`)

if err != nil {
log.Fatal(err)
}
}


func loadCierreDeInscripcion (){
	db,err := sql.Open("postgres", "user=postgres host=localhost dbname=garcia_montoro_moralez_rodriguez_db1 sslmode=disable")
	if err!= nil{
	log.Fatal(err)
	}
	defer db.Close()

	//hardcodeo inscripcion abierta
	_, err = db.Exec(`insert into periodo values('2025-2', 'inscripcion')`)
	if err!= nil{
	log.Fatal(err)
	}
	_, err = db.Exec(`
		create function cierreDeInscripcion(semestre_buscado text) returns void as $$
		declare
			resultado_periodo periodo%rowtype;

		begin

			select * into resultado_año from periodo where semestre = semestre_buscado and estado = 'inscripcion' ;

			if not found then
			raise 'el semestre no existe en periodo de inscripcion';
			end if;

			update periodo set estado = 'cierre inscrip' where semestre = semestre_buscado;

			end;
			$$ language plpgsql;
	`)
	}

func loadAplicacionDeCupos(){
	db,err := sql.Open("postgres", "user=postgres host=localhost dbname=garcia_montoro_moralez_rodriguez_db1 sslmode=disable")
	if err!= nil{
		log.Fatal(err)
	}
	defer db.Close()
	
	//Hardcodeo cierre de inscripcion
	_, err = db.Exec(`insert into periodo values('2025-1', 'cierre inscrip')`)
	if err!= nil{
		log.Fatal(err)
	}
	
	_, err = db.Exec(`
		create function aplicacion_cupos(semestre_buscado varchar(6)) returns void as $$
		declare
			periodo_encontrado periodo%rowtype;
			cupo_materia int;
			id_materia_buscada int := 1;
			id_comision_buscada int;
			alumne_inscripte cursada%rowtype;
			comision_materia comision%rowtype;
			materia comision%rowtype;
		begin
			select * into periodo_encontrado from periodo where semestre = semestre_buscado and estado = 'cierre inscrip';
			
			if not found then
				raise 'el semestre % no se encuentra en un período válido para aplicar cupos', semestre_buscado;
			end if;
			
			loop
				perform 1 from comision where id_materia = id_materia_buscada;
				exit when not found;
			
				for comision_materia in select * from comision where id_materia = id_materia_buscada loop
					id_comision_buscada := comision_materia.id_comision;
					
					select cupo into cupo_materia from comision where id_materia = id_materia_buscada and id_comision = comision_materia.id_comision;
					
					for alumne_inscripte in (select * from cursada where id_materia = id_materia_buscada and id_comision = id_comision_buscada and estado = 'ingresade' order by f_inscripcion asc limit cupo_materia) loop
						update cursada set estado = 'aceptade' 
						where id_alumne = alumne_inscripte.id_alumne 
						and id_materia = alumne_inscripte.id_materia 
						and id_comision = alumne_inscripte.id_comision;
					end loop;
				
				update cursada set estado = 'en espera'
				where id_materia = id_materia_buscada and id_comision = comision_materia.id_comision and estado = 'ingresade';
				
				end loop;
				
				id_materia_buscada := id_materia_buscada + 1;
			end loop;
			
			update periodo set estado = 'cursada'
			where estado = 'cierre inscrip';
			
		end;
		$$ language plpgsql;
	`)
	if err!= nil{
		log.Fatal(err)
	}
}

// CARGA LA NOTA DE CURSADA
func loadIngresoNota() {
	db,err := sql.Open("postgres", "user=postgres host=localhost dbname=garcia_montoro_moralez_rodriguez_db1 sslmode=disable")
	if err!= nil{
		log.Fatal(err)
	}
	defer db.Close()
	
	_, err = db.Exec(`
    create or replace function ingreso_nota(id_alumne_buscado int, id_materia_buscada int, id_comision_buscada int, nota_ingresada int, out p_result boolean, out p_error_message text) as $$
    declare
        v_count int;
    begin
        p_error_message := '';

        select count(*) into v_count from periodo where estado = 'cursada';

		if v_count = 0 then
            p_result := false;
            p_error_message := 'periodo de cursada cerrado';
            return;
        end if;
        
		if not exists (select 1 from alumne where id_alumne = id_alumne_buscado) then
	        p_result := false;
	        p_error_message := 'id de alumne no valido';
            return;
        end if;
        
		if not exists (select 1 from materia where id_materia = id_materia_buscada) then
	        p_result := false;
	        p_error_message := 'id de materia no valido';
            return;
        end if;
      
		if not exists (
			select 1 from comision
			where id_materia = id_materia_buscada and
			id_comision = id_comision_buscada
			) then
				p_result := false;
				p_error_message := 'id de comision no valido para la materia';
				return;
        end if;
        
		if not exists (
			select 1 from cursada
			where id_alumne = id_alumne_buscado and
			id_materia = id_materia_buscada and
			id_comision = id_comision_buscada
			) then
				p_result := false;
				p_error_message := 'alumne no cursa en la comision';
				return;
		end if;
        
		if nota < 0 or nota > 10 then
            p_result := false;
            p_error_message := 'nota no valida: ' || nota;
            return;
        end if; 
        
		update cursada 
		set nota = nota_ingresada
		where id_alumne = id_alumne_buscado 
		and id_materia = id_materia_buscada 
		and id_comision = id_comision_buscada;

		p_result := true;
		
	end;
	$$ language plpgsql;
	`)
	if err!= nil{
		log.Fatal(err)
	}
}

//Se dispara cuando una inscripcion es registrada
func loadEmailAltaInscripcion() {
	db,err := sql.Open("postgres", "user=postgres host=localhost dbname=garcia_montoro_moralez_rodriguez_db1 sslmode=disable")
	if err!= nil{
		log.Fatal(err)
	}
	defer db.Close()
	
	query := `
	create or replace function email_alta_inscipcion() returns trigger as $$
	declare
		v_materia_nombre text;
		v_comision_numero text;
		v_alumne_nombre text;
		v_alumne_apellido text;
		v_email_alumne text;
	begin
		select nombre into v_materia_nombre from materia where id_materia = NEW.id_materia;
		select id_comision into v_comision_numero from comision where id_materia = NEW.id_materia and id_comision = NEW.id_comision;
		select nombre, apellido, email into v_alumne_nombre, v_alumne_apellido, v_email_alumne from alumne where id_alumne = NEW.id_alumne;
		
		insert into envio_email (f_generacion, email_alumne, asunto, cuerpo, f_envio, estado)
		values (current_timestamp, v_email_alumne, 'Inscripcion registrada', 
		'Hola ' || alumne_nombre || ' ' || alumne_apellido || ', tu inscripcion a la materia ' || materia_nombre || ', comision ' || comision_numero ||' ha sido registrada.',
		null, 'pendiente' 
		);
				
		return old;
	end;
	$$ language plpgsql;
	
	create trigger email_alta_inscripcion_trg
	after insert on cursada
	for each row
	when (NEW.estado = 'aceptade')
	execute function email_alta_inscipcion();
	`
	
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Trigger 'email alta inscipcion' creado con exito.\n")
}

//Se dispara cuando la inscripcion es dada de baja
func loadEmailBajaInscripcion() {
	db,err := sql.Open("postgres", "user=postgres host=localhost dbname=garcia_montoro_moralez_rodriguez_db1 sslmode=disable")
	if err!= nil{
		log.Fatal(err)
	}
	defer db.Close()
	
	query := `
	create or replace function email_baja_inscripcion() returns trigger as $$
	declare
		v_materia_nombre text;
		v_comision_numero text;
		v_alumne_nombre text;
		v_alumne_apellido text;
		v_email_alumne text;
	begin
		select nombre into v_materia_nombre from materia where id_materia = OLD.id_materia;
		select id_comision into v_comision_numero from comision where id_materia = OLD.id_materia and id_comision = OLD.id_comision;
		select nombre, apellido, email into v_alumne_nombre, v_alumne_apellido, v_email_alumne from alumne where id_alumne = OLD.id_alumne;
		
		insert into envio_email (f_generacion, email_alumne, asunto, cuerpo, f_envio, estado)
		values (current_timestamp, v_email_alumne, 'Inscripcion dada de baja',
		'Hola ' || alumne_nombre || ' ' || alumne_apellido || ', tu inscripcion a la materia ' || materia_nombre || ', comision ' || comision_numero ||' ha sido dada de baja.',
		null, 'pendiente'
		);
		
		return old;
	end;
	$$ language plpgsql;
		
	create trigger email_baja_inscripcion_trg
	after update on cursada
	for each row
	when (NEW.estado = 'dade de baja')
	execute function email_baja_inscripcion();
	`
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Trigger 'email baja inscipcion' creado con exito.\n")
	
}

//Se dispara cuando la inscripcion es aceptada
func loadEmailAplicacionCupos() {
	db,err := sql.Open("postgres", "user=postgres host=localhost dbname=garcia_montoro_moralez_rodriguez_db1 sslmode=disable")
	if err!= nil{
		log.Fatal(err)
	}
	defer db.Close()
	
	query := `
	create or replace function email_aplicacion_cupos() returns trigger as $$
	declare
		v_materia_nombre text;
		v_comision_numero text;
		v_alumne_nombre text;
		v_alumne_apellido text;
		v_email_alumne text;
		v_estado_inscripcion char(12);
	begin
		select nombre into v_materia_nombre from materia where id_materia = NEW.id_materia;
		select id_comision into v_comision_numero from comision where id_materia = NEW.id_materia and id_comision = NEW.id_comision;
		select nombre, apellido, email into v_alumne_nombre, v_alumne_apellido, v_email_alumne from alumne where id_alumne = NEW.id_alumne;
		
		if new.estado = 'aceptade' then
			v_estado_inscripcion := 'aceptade';
		else
			v_estado_inscripcion := 'en espera';
		end if;
		
		insert into envio_email (f_generacion, email_alumne, asunto, cuerpo, f_envio, estado)
		values (current_timestamp, v_email_alumne, v_estado_inscripcion,
		'Hola ' || alumne_nombre || ' ' || alumne_apellido || ', tu inscripcion a la materia ' || materia_nombre || ', comision ' || comision_numero ||' se encuentra: ' || 'v_estado_inscripcion.',
		null, 'pendiente'
		);
		
		return new;
	end;
	$$ language plpgsql;

	create trigger email_aplicacion_cupos_trg
	after update on cursada
	for each row
	when (OLD.estado = 'en espera' and NEW.estado = 'aceptade')
	execute	function email_aplicacion_cupos();
	`
	
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Trigger 'email aplicacion de cupo' creado con exito.\n")
}


// Se dispara cuando la inscripcion en espera pasa a aceptada
func loadEmailInscripcionEnEspera() {
	db,err := sql.Open("postgres", "user=postgres host=localhost dbname=garcia_montoro_moralez_rodriguez_db1 sslmode=disable")
	if err!= nil{
		log.Fatal(err)
	}
	defer db.Close()
	
	query := `
	create or replace function email_inscripcion_lista_espera() returns trigger as $$
	declare
		v_materia_nombre text;
		v_comision_numero text;
		v_alumne_nombre text;
		v_alumne_apellido text;
		v_email_alumne text;
	begin
		select nombre into v_materia_nombre from materia where id_materia = NEW.id_materia;
		select id_comision into v_comision_numero from comision where id_materia = NEW.id_materia and id_comision = NEW.id_comision;
		select nombre, apellido, email into v_alumne_nombre, v_alumne_apellido, v_email_alumne from alumne where id_alumne = NEW.id_alumne;
		
		insert into envio_email (f_generacion, email_alumne, asunto, cuerpo, f_envio, estado)
		values (current_timestamp, v_email_alumne, 'Inscripcion aceptada',
		'Hola ' || alumne_nombre || ' ' || alumne_apellido || ', tu inscripcion a la materia ' || materia_nombre || ', comision ' || comision_numero ||' ha sido finalmente aceptada',
		null, 'pendiente'
		);
	
		return new;
	end;
	$$ language plpgsql;
	
	create trigger email_inscripcion_lista_espera_trg
	after update on cursada
	for each row
	when (OLD.estado = 'en espera' and NEW.estado = 'aceptade')
	execute function email_inscripcion_lista_espera();
	`
	
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Trigger 'email inscripcion en espera' creado con exito.\n")
}

//Se ejecuta cuando se cierra la cursada de una comision_numero
func loadEmailCierreCursada() {
	db,err := sql.Open("postgres", "user=postgres host=localhost dbname=garcia_montoro_moralez_rodriguez_db1 sslmode=disable")
	if err!= nil{
		log.Fatal(err)
	}
	defer db.Close()
	
	query := 	`
	create or replace function email_cierre_cursada() returns trigger as $$
	declare
		v_materia_nombre text;
		v_comision_numero text;
		v_alumne_nombre text;
		v_alumne_apellido text;
		v_email_alumne text;
		v_estado_academico char(15);
		v_nota_regular int;
		v_nota_final int;
		v_semestre_actual text;
	begin
		select nombre into v_materia_nombre from materia where id_materia = NEW.id_materia;
		select id_comision into v_comision_numero from comision where id_materia = NEW.id_materia and id_comision = NEW.id_comision;
		select nombre, apellido, email into v_alumne_nombre, v_alumne_apellido, v_email_alumne from alumne where id_alumne = NEW.id_alumne;
		select semestre into v_semestre_actual from periodo where estado = 'cursada';
		select estado, nota_regular, nota_final into v_estado_academico, v_nota_regular, v_nota_final from historia_academica
		where id_alumne = NEW.id_alumne and id_materia = NEW.id_materia and semestre = v_semestre_actual;
		
		insert into envio_email (f_generacion, email_alumne, asunto, cuerpo, f_envio, estado)
		values (current_timestamp, 'Cierre de cursada',
		'Hola ' || alumne_nombre || ' ' || alumne_apellido || ', tu inscripcion a la materia ' || materia_nombre || ', comision ' || comision_numero ||
		' ha sido cerrada. Estado: ' || v_estado_academico || ', Nota regular: ' || COALESCE(v_nota_regular::text, 'N/A') || ', Nota final: ' || COALESCE(v_nota_final::text, 'N/A') || '.',
		null, 'pendiente'
		);
		
		return new;
	end;
	
	$$ language plpgsql;
	
	create trigger email_cierre_cursada_trg
	after update on cursada
	for each row
	when (NEW.estado = 'cerrado')
	execute function email_cierre_cursada();
	`
	
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Trigger 'email cierre cursada' creado con exito.\n")
}









