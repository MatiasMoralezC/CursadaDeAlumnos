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
