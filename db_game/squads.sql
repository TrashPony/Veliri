 CREATE TABLE squads (
  id      SERIAL PRIMARY KEY,
  name    VARCHAR(64),
  active  BOOLEAN,                    /* у игрока может быть в 1 момент времени только 1 отряд, это переменная говорит какой */
  id_user INT REFERENCES users (id),  /* кому принадлежит */
  in_game BOOLEAN                     /* отряд в бою */
);

CREATE TABLE squad_units (
  id                  SERIAL PRIMARY KEY,
  id_squad            INT REFERENCES squads (id),          /* ид отряда к которому принадлежит юнит */

  /* из чего состоит юнит */
  id_body             INT REFERENCES body_type (id),       /* ид тела юнита */

  slot                INT, /* номер слота который занимает юнит в материнской машине */

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
  type          VARCHAR(64),                          /* оружие(weapon), снаряжение(equip) или боеприпасы (ammo)*/
  id_squad_unit INT REFERENCES squad_units,           /* ид юнита к которому прикреплено снаряжение */
  id_equipping  INT REFERENCES equipping_type (id),   /* ид снаряжения */
  slot_in_body  INT                                   /* слот который занимает снаряжения, тип слота определяется типом слота снаряжения */
);

CREATE TABLE squad_mother_ship_equipping ( /* таблица снаряжения которое нацеплино на мазер шипа */
  id                   SERIAL PRIMARY KEY,
  id_squad             INT REFERENCES squads (id),
  id_squad_mother_ship INT REFERENCES squad_mother_ship,   /* ид мазершипа к которому прикреплено снаряжение */
  type                 VARCHAR(64),                        /* оружие(weapon), снаряжение(equip) или боеприпасы (ammo)*/
  id_equipping         INT REFERENCES equipping_type (id), /* ид снаряжения */
  slot_in_body         INT                                 /* слот который занимает снаряжения, тип слота определяется типом слота снаряжения */
);

CREATE TABLE squad_inventory (                     /* инвентарь отряда не боевой параметр */
  id                   SERIAL PRIMARY KEY,
  id_squad             INT REFERENCES squads (id), /* какому отряду принаджелит */
  slot                 INT,                        /* какой слот занимает итем */
  item_type            VARCHAR(64),                /* оружие(weapon), снаряжение(equip) или боеприпасы (ammo), корпуса (body) */
  item_id              INT,                        /* ид итема определяет конкретный итем тип + ид*/
  quantity             INT                         /* количество предметов в слоте */
);