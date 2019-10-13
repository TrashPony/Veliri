CREATE TABLE weapon_type
(
  id                    SERIAL PRIMARY KEY,
  name                  VARCHAR(64),
  min_attack_range      INT,
  range_attack          INT,
  accuracy              INT,
  ammo_capacity         INT, /* кол-во боезапаса вмещаемого в орудие до перезарядки */
  artillery             BOOLEAN, /* параметр отвечает за игнорирование препятвий (артилерия, тактическая ракета) */
  power                 INT, /* кол-во потребляемой энергии в реакторе */
  max_hp                INT, /* макс кол-во хп */
  type                  VARCHAR(64), /* firearms, missile, laser */
  standard_size         INT, /* small - 1, medium - 2, big - 3 */
  size                  REAL, /* занимаемый обьем в кубо метрах */
  equip_damage          INT, /* урон по эквипу */
  equip_critical_damage INT, /* шанс нанести критический урон по эквипу  */
  rotate_speed          INT, /* скорость поворота башни */
  /* скорость полета пути, складывается с скорость из ammunition_type т.к. для танкового оружия будет задавать скорость оружие а для ракетницы снаряды  */
  bullet_speed          INT,
  reload_time           INT, /* время перезарядки в секудах */

  y_attach              INT, /* точка крепления оружия к точке крепления корпуса, по сути центр оружия */
  x_attach              INT, /* точка крепления оружия к точке крепления корпуса, по сути центр оружия */

  count_fire_bullet     INT, /* определяет сколько пуль будет выпущено за выстрел */

  /* если count_fire_bullet > 1 то это параметр указывает задержку между залпали, если 0 то выстрел происходит одновременно из всех стволов */
  delay_following_fire  INT,
  /* [ {"x": 1, "y": 1}, {"x": 2, "y": 2} ] описывают точки появления снарядов при запуске относительно оружия */
  fire_positions        json
);

CREATE TABLE ammunition_type
(
  id                    SERIAL PRIMARY KEY,
  name                  VARCHAR(64), /* спрайти патрона для оторажения в игре */
  type                  VARCHAR(64), /* определяет к какому оружию подойдет оружие firearms, missile_weapon, laser_weapon */
  standard_size         INT, /* small - 1, medium - 2, big - 3 */
  type_attack           VARCHAR(64),
  min_damage            INT,
  max_damage            INT,
  area_covers           INT, /* зона покрытия уроном в пикселях - радиус */
  equip_damage          INT, /* дополнительный урон по эквипу складывается с оружейным */
  equip_critical_damage INT, /* дополнительный шанс нанести критический урон по эквипу складывается с оружейным */
  size                  REAL,
  chase_target          BOOLEAN, /* снаряд преследует цель. Нарпмер самонаводящиеся ракета */
  /* скорость полета пути, складывается с скорость из weapon_type т.к. для танкового оружия будет задавать скорость оружие а для ракетницы снаряды  */
  bullet_speed          INT
);