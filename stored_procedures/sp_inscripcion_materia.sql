create function inscripcion_materia(id_alumne_buscado integer, id_materia_buscada integer, id_comision_buscada integer, out p_result boolean, out p_error_message text) as $$
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
	p_error_message := '';

	select * into resultado_periodo from periodo where estado = 'inscripcion';
	
	if not found then
		insert into error values(nextval('error_id_seq'), 'alta inscrip', resultado_periodo.semestre, id_alumne_buscado, id_materia_buscada, id_comision_buscada, current_timestamp, 'Periodo de inscripción cerrado.');
		p_error_message := 'periodo de inscripción cerrado';
		p_result := false;
		return;
	end if;
	
	select * into resultado_alumne from alumne where id_alumne = id_alumne_buscado;
	
	if not found then
		insert into error values(nextval('error_id_seq'), 'alta inscrip', resultado_periodo.semestre, id_alumne_buscado, id_materia_buscada, id_comision_buscada, current_timestamp, 'Id de alumne no válido.');
		p_error_message := 'id de alumne no válido';
		p_result := false;
		return;
	end if;
	
	select * into resultado_materia from materia where id_materia = id_materia_buscada;
	
	if not found then
		insert into error values(nextval('error_id_seq'), 'alta inscrip', resultado_periodo.semestre, id_alumne_buscado, id_materia_buscada, id_comision_buscada, current_timestamp, 'Id de materia no válido.');
		p_error_message := 'id de materia no válido';
		p_result := false;
		return;
	end if;
	
	select * into resultado_comision from comision where id_materia = id_materia_buscada and id_comision = id_comision_buscada;
	
	if not found then
		insert into error values(nextval('error_id_seq'), 'alta inscrip', resultado_periodo.semestre, id_alumne_buscado, id_materia_buscada, id_comision_buscada, current_timestamp, 'Id de comision no válido.');
		p_error_message := 'id de comision no válido';
		p_result := false;
		return;
	end if;
	
	select * into resultado_cursada from cursada where id_alumne = id_alumne_buscado and id_materia = id_materia_buscada and id_comision = id_comision_buscada and estado = 'ingresade';
	
	if found then
		insert into error values(nextval('error_id_seq'), 'alta inscrip', resultado_periodo.semestre, id_alumne_buscado, id_materia_buscada, id_comision_buscada, current_timestamp, 'Alumne ya inscripte en la materia');
		p_error_message := 'alumne ya inscripte en la materia';
		p_result := false;
		return;
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
		insert into error values(nextval('error_id_seq'), 'alta inscrip', resultado_periodo.semestre, id_alumne_buscado, id_materia_buscada, id_comision_buscada, current_timestamp, 'Alumne no cumple requisitos de correlatividad');
		p_error_message := 'alumne no cumple requisitos de correlatividad';
		p_result := false;
		return;
	end if;
	
	insert into cursada values(id_materia_buscada, id_alumne_buscado, id_comision_buscada, current_timestamp, null, 'ingresade');
	
	p_result := true;
end;
$$ language plpgsql;
