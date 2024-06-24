create function cierreDeInscripcion(semestre_buscado text) returns void as $$
declare
	resultado_periodo periodo%rowtype;

begin

	select * into resultado_periodo from periodo where semestre = semestre_buscado and estado = 'inscripcion' ;

	if not found then
	raise 'el semestre no existe en periodo de inscripcion';
	end if;

	update periodo set estado = 'cierre inscrip' where semestre = semestre_buscado;

end;
$$ language plpgsql;
