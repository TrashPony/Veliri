CREATE TABLE users (
  id               SERIAL PRIMARY KEY,
  name             VARCHAR(64),
  password         VARCHAR(255),
  mail             VARCHAR(64),
  credits          INT, /* внутреигровая валюта  */
  experience_point INT, /* накопленые очки опыта */

  /* этот инт показывает раздел обучения на котором остановился игрок, существует только ради обучения) */
  training         INT
);