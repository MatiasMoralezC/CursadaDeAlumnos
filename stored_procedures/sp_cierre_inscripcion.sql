create function cierreDeInscripcion(anio_ingresado int, nro_semestre_ingresado int) returns void as $$
declare
	resultado_periodo periodo%rowtype;
	semestre_buscado varchar(6);
begin
	semestre_buscado := to_char(anio_ingresado, 'FM999999') || '-' || to_char(nro_semestre_ingresado, 'FM999999');

	select * into resultado_periodo from periodo where semestre = semestre_buscado and estado = 'inscripcion' ;

	if not found then
	raise 'el semestre no existe en periodo de inscripcion';
	end if;

	update periodo set estado = 'cierre inscrip' where semestre = semestre_buscado;

end;
$$ language plpgsql;
