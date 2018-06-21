CREATE TABLE squads (
  id      SERIAL PRIMARY KEY,
  name    VARCHAR(64),
  id_user INT REFERENCES users (id),
  in_game BOOLEAN
);

CREATE TABLE squad_units (
  id                  SERIAL PRIMARY KEY,
  id_squad            INT REFERENCES squads (id),
  id_weapon           INT REFERENCES weapon_type (id),
  id_ammunition       INT REFERENCES ammunition_type (id),
  id_body             INT REFERENCES body_type (id),

  slot_in_mother_ship INT, /* номер слота который занимает юнит в материнской машине */

  /* Позиция */
  x                   INT,
  y                   INT,
  rotate              INT,
  on_map              BOOLEAN,

  /* Игровая статистика */
  action              BOOLEAN,
  target              VARCHAR(64),
  queue_attack        INT,

  /* Характиристики */
  hp                  INT
);

CREATE TABLE squad_mother_ship (
  id            SERIAL PRIMARY KEY,
  id_squad      INT REFERENCES squads (id),
  id_weapon     INT REFERENCES weapon_type (id),
  id_ammunition INT REFERENCES ammunition_type (id),
  id_body       INT REFERENCES body_type (id),

  /* Позиция */
  x             INT,
  y             INT,
  rotate        INT,

  /* Игровая статистика */
  action        BOOLEAN,
  target        VARCHAR(64),
  queue_attack  INT,

  /* Характиристики */
  hp            INT
);

CREATE TABLE squad_units_equipping (
  id            SERIAL PRIMARY KEY,
  id_squad      INT REFERENCES squads (id),
  id_squad_unit INT REFERENCES squad_units,
  id_equipping  INT REFERENCES equipping_type (id),
  slot_in_unit  INT
);

CREATE TABLE squad_mother_ship_equipping (
  id                   SERIAL PRIMARY KEY,
  id_squad             INT REFERENCES squads (id),
  id_squad_mother_ship INT REFERENCES squad_mother_ship,
  id_equipping         INT REFERENCES equipping_type (id),
  slot_in_mother_ship  INT
);

CREATE TABLE squad_inventory (
  id                   SERIAL PRIMARY KEY,
  id_squad             INT REFERENCES squads (id),
  slot                 INT
);