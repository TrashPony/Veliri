CREATE TABLE bases (
  id              SERIAL PRIMARY KEY,
  base_name       varchar(64),
  /* id сектора то есть карты где находиться база */
  id_map          INT REFERENCES maps (id),
  /* позиция базы, на иговой карте берется обьект на этой координате и накладывается событие при нажатии */
  q               int,
  r               int,
  resp_q          int, /* resp q,r это точка выхода из базы, когда игрок нажимает выйти из базы он попадает сюда */
  resp_r          int,
  transport_count int, /* количество эвакуатор у базы */
  defender_count  int  /* количество защитников */
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