CREATE TABLE equipping_type (
  id            SERIAL PRIMARY KEY,
  type          VARCHAR(64),
  specification VARCHAR(255), /* описание снаряжения */
  applicable    VARCHAR(64), /* прменимо к своим "my_units", вражеским "hostile_units", всем "all",  карте "map"*/
  region        INT                  /* регион на которое работает снаряжение*/
);

CREATE TABLE equip_effects (
  id_equip  INT REFERENCES equipping_type (id), /* ид снаряжения */
  id_effect INT REFERENCES effects_type (id)   /* ид эфекта которое накладывается снаряжением */
);