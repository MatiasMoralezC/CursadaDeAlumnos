\c garcia_montoro_moralez_rodriguez_db1

/*Hardcodeado*/
insert into entrada_trx values(1, 'apertura', 2023, 2, null, null, null, null);
insert into entrada_trx values(2, 'alta inscrip', null, null, 10, 3, 1, null);
insert into entrada_trx values(3, 'cierre inscrip', 2024, 1, null, null, null, null);
insert into entrada_trx values(4, 'aplicacion cupo', 2024, 1, null, null, null, null);
insert into entrada_trx values(5, 'ingreso nota', null, null, 10, 3, 1, 10);
/*insert into entrada_trx values(6, 'cierre cursada', null, null, null, 3, 1, null);*/


create function ejecutar_entradas() returns void as $$
	declare
		v_operacion text;
		v_año int;
		v_nro_semestre int;
		v_id_alumne int;
		v_id_materia int;
		v_id_comision int;
		v_nota int;
		v_id_orden int := 1;
	begin
		loop
		
		perform 1 from entrada_trx where id_orden = v_id_orden;
		exit when not found;
				
		select operacion, año, nro_semestre, id_alumne, id_materia, id_comision, nota into v_operacion, v_año, v_nro_semestre, v_id_alumne, v_id_materia, v_id_comision, v_nota
		from entrada_trx where id_orden = v_id_orden;
		
		if v_operacion = 'apertura' then
			perform apertura_inscripcion('2025-1');
		end if;
		
		if v_operacion = 'alta inscrip' then
			perform inscripcion_materia(v_id_alumne, v_id_materia, v_id_comision);
		end if;
		
		if v_operacion = 'cierre inscrip' then
			perform cierreDeInscripcion('2025-1');
		end if;
		
		if v_operacion = 'aplicacion cupo' then
			perform aplicacion_cupos('2025-1');
		end if;
		
		if v_operacion = 'ingreso nota' then
			perform ingreso_nota(v_id_alumne, v_id_materia, v_id_comision, v_nota);
		end if;
		
		/*if v_operacion = 'cierre cursada' then
			perform inscripcion_materia(v_id_alumne, v_id_materia, v_id_comision);
		end if;*/
		
		v_id_orden := v_id_orden + 1;
		
		end loop;
	end;
	$$ language plpgsql;
	
select ejecutar_entradas();
		
	
