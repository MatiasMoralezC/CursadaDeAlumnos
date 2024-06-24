create or replace function cierre_cursada(id_materia_buscada int, id_comision_buscada int, out p_result boolean, out p_error_message text) as $$
declare
	v_count int;
	v_semestre_actual text;
	v_nota_regular int;
	v_nota_final int;
begin

	/*--HARDCODEO
	insert into cursada
	values (1, 1, 1, current_timestamp, 8, 'ingresade');*/
	
	p_error_message := '';
	
	select semestre into v_semestre_actual from periodo where estado = 'cursada' limit 1;

	if v_semestre_actual is null then
		p_result := false;
		p_error_message := 'periodo de cursada cerrado';
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
	
	select count(*) into v_count from cursada where id_comision = id_comision_buscada;
	
	if v_count = 0 then
		p_result := false;
		p_error_message := 'comision sin alumnes inscriptes';
		return;
	end if;
	
	select count(*) into v_count from cursada where id_comision = id_comision_buscada and
	nota is not null;
	
	if v_count = 0 then
		p_result := false;
		p_error_message := 'la carga de notas no esta completa';
		return;
	end if;
	
	insert into historia_academica (id_alumne, semestre, id_materia, id_comision, estado, nota_regular, nota_final)
	select c.id_alumne, v_semestre_actual, c.id_materia, c.id_comision,
		case
			when c.nota is null then 'ausente'
			when c.nota between 1 and 3 then 'reprobada'
			when c.nota between 4 and 6 then 'regular'
			when c.nota between 7 and 10 then 'aprobada'
		end as estado,
		c.nota as nota_regular,
		case 
			when c.nota between 7 and 10 then c.nota
			else null
		end as nota_final
	from cursada c
	where c.id_comision = id_comision_buscada and
	c.id_materia = id_materia_buscada and
	c.estado = 'aceptade';
	
	delete from cursada
	where id_comision = id_comision_buscada and
	id_materia = id_materia_buscada;
	
	p_result := true;
	
end;
$$ language plpgsql;
