CREATE TABLE bases (
  id        SERIAL PRIMARY KEY,
  base_name varchar(64)
  /* пока не доработана логика, сейчас тупо ради ангара и базового инвентаря существует */
);

CREATE TABLE base_users ( /* игроки которые сейчас сидят на конкретной базу */
  id        SERIAL PRIMARY KEY,
  base_id   INT REFERENCES bases (id),
  user_id   INT REFERENCES users (id)
);

CREATE TABLE base_storage ( /* инвнтерь конкретной базы конктретного игрока */
  id        SERIAL PRIMARY KEY,
  base_id   INT REFERENCES bases (id),
  user_id   INT REFERENCES users (id),
  slot      INT,            /* какой слот занимает итем */
  item_type VARCHAR(64),    /* оружие(weapon), снаряжение(equip) или боеприпасы (ammo), корпуса (body) */
  item_id   INT,            /* ид итема определяет конкретный итем тип + ид*/
  quantity  INT,            /* количество предметов в слоте */
  hp        INT             /* сколько осталось хп у эквипа, до поломки*/
);