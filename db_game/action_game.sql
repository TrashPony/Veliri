CREATE TABLE action_games (
  id     SERIAL PRIMARY KEY,
  name   VARCHAR(64),
  id_map INT REFERENCES maps (id),
  step   INT,
  phase  VARCHAR(64),
  winner VARCHAR(64)
);

CREATE TABLE action_game_squads (
  id_game  INT REFERENCES action_games (id),
  id_squad INT REFERENCES squads (id)
);

CREATE TABLE action_game_user (
  id_game INT REFERENCES action_games (id),
  id_user INT REFERENCES users (id),
  ready   BOOLEAN
);

CREATE TABLE action_game_unit_effects (/* эфекты которые в данный момент висят на юнитах */
  id         SERIAL PRIMARY KEY,
  id_unit    INT REFERENCES squad_units (id),
  id_squad   INT REFERENCES squads (id),
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

CREATE TABLE action_game_reload_equip (
  id_squad_equip    INT REFERENCES squad_units_equipping (id), /* ид эквипа в отряде */
  reload            INT                                        /* сколько он еще будет перезаряжаться если 0 то он заряжен */
)