CREATE TABLE missions
(
  id   SERIAL PRIMARY KEY,
  name VARCHAR(64) not null default ''
);

CREATE TABLE actions
(
  id SERIAL PRIMARY KEY,
  type_monitor varchar(64)
)