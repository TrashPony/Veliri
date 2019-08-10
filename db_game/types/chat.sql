CREATE TABLE chats
(
  id       SERIAL PRIMARY KEY,
  name     VARCHAR(64),
  public   BOOLEAN,
  password varchar(64),
  fraction varchar(64), --Replics Explores Reverses
  private  BOOLEAN not null default false,
  private_key text not null default ''
);

CREATE TABLE users_in_chat
(
  id      SERIAL PRIMARY KEY,
  id_user INT REFERENCES users (id),
  id_chat INT REFERENCES chats (id)
);
