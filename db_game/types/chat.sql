CREATE TABLE chats
(
  id             SERIAL PRIMARY KEY,
  name           VARCHAR(64),
  avatar         text,                           -- todo надо переделать на bytea забирать и обновлять методами decode(string text, format text) и encode(data bytea, format text) но мне лень
  greetings      text,                           -- сообщение которое выводится когда чат загружается самым перым сообщением.
  public         BOOLEAN,
  password       varchar(64),
  fraction       varchar(64),                    --Replics Explores Reverses
  private        BOOLEAN not null default false,
  private_key    text    not null default '',
  user_create    boolean not null default false, -- есть чат кастоный то true, для чатов фракционных или "помощь новичкам" и подобные false
  user_id_create int     not null default 0      -- ид пользователя который создал, да предыдущий параметр как бы уже и не нужен...
);

CREATE TABLE users_in_chat
(
  id      SERIAL PRIMARY KEY,
  id_user INT REFERENCES users (id),
  id_chat INT REFERENCES chats (id)
);
