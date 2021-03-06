CREATE TABLE bases
(
  id                           SERIAL PRIMARY KEY,
  base_name                    varchar(64),
  /* id сектора то есть карты где находиться база */
  id_map                       INT REFERENCES maps (id),
  /* позиция базы, на иговой карте берется обьект на этой координате и накладывается событие при нажатии */
  x                            int,
  y                            int,
  transport_count              int, /* количество эвакуатор у базы */
  defender_count               int, /* количество защитников */
  gravity_radius               int, /* радиус стабильной гравитации вокруг баз */
  --Replics Explores Reverses, нация за кторую играет игрок
  fraction                     varchar(64),
  capital                      boolean not null default false, -- сталица фракции
  -- количество ресурсов ниже которого будет снижатся налоги на переработку (на каждый ресурс индивидуально)
  boundary_amount_of_resources int not null default 1000,

  -- количество ресурсов(сумарное для всех) ниже которого помвышаются налоги на услуги базы (1% недостатка = 1% налога)
  sum_work_resources           int not null default 10000
);

-- таблица в себе содержит точки респаунов на
/* resp x,y это точка выхода из базы, когда игрок нажимает выйти из базы он попадает сюда */
CREATE TABLE bases_respawns
(
  id      SERIAL PRIMARY KEY,
  base_id INT REFERENCES bases (id),
  x       int,
  y       int,
  rotate  int not null default 0 -- направление выхода с базы, угол от 0 до 360 который принимает корпус при выходе
);

CREATE TABLE base_users
( /* игроки которые сейчас сидят на конкретной базу */
  id      SERIAL PRIMARY KEY,
  base_id INT REFERENCES bases (id),
  user_id INT REFERENCES users (id)
);

CREATE TABLE base_storage
( /* инвнтерь конкретной базы конктретного игрока */
  id            SERIAL PRIMARY KEY,
  base_id       INT REFERENCES bases (id),
  user_id       INT REFERENCES users (id),
  /* оружие(weapon), снаряжение(equip) или боеприпасы (ammo), корпуса (body), ресурсы (resource),
   переработака (recycle), ящики (boxes), детали (detail), чертеж (blueprints) */
  item_type     VARCHAR(64),
  slot          INT, /* какой слот занимает итем */
  item_id       INT, /* ид итема определяет конкретный итем тип + ид*/
  quantity      INT, /* количество предметов в слоте */
  hp            INT, /* сколько осталось хп у эквипа, до поломки*/
  place_user_id INT -- ид игрока который туда положить предмет(обновли последним), необходимо для публичных ящиков
);