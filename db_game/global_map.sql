CREATE TABLE bases (
    id serial primary key,
    id_user int references users(id),
    name varchar(64),
    type varchar(64),
    x int,
    y int
);

CREATE TABLE structures (
    id serial primary key,
    id_base int references users(id),
    type_id int,
    lvl int
);

CREATE TABLE mather_ships (
    id serial primary key,
    id_user int references users(id),
    id_type int references mother_ship_type(id)
);

CREATE TABLE user_parameters (
    id serial primary key,
    id_user int references users(id),
    experience int(64),
    lvl int(3)
);

CREATE TABLE structure_type (
    id serial primary key
);

CREATE TABLE global_map (
    id serial primary key
);

CREATE TABLE global_map (
    id serial primary key,
    id_map int references global_map(id),
    X int,
    Y int,
    type varchar(64),
    Object_type varchar(64),
    id_user int references users(id)
);

CREATE TABLE structure_type_in_map (
    id serial primary key
);

