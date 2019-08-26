CREATE TABLE box_in_map
(/* таблица описывает все ящики которые лежат на карте */
  id           SERIAL PRIMARY KEY,
  password     int, /* если ящик общий то пароль - 0, если запаролен то 1-4 числа */
  destroy_time timestamp, /* время когда ящик самоликвидируется */
  id_map       INT REFERENCES maps (id),
  id_box_type  INT REFERENCES box_type (id),
  q            int,
  r            int,
  rotate       int,
  current_hp   int not null default 1
);

CREATE TABLE box_type
(
  id            SERIAL PRIMARY KEY,
  name          varchar(64),
  type          varchar(64), /* описывает какую текстуру загружать */
  capacity_size REAL, /* вместимость в кубо-метрах */
  fold_size     REAL, /* размер если ящик нести в инвентаре */
  protect       BOOLEAN, /* тру-на ящик можно поставить пароль */
  protect_lvl   int, /* 1-5 число описывающие сложность замка */
  underground   BOOLEAN, /* если ящик под землей то его нельзя задавить */
  hp            int not null default 100
);

CREATE TABLE box_storage
(
  id            SERIAL PRIMARY KEY,
  id_box        INT REFERENCES box_in_map (id),
  slot          INT, /* какой слот занимает итем */
  /* оружие(weapon), снаряжение(equip) или боеприпасы (ammo), корпуса (body), ресурсы (resource),
  переработака (recycle), ящики (boxes), детали (detail), чертеж (blueprints) */
  item_type     VARCHAR(64),
  item_id       INT, /* ид итема определяет конкретный итем тип + ид*/
  quantity      INT, /* количество предметов в слоте */
  hp            INT, /* сколько осталось хп у эквипа, до поломки*/
  place_user_id INT -- ид игрока который туда положить предмет(обновли последним), необходимо для публичных ящиков
);