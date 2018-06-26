 CREATE TABLE squads (
  id      SERIAL PRIMARY KEY,
  name    VARCHAR(64),
  id_user INT REFERENCES users (id),  /* кому принадлежит */
  in_game BOOLEAN                     /* отряд в бою */
);

CREATE TABLE squad_units (
  id                  SERIAL PRIMARY KEY,
  id_squad            INT REFERENCES squads (id),          /* ид отряда к которому принадлежит юнит */

  /* из чего состоит юнит */
  id_weapon           INT REFERENCES weapon_type (id),     /* ид оружия которое на юните */
  id_ammunition       INT REFERENCES ammunition_type (id), /* ид боеприпаса которое на юните */
  id_body             INT REFERENCES body_type (id),       /* ид тела юнита */

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

  /* из чего состоит мазер шип */
  id_weapon     INT REFERENCES weapon_type (id),     /* ид оружия которое на юните */
  id_ammunition INT REFERENCES ammunition_type (id), /* ид боеприпаса которое на юните */
  id_body       INT REFERENCES body_type (id),       /* ид тела юнита */

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

CREATE TABLE squad_units_equipping ( /* таблица снаряжения которое нацеплино на юнита */
  id            SERIAL PRIMARY KEY,
  id_squad      INT REFERENCES squads (id),

  id_squad_unit INT REFERENCES squad_units,           /* ид юнита к которому прикреплено оружие */
  id_equipping  INT REFERENCES equipping_type (id),   /* ид снаряжения */
  slot_in_unit  INT                                   /* слот который занимает снаряжения, тип слота определяется типом слота снаряжения */
);

CREATE TABLE squad_mother_ship_equipping ( /* таблица снаряжения которое нацеплино на мазер шипа */
  id                   SERIAL PRIMARY KEY,
  id_squad             INT REFERENCES squads (id),
  id_squad_mother_ship INT REFERENCES squad_mother_ship,   /* ид мазершипа к которому прикреплено оружие */
  id_equipping         INT REFERENCES equipping_type (id), /* ид снаряжения */
  slot_in_mother_ship  INT                                 /* слот который занимает снаряжения, тип слота определяется типом слота снаряжения */
);

CREATE TABLE squad_inventory (                     /* инвентарь отряда не боевой параметр */
  id                   SERIAL PRIMARY KEY,
  id_squad             INT REFERENCES squads (id), /* какому отряду принаджелит */
  slot                 INT,                        /* какой слот занимает итем */
  item_type            VARCHAR(64),                /* тип итема (снаряжение ресурсы патроны и тд)*/
  item_id              INT                         /* ид итема определяет конкретный итем тип + ид*/
);