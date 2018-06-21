CREATE TABLE chassis_type (
  id              SERIAL PRIMARY KEY,
  name            VARCHAR(64),
  type            VARCHAR(64),
  carrying        INT,
  Maneuverability INT,
  max_speed       INT
);

CREATE TABLE weapon_type (
  id               SERIAL PRIMARY KEY,
  name             VARCHAR(64),
  type             VARCHAR(64),
  weight           INT,
  damage           INT,
  min_attack_range INT,
  range_attack     INT,
  accuracy         INT,
  area_covers      INT
);

CREATE TABLE tower_type (
  id                         SERIAL PRIMARY KEY,
  name                       VARCHAR(64),
  type                       VARCHAR(64),
  weight                     INT,
  hp                         INT,
  power_radar                INT,
  armor                      INT,
  vulnerability_to_kinetics  INT,
  vulnerability_to_thermo    INT,
  vulnerability_to_em        INT,
  vulnerability_to_explosion INT
);

CREATE TABLE body_type (
  id                         SERIAL PRIMARY KEY,
  name                       VARCHAR(64),
  type                       VARCHAR(64),
  weight                     INT,
  hp                         INT,
  max_tower_weight           INT,
  armor                      INT,
  vulnerability_to_kinetics  INT,
  vulnerability_to_thermo    INT,
  vulnerability_to_em        INT,
  vulnerability_to_explosion INT
);

CREATE TABLE radar_type (
  id       SERIAL PRIMARY KEY,
  name     VARCHAR(64),
  type     VARCHAR(64),
  weight   INT,
  power    INT,
  through  BOOL,
  analysis INT
);