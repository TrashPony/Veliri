CREATE TABLE orders (
  id          SERIAL PRIMARY KEY,
  id_user     INT REFERENCES users (id),
  price       INT,
  count       INT,
  type        varchar(64), /* buy/sell */
  min_buy_out int,
  type_item   varchar(64), /* body, ammo, weapon, equip */
  id_item     int,
  expires     timestamp
);