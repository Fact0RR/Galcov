CREATE OR REPLACE PROCEDURE t2s(id_user int, id_val_old int, id_val_new int, p_count numeric)
LANGUAGE plpgsql
AS $$
begin


    IF EXISTS (SELECT login_id, val_id FROM ACC_VAL where login_id = id_user and val_id = id_val_old and p_count <= count) THEN
        
        IF EXISTS (SELECT login_id, val_id FROM ACC_VAL where login_id = id_user and val_id = id_val_new) THEN
            
            update ACC_VAL set count = count - p_count where login_id = id_user and val_id = id_val_old;
            update ACC_VAL set count = count + p_count * (select price from val where id = id_val_old)/(select price from val where id = id_val_new) where login_id = id_user and val_id = id_val_new;

        ELSE

            insert into ACC_VAL (login_id,val_id,count) values (id_user,id_val_new,0);
            update ACC_VAL set count = count - p_count where login_id = id_user and val_id = id_val_old;
            update ACC_VAL set count = count + p_count * (select price from val where id = id_val_old)/(select price from val where id = id_val_new) where login_id = id_user and val_id = id_val_new;

        END IF;

    END IF;
    commit;
end;
$$;


CREATE OR REPLACE PROCEDURE t2u(id_user_old int,id_user_new int, id_val int, p_count numeric)
LANGUAGE plpgsql
AS $$
begin
    
    IF EXISTS (SELECT login_id, val_id FROM ACC_VAL where login_id = id_user_old and val_id = id_val and p_count <= count) THEN

        IF EXISTS (SELECT login_id, val_id FROM ACC_VAL where login_id = id_user_new and val_id = id_val) THEN
        
            update ACC_VAL set count = count - p_count where login_id = id_user_old and val_id = id_val;
            update ACC_VAL set count = count + p_count where login_id = id_user_new and val_id = id_val;

        ELSE

            insert into ACC_VAL (login_id,val_id,count) values (id_user_new,id_val,0);
            update ACC_VAL set count = count - p_count where login_id = id_user_old and val_id = id_val;
            update ACC_VAL set count = count + p_count where login_id = id_user_new and val_id = id_val;
        END IF;
    END IF;


end;
$$;