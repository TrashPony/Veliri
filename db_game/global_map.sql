CREATE TABLE bases (
  id      SERIAL PRIMARY KEY,
  id_user INT REFERENCES users (id),
  name    VARCHAR(64),
  type    VARCHAR(64),
  x       INT,
  y       INT
);

CREATE TABLE structures (
  id      SERIAL PRIMARY KEY,
  id_base INT REFERENCES users (id),
  type_id INT,
  lvl     INT
);

CREATE TABLE mather_ships (
  id      SERIAL PRIMARY KEY,
  id_user INT REFERENCES users (id),
  id_type INT REFERENCES mother_ship_type (id)
);

CREATE TABLE user_parameters (
  id         SERIAL PRIMARY KEY,
  id_user    INT REFERENCES users (id),
  experience INT (64),
  lvl        INT (3)
);

CREATE TABLE structure_type (
  id SERIAL PRIMARY KEY
);

CREATE TABLE global_map (
  id SERIAL PRIMARY KEY
);

CREATE TABLE global_map (
  id          SERIAL PRIMARY KEY,
  id_map      INT REFERENCES global_map (id),
  X           INT,
  Y           INT,
  type        VARCHAR(64),
  Object_type VARCHAR(64),
  id_user     INT REFERENCES users (id)
);

CREATE TABLE structure_type_in_map (
  id SERIAL PRIMARY KEY
);


