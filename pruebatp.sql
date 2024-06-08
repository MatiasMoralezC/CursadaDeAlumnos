DROP DATABASE IF EXISTS garcia_montoro_moralez_rodriguez_db1;
CREATE DATABASE garcia_montoro_moralez_rodriguez_db1;

\C garcia_montoro_moralez_rodriguez_db1

CREATE TABLE alumne(
		id_alumne INTEGER PRIMARY KEY,
		nombre VARCHAR(64),
		apellido VARCHAR(64),
		dni VARCHAR(64),
		fecha_nacimiento DATE,
		telefono VARCHAR(64),
		email VARCHAR(64)
);

CREATE TABLE materia(
		id_materia: INTEGER PRIMARY KEY,
		nombre VARCHAR(64)
);

CREATE TABLE correlatividad(
		id_materia INTEGER,
		id_mat_correlativa INTEGER,
		PRIMARY KEY (id_materia, id_mat_correlativa),
		FOREIGN KEY (id_materia) REFERENCES materia (id_materia),
		FOREIGN KEY (id_mat_correlativa) REFERENCES materia _id_materia)
);

CREATE TABLE comision(
		id_materia INTEGER,
		id_comision INTEGER,
		cupo INTEGER,
		PRIMARY KEY (id_materia, id_comision),
		FOREIGN KEY (id_materia) REFERENCES materia (id_materia)
);
