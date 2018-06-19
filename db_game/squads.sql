CREATE TABLE mother_ship_type (
    id serial primary key,
    type varchar(64),
    hp int,
    armor int,
    unit_slots int,
    unit_slot_size int,
    equipment_slots int,
    range_view int
);

CREATE TABLE squads (
	id serial primary key,
	name varchar(64),
	id_user int references users(id)
);

CREATE TABLE squad_units (
	id serial primary key,
	id_squad int references squads(id),
	slot_in_mother_ship int,
	id_chassis int references chassis_type(id),
	id_weapon int references weapon_type(id),
	id_tower int references tower_type(id),
	id_body int references body_type(id),
	id_radar int references radar_type(id)
);

CREATE TABLE squad_mother_ship (
    id serial primary key,
	id_squad int references squads(id),
	id_mother_ship int references mother_ship_type(id)
);

CREATE TABLE squad_equipping (
	id serial primary key,
	id_squad int references squads(id),
	slot_in_mother_ship int,
    id_equipping int references equipping_type(id)
);