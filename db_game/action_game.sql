CREATE TABLE action_games (
  id     SERIAL PRIMARY KEY,
  name   VARCHAR(64),
  id_map INT REFERENCES maps (id),
  step   INT,
  phase  VARCHAR(64),
  winner VARCHAR(64)
);

CREATE TABLE action_game_unit (
  id               SERIAL PRIMARY KEY,

  /* Методанные об игре и владельце */
  id_user          INT REFERENCES users (id),
  id_game          INT REFERENCES action_games (id),

  /* Части юнита */
  id_chassis       INT REFERENCES chassis_type (id),
  id_weapons       INT REFERENCES weapon_type (id),
  id_tower         INT REFERENCES tower_type (id),
  id_body          INT REFERENCES body_type (id),
  id_radar         INT REFERENCES radar_type (id),

  /* Позиция */
  x                INT,
  y                INT,
  rotate           INT,
  on_map           BOOLEAN,

  /* Игровая статистика */
  action           BOOLEAN,
  target           VARCHAR(64),
  queue_attack     INT,

  /* Характиристики */
  weight           INT,
  speed            INT,
  initiative       INT,
  damage           INT,
  range_attack     INT,
  min_attack_range INT,
  area_attack      INT,
  type_attack      VARCHAR(64),
  max_hp           INT,
  hp               INT,
  armor            INT,
  evasion_critical INT,
  vul_kinetics     INT,
  vul_thermal      INT,
  vul_em           INT,
  vul_explosive    INT,
  range_view       INT,
  accuracy         INT,
  wall_hack        BOOLEAN
);

CREATE TABLE action_mother_ship (
  id      SERIAL PRIMARY KEY,
  id_game INT REFERENCES action_games (id),
  id_type INT REFERENCES mother_ship_type (id),
  id_user INT REFERENCES users (id),
  x       INT,
  y       INT
);

CREATE TABLE action_game_user (
  id_game INT REFERENCES action_games (id),
  id_user INT REFERENCES users (id),
  ready   BOOLEAN
);

CREATE TABLE action_game_equipping (/* снаряжение у игрока */
  id      SERIAL PRIMARY KEY,
  id_game INT REFERENCES action_games (id),
  id_user INT REFERENCES users (id),
  id_type INT REFERENCES equipping_type (id),
  used    BOOLEAN
);

CREATE TABLE action_game_unit_effects (/* эфекты которые в данный момент висят на юнитах */
  id         SERIAL PRIMARY KEY,
  id_unit    INT REFERENCES action_game_unit (id),
  id_effect  INT REFERENCES effects_type (id),
  left_steps INT
);

CREATE TABLE action_game_zone_effects (/* эфекты которые в данный момент висят на ячейках карты */
  id         SERIAL PRIMARY KEY,
  id_game    INT REFERENCES action_games (id),
  id_effect  INT REFERENCES effects_type (id),
  x          INT,
  y          INT,
  left_steps INT
);