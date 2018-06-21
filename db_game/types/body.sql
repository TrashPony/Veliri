CREATE TABLE body_type (
  id                         SERIAL PRIMARY KEY,
  name                       VARCHAR(64),
  mother_ship                BOOLEAN,
  speed                      INT,
  initiative                 INT,
  max_hp                     INT,
  armor                      INT,
  evasion_critical           INT,
  vulnerability_to_kinetics  INT,
  vulnerability_to_thermo    INT,
  vulnerability_to_em        INT,
  vulnerability_to_explosion INT,
  range_view                 INT,
  accuracy                   INT,
  power                      INT, /* сила встроеного генератора */
  wall_hack                  BOOLEAN
);

CREATE TABLE body_slots (
  id_body     INT REFERENCES body_type (id), /* ид корпуса которому принадлежит слот*/
  type_slot   INT,                           /* тип слота куда встаривается оборудование I (1) , II (2), III (3), IV (4), V (5) */
  number_slot INT,                           /* номер слота в корпусе */
  weapon      BOOLEAN,                       /* труе если слот приналдлежит оружию */
  weapon_type varchar(64)                    /* тип оружия например "artillery" */
);