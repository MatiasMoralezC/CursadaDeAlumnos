create function baja_de_inscripcion(id_alumne_buscade integer, id_materia_buscada integer) returns void as $$
declare
	resultado_periodo periodo%rowtype;
	resultado_alumne alumne%rowtype;
	resultado_materia materia%rowtype;
	resultado_comision comision%rowtype;
	resultado_cursada cursada%rowtype;

	alumne_enespera record;

begin
	select * into resultado_periodo from periodo where estado = 'inscripcion' or estado = 'cursada';

	if not found then
	raise 'no se permiten bajas en este periodo';
	end if;

	select * into resultado_alumne from alumne where id_alumne = id_alumne_buscade;

	if not found then
	raise 'id de alumne no válido';
	end if;

	select * into resultado_materia from materia where id_materia = id_materia_buscada;

	if not found then
	raise 'id de materia no válido';
	end if;

	select * into resultado_cursada from cursada where id_alumne = id_alumne_buscade and id_materia = id_materia_buscada and estado = 'aceptade';

	if not found then
	raise 'alumne no inscripte en la materia';
	end if;

	update cursada set estado = 'dade de baja' where cursada.id_alumne = id_alumne_buscade and cursada.id_materia = id_materia_buscada;
	
	if resultado_periodo.estado = 'cursada' then
		select * into alumne_enespera from cursada 
		where id_materia = id_materia_buscada and id_comision = resultado_cursada.id_comision and estado = 'en espera' 
		order by f_inscripcion asc limit 1;
		if not found then
			raise 'no hay alumne en espera';
		end if;
		
		update cursada set estado = 'aceptade' 
		where id_alumne = alumne_enespera.id_alumne and id_materia = id_materia_buscada and id_comision = resultado_cursada.id_comision;
	end if;
end;
$$ language plpgsql;
