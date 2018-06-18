CREATE TABLE users(
    id serial primary key,
    name varchar(64),
    password varchar(255),
    mail varchar(64)
);