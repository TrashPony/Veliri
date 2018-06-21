CREATE TABLE mother_ship_type (
  id              SERIAL PRIMARY KEY,
  type            VARCHAR(64),
  hp              INT,
  armor           INT,
  unit_slots      INT,
  unit_slot_size  INT,
  equipment_slots INT,
  range_view      INT
);

CREATE TABLE squads (
  id      SERIAL PRIMARY KEY,
  name    VARCHAR(64),
  id_user INT REFERENCES users (id)
);

CREATE TABLE squad_units (
  id                  SERIAL PRIMARY KEY,
  id_squad            INT REFERENCES squads (id),
  slot_in_mother_ship INT,
  id_chassis          INT REFERENCES chassis_type (id),
  id_weapon           INT REFERENCES weapon_type (id),
  id_tower            INT REFERENCES tower_type (id),
  id_body             INT REFERENCES body_type (id),
  id_radar            INT REFERENCES radar_type (id)
);

CREATE TABLE squad_mother_ship (
  id             SERIAL PRIMARY KEY,
  id_squad       INT REFERENCES squads (id),
  id_mother_ship INT REFERENCES mother_ship_type (id)
);

CREATE TABLE squad_equipping (
  id                  SERIAL PRIMARY KEY,
  id_squad            INT REFERENCES squads (id),
  slot_in_mother_ship INT,
  id_equipping        INT REFERENCES equipping_type (id)
);