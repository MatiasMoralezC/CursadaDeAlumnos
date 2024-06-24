create function baja_de_inscripcion(id_alumne_buscado integer, id_materia_buscada integer, out p_result boolean, out p_error_message text) as $$
declare
	resultado_periodo periodo%rowtype;
	resultado_alumne alumne%rowtype;
	resultado_materia materia%rowtype;
	resultado_comision comision%rowtype;
	resultado_cursada cursada%rowtype;

	alumne_enespera record;

begin
	select * into resultado_periodo from periodo where estado = 'inscripcion' or estado = 'cursada';

	if not found then
		insert into error values(nextval('error_id_seq'), 'baja inscrip', null, id_alumne_buscado, id_materia_buscada, null, current_timestamp, 'No se permiten bajas en este periodo');
		p_error_message := 'no se permiten bajas en este periodo';
		p_result := false;
		return;
	end if;

	select * into resultado_alumne from alumne where id_alumne = id_alumne_buscado;

	if not found then
		insert into error values(nextval('error_id_seq'), 'baja inscrip', resultado_periodo.semestre, id_alumne_buscado, id_materia_buscada, null, current_timestamp, 'Id de alumne no válido');
		p_error_message := 'id de alumne no válido';
		p_result := false;
		return;
	end if;

	select * into resultado_materia from materia where id_materia = id_materia_buscada;

	if not found then
		insert into error values(nextval('error_id_seq'), 'baja inscrip', resultado_periodo.semestre, id_alumne_buscado, id_materia_buscada, null, current_timestamp, 'Id de materia no válido');
		p_error_message := 'id de materia no válido';
		p_result := false;
		return;
	end if;

	select * into resultado_cursada from cursada where id_alumne = id_alumne_buscado and id_materia = id_materia_buscada and estado = 'aceptade';

	if not found then
		insert into error values(nextval('error_id_seq'), 'baja inscrip', resultado_periodo.semestre, id_alumne_buscado, id_materia_buscada, null, current_timestamp, 'Alumne no inscripte en la materia');
		p_error_message := 'alumne no inscripte en la materia';
		p_result := false;
		return;
	end if;

	update cursada set estado = 'dade de baja' where cursada.id_alumne = id_alumne_buscado and cursada.id_materia = id_materia_buscada;
	
	if resultado_periodo.estado = 'cursada' then
		select * into alumne_enespera from cursada 
		where id_materia = id_materia_buscada and id_comision = resultado_cursada.id_comision and estado = 'en espera' 
		order by f_inscripcion asc limit 1;
		
		update cursada set estado = 'aceptade' 
		where id_alumne = alumne_enespera.id_alumne and id_materia = id_materia_buscada and id_comision = resultado_cursada.id_comision;
	end if;
end;
$$ language plpgsql;
