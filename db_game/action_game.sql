CREATE TABLE action_games(
   id serial primary key,
   name varchar(64),
   id_map int references maps(id),
   step int,
   phase varchar(64),
   winner varchar(64)
);

CREATE TABLE action_game_unit(
	id serial primary key,

	/* Методанные об игре и владельце */
	id_user int references users(id),
	id_game int references action_games(id),

	/* Части юнита */
	id_chassis  int references chassis_type(id),
	id_weapons  int references weapon_type(id),
	id_tower    int references tower_type(id),
	id_body     int references body_type(id),
	id_radar    int references radar_type(id),

	/* Позиция */
	x       int,
	y       int,
	rotate  int,
	on_map  boolean,

	/* Игровая статистика */
	action        boolean,
	target        varchar(64),
	queue_attack  int,

	/* Характиристики */
	weight           int,
    speed            int,
    initiative       int,
    damage           int,
    range_attack     int,
    min_attack_range int,
    area_attack      int,
    type_attack      varchar(64),
    max_hp           int,
    hp               int,
    armor            int,
    evasion_critical int,
    vul_kinetics     int,
    vul_thermal      int,
    vul_em           int,
    vul_explosive    int,
    range_view       int,
    accuracy         int,
    wall_hack        boolean
);

CREATE TABLE  action_mother_ship(
    id serial primary key,
    id_game int references action_games(id),
    id_type int references mother_ship_type(id),
    id_user int references users(id),
    x       int,
    y       int
);

CREATE TABLE action_game_user(
   id_game  int references action_games(id),
   id_user  int references users(id),
   ready    boolean
);

CREATE TABLE action_game_equipping( /* снаряжение у игрока */
   id           serial primary key,
   id_game      int references action_games(id),
   id_user      int references users(id),
   id_type      int references equipping_type(id),
   used         boolean
);

CREATE TABLE action_game_unit_effects( /* эфекты которые в данный момент висят на юнитах */
   id serial    primary key,
   id_unit      int references action_game_unit(id),
   id_effect    int references effects_type(id),
   left_steps   int
);

CREATE TABLE action_game_zone_effects( /* эфекты которые в данный момент висят на ячейках карты */
   id serial    primary key,
   id_game      int references action_games(id),
   id_effect    int references effects_type(id),
   x            int,
   y            int,
   left_steps   int
);