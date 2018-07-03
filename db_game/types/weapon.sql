CREATE TABLE weapon_type (
  id               SERIAL PRIMARY KEY,
  name             VARCHAR(64),
  min_attack_range INT,
  range_attack     INT,
  accuracy         INT,
  ammo_capacity    INT,                 /* кол-во боезапаса вмещаемого в орудие до перезарядки */
  artillery        BOOLEAN,
  power            INT                  /* кол-во потребляемой энергии */
);

CREATE TABLE ammunition_type (
  id               SERIAL PRIMARY KEY,
  name             VARCHAR(64),
  type             VARCHAR(64),
  type_attack      VARCHAR(64),
  damage           INT,
  area_covers      INT
);