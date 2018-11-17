CREATE TABLE orders (
  id          SERIAL PRIMARY KEY,
  id_user     INT REFERENCES users (id),
  price       INT,                      /* цена за еденицу */
  count       INT,                      /* количество итемов */
  type        varchar(64),              /* buy/sell */
  min_buy_out int,                      /* минимальное количество для покупки */
  type_item   varchar(64),              /* body, ammo, weapon, equip */
  id_item     int,                      /* ид итема */
  expires     timestamp,
  place_name  varchar(64),              /* место продажи */
  place       INT REFERENCES bases (id) /* ссылка на саму базу */
);