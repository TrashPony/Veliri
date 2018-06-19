CREATE TABLE chassis_type (
	id serial primary key,
	name varchar(64),
	type varchar(64),
	carrying int,
	Maneuverability int,
	max_speed int
);

CREATE TABLE weapon_type (
	id serial primary key,
	name varchar(64),
	type varchar(64),
	weight int,
	damage int,
	min_attack_range int,
	range_attack int,
	accuracy int,
	area_covers int
);

CREATE TABLE tower_type (
	id serial primary key,
    name varchar(64),
	type varchar(64),
	weight	int,
	hp int,
	power_radar int,
	armor int,
	vulnerability_to_kinetics int,
	vulnerability_to_thermo int,
	vulnerability_to_em int,
	vulnerability_to_explosion int
);

CREATE TABLE body_type (
	id serial primary key,
	name varchar(64),
	type varchar(64),
	weight	int,
	hp int,
	max_tower_weight int,
	armor int,
	vulnerability_to_kinetics int,
	vulnerability_to_thermo int,
	vulnerability_to_em int,
	vulnerability_to_explosion int
);

CREATE TABLE radar_type (
	id serial primary key,
    name varchar(64),
	type varchar(64),
	weight	int,
	power int,
	through bool,
	analysis int
);