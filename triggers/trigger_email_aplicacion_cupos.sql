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
