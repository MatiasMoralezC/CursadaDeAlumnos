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
