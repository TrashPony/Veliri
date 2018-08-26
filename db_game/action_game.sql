CREATE TABLE action_games (
  id     SERIAL PRIMARY KEY,
  name   VARCHAR(64),
  id_map INT REFERENCES maps (id), /* карта боя */
  step   INT,                      /* шаг боя */
  phase  VARCHAR(64),              /* фаза боя */
  winner VARCHAR(64)               /* победитель */
);

CREATE TABLE action_game_squads (              /* отряды которые участвую в игре */
  id_game  INT REFERENCES action_games (id),
  id_squad INT REFERENCES squads (id)
);

CREATE TABLE action_game_user (                /* пользователи которые участвуют в игре */
  id_game INT REFERENCES action_games (id),
  id_user INT REFERENCES users (id),
  ready   BOOLEAN                              /* готовность пользователя */
);

CREATE TABLE action_game_unit_effects (        /* эфекты которые в данный момент висят на юнитах */
  id         SERIAL PRIMARY KEY,
  id_unit    INT REFERENCES squad_units (id),  /* на каком юните висит эффект */
  id_squad   INT REFERENCES squads (id),
  id_effect  INT REFERENCES effects_type (id), /* какой эффект */
  left_steps INT                               /* сколько шагов ему еще висеть */
);

CREATE TABLE action_game_zone_effects (        /* эфекты которые в данный момент висят на ячейках карты */
  id         SERIAL PRIMARY KEY,
  id_game    INT REFERENCES action_games (id),
  id_effect  INT REFERENCES effects_type (id), /* какой эффект */
  x          INT,
  y          INT,
  left_steps INT                               /* сколько шагов ему еще висеть */
);