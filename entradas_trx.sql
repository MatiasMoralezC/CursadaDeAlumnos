\c garcia_montoro_moralez_rodriguez_db1

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
			perform apertura_inscripcion(v_año, v_nro_semestre);
		end if;
		
		if v_operacion = 'alta inscrip' then
			perform inscripcion_materia(v_id_alumne, v_id_materia, v_id_comision);
		end if;
		
		if v_operacion = 'cierre inscrip' then
			perform cierreDeInscripcion(v_año, v_nro_semestre);
		end if;
		
		if v_operacion = 'aplicacion cupo' then
			perform aplicacion_cupos(v_año, v_nro_semestre);
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
		
	
