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
