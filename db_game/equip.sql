CREATE TABLE equipping_type (
    id serial primary key,
    type varchar(64),
    specification varchar(255), /* описание снаряжения */
    applicable varchar(64),     /* прменимо к своим "my_units", вражеским "hostile_units", всем "all",  карте "map"*/
    region int                  /* регион на которое работает снаряжение*/
);

CREATE TABLE equip_effects (
    id_equip int references equipping_type(id), /* ид снаряжения */
    id_effect int references effects_type(id)   /* ид эфекта которое накладывается снаряжением */
);