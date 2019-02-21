CREATE TABLE dialogs
(
  id   SERIAL PRIMARY KEY,
  name VARCHAR(64)
);

CREATE TABLE dialog_pages
(
  id   SERIAL PRIMARY KEY,
  name VARCHAR(64),
  text VARCHAR()
);

CREATE TABLE dialog_asc
(
  id   SERIAL PRIMARY KEY,
  name VARCHAR(64)
);