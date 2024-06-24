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
