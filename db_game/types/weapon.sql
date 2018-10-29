CREATE TABLE weapon_type (
  id               SERIAL PRIMARY KEY,
  name             VARCHAR(64),
  min_attack_range INT,
  range_attack     INT,
  accuracy         INT,
  ammo_capacity    INT,                 /* кол-во боезапаса вмещаемого в орудие до перезарядки */
  artillery        BOOLEAN,             /* параметр отвечает за игнорирование препятвий */
  power            INT,                 /* кол-во потребляемой энергии */
  max_hp           INT,                 /* макс кол-во хп */
  type             VARCHAR(64),         /* firearms, missile_weapon, laser_weapon */
  standard_size    INT,                 /* small - 1, medium - 2, big - 3 */
  size             REAL,                /* занимаемый обьем в кубо метрах */
  initiative       INT                  /* инициаива, определяет порядок действия в фазе атаки */
);

CREATE TABLE ammunition_type (
  id               SERIAL PRIMARY KEY,
  name             VARCHAR(64),
  type             VARCHAR(64),         /* определяет к какому оружию подойдет оружие firearms, missile_weapon, laser_weapon */
  standard_size    INT,                 /* small - 1, medium - 2, big - 3 */
  type_attack      VARCHAR(64),
  min_damage       INT,
  max_damage       INT,
  area_covers      INT,
  size             REAL
);