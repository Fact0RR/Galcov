
CREATE TABLE USERS(
  id serial PRIMARY KEY,
  login text NOT NULL,
  password text NOT NULL
);


CREATE TABLE VAL(
   id serial PRIMARY KEY,
   name varchar(127) NOT NULL,
   price numeric NOT NULL,
   description text
);

CREATE TABLE ACC_VAL(
    id serial PRIMARY KEY,
    login_id int NOT NULL,
    val_id int NOT NULL,
    count numeric NOT NULL,

    UNIQUE (login_id, val_id),

    FOREIGN KEY (login_id) REFERENCES USERS (id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (val_id) REFERENCES VAL (id) ON DELETE CASCADE ON UPDATE CASCADE
);
