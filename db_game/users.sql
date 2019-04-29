CREATE TABLE users
(
  id                SERIAL PRIMARY KEY,
  name              VARCHAR(64),
  password          VARCHAR(255),
  mail              VARCHAR(64),
  credits           INT, /* внутреигровая валюта  */

  scientific_points INT,
  attack_points     INT,
  production_points INT,

  -- последняя база которую посетил игрок
  last_base_id      INT default 0 not null,

  /* этот инт показывает раздел обучения на котором остановился игрок, существует только ради обучения) */
  training          INT,
  --Replics Explores Reverses, нация за кторую играет игрок
  fraction          varchar(64),
  avatar            text, -- todo надо переделать на bytea забирать и обновлять методами decode(string text, format text) и encode(data bytea, format text) но мне лень
  biography         text,
  title             text
);

CREATE TABLE user_current_mission
(
  id         SERIAL PRIMARY KEY,
  id_user    INT REFERENCES users (id),
  id_mission INT REFERENCES missions (id),
  data       json -- тут текущее состояние квеста ¯\_(ツ)_/¯
)

-- todo линия выполнения заданий что бы старые задания оставались и не выдавались пока не пройдет весь цикл