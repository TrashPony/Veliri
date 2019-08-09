CREATE TABLE action_games (
  id       SERIAL PRIMARY KEY,
  name     VARCHAR(64),
  id_map   INT REFERENCES maps (id), /* карта боя */
  step     INT,                      /* шаг боя */
  phase    VARCHAR(64),              /* фаза боя */
  winner   VARCHAR(64),             /* победитель */
  end_game boolean
);

CREATE TABLE action_game_squads (              /* отряды которые участвую в игре */
  id_game  INT REFERENCES action_games (id),
  id_squad INT REFERENCES squads (id)
);

CREATE TABLE action_game_user (                       /* пользователи которые участвуют в игре */
  id_game        INT REFERENCES action_games (id),
  id_user        INT REFERENCES users (id),
  ready          BOOLEAN,                             /* готовность пользователя */
  leave          BOOLEAN                              -- игрок ливнул из игры по той или иной причине
);

-- таблица которая хранит текущие состояния союзов игроков, связи многие ко многим
CREATE TABLE action_game_pacts (
  id_game        INT REFERENCES action_games (id),
  id_user        INT REFERENCES users (id),
  id_to_user     INT REFERENCES users (id)
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
  q          INT,
  r          INT,
  left_steps INT                               /* сколько шагов ему еще висеть */
);

CREATE TABLE user_memory_unit (
  id      SERIAL PRIMARY KEY,
  id_user INT,
  id_game INT,
  id_unit INT,
  unit    JSON                                 /* хранит текущие состояние юнита в виде строки json */
);

-- таблица хранит в себе юнитов игроков которые вышли из игры
CREATE TABLE game_leave_unit (
  id      SERIAL PRIMARY KEY,
  id_user INT REFERENCES users (id),
  id_game INT REFERENCES action_games (id),
  unit    JSON                                 /* хранит текущие состояние юнита в виде строки json */
);