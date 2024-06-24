create function aplicacion_cupos(anio_ingresado int, nro_semestre_ingresado int) returns void as $$
declare
	periodo_encontrado periodo%rowtype;
	cupo_materia int;
	id_materia_buscada int := 1;
	id_comision_buscada int;
	alumne_inscripte cursada%rowtype;
	comision_materia comision%rowtype;
	materia comision%rowtype;
	semestre_buscado varchar(6);
begin
	semestre_buscado := to_char(anio_ingresado, 'FM999999') || '-' || to_char(nro_semestre_ingresado, 'FM999999');

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
