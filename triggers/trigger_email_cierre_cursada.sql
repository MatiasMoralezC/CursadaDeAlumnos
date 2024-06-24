create or replace function email_cierre_cursada() returns trigger as $$
declare
	v_materia_nombre text;
	v_comision_numero text;
	v_alumne_nombre text;
	v_alumne_apellido text;
	v_email_alumne text;
	v_estado_academico char(15);
	v_nota_regular int;
	v_nota_final int;
	v_semestre_actual text;
begin
	select nombre into v_materia_nombre from materia where id_materia = new.id_materia;
	select id_comision into v_comision_numero from comision where id_materia = new.id_materia and id_comision = new.id_comision;
	select nombre, apellido, email into v_alumne_nombre, v_alumne_apellido, v_email_alumne from alumne where id_alumne = new.id_alumne;
	select semestre into v_semestre_actual from periodo where estado = 'cursada';
	select estado, nota_regular, nota_final into v_estado_academico, v_nota_regular, v_nota_final from historia_academica
	where id_alumne = NEW.id_alumne and id_materia = NEW.id_materia and semestre = v_semestre_actual;
	
	insert into envio_email
	values (nextval('envio_email_id_seq'), current_timestamp, 'Cierre de cursada',
	'Hola ' || v_alumne_nombre || ' ' || v_alumne_apellido || ', tu inscripcion a la materia ' || v_materia_nombre || ', comision ' || v_comision_numero ||
	' ha sido cerrada. Estado: ' || v_estado_academico || ', Nota regular: ' || coalesce(v_nota_regular::text, 'n/a') || ', Nota final: ' || coalesce(v_nota_final::text, 'n/a') || '.',
	null, 'pendiente'
	);
		
	return new;
end;
	
$$ language plpgsql;
	
create trigger email_cierre_cursada_trg
after update on cursada
for each row
when (new.estado = 'cerrado')
execute function email_cierre_cursada();
