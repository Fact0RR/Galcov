CREATE OR REPLACE FUNCTION add_user(p_login text, p_password text)
RETURNS int
language plpgsql
as $$
DECLARE id_user int;
begin
    insert into users (login,password) values (p_login,p_password)
    RETURNING id into id_user;
    return id_user;
end;
$$;

CREATE OR REPLACE FUNCTION check_user(p_login text, p_password text)
RETURNS int
language plpgsql
as $$
DECLARE 
    id_user int;
begin
	
	SELECT case when sum(id) is null then 0 else sum(id) end into id_user from users where login = p_login and password = p_password;

    return id_user;
end;
$$;

