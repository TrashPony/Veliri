CREATE TABLE missions
(
  id                     SERIAL PRIMARY KEY,
  name                   VARCHAR(64) not null default '',
  type                   text        not null default '', -- delivery
  start_dialog_id        INT         not null default 0,
  not_finished_dialog_id INT         not null default 0,  -- диалог который отыгрывается если игрок не выполнил задание и вернулся к нанимающему
  reward_cr              INT         not null default 0,
  -- end_dialog_id    INT         not null default 0, определяет последний экшен
  -- end_base_id      INT         not null default 0, определяет последний экшен
  fraction               text        not null default '',
  start_base_id          INT         not null default 0,
  main_story             BOOLEAN     not null default false,
  story                  INT         not null default 1   -- порядный номер главной сюжетной линии
  -- предмет которые выдается при взятие задания, обычно его надо доставить, не является игровым предметом таблица - trash_type
  -- delivery_item_id INT         not null default 0, определяет экшон!
);

-- таблица наград в виде предметов за квест
CREATE TABLE reward_items
(
  id            SERIAL PRIMARY KEY,
  id_mission    INT REFERENCES missions (id),
  /* оружие(weapon), снаряжение(equip) или боеприпасы (ammo), корпуса (body), ресурсы (resource),
   переработака (recycle), ящики (boxes), детали (detail), чертеж (blueprints) */
  item_type     VARCHAR(64),
  slot          INT, /* какой слот занимает итем */
  item_id       INT, /* ид итема определяет конкретный итем тип + ид*/
  quantity      INT, /* количество предметов в слоте */
  hp            INT, /* сколько осталось хп у эквипа, до поломки*/
  place_user_id INT -- ид игрока который туда положить предмет(обновли последним), необходимо для публичных ящиков
);

CREATE TABLE actions
(
  id                    SERIAL PRIMARY KEY,
  id_mission            INT REFERENCES missions (id),

  type_monitor          text    not null default '',
  --complete          BOOLEAN not null default false, -- этот параметр тут не нужен, ибо он индвидуален для игрока
  description           text    not null default '',
  short_description     text    not null default '',
  base_id               INT     not null default 0,
  Q                     INT     not null default 0,
  R                     INT     not null default 0,
  map_id                INT     not null default 0,    -- ид карты куда достигать Q R )
  radius                int     not null default 0,    -- Q,R являются центром цели радиус показывает растояние от цели
  sec                   int     not null default 0,    -- количество секунд, например надо записать показания в точке QR (постоять там секунд 30)
  count                 int     not null default 0,
  -- current_count     int     not null default 0, -- этот параметр тут не нужен, ибо он индвидуален для игрока
  -- player_id         int     not null default 0, -- этот параметр тут не нужен, ибо он индвидуален для игрока
  -- bot_id            text    not null default 0, -- этот параметр тут не нужен, ибо он индвидуален для игрока
  dialog_id             INT     not null default 0,
  alternative_dialog_id INT     not null default 0,
  -- опция говорит порядок выполнения задания
  number                INT     not null default 0,
  -- опция говорит что этот экшон можно выполнять паралельно а не после предыдущего
  async                 BOOLEAN not null default false,
  owner_place           BOOLEAN not null default true, -- указывает что это должен сделать владелец задания и только он (например положить предмет куда то)
  end_text              text    not null default ''    -- текст который отсылается игроку при завершение экшона
);

CREATE TABLE need_action_items
(
  id            SERIAL PRIMARY KEY,
  id_actions    INT REFERENCES actions (id),
  /* оружие(weapon), снаряжение(equip) или боеприпасы (ammo), корпуса (body), ресурсы (resource),
   переработака (recycle), ящики (boxes), детали (detail), чертеж (blueprints) */
  item_type     VARCHAR(64),
  slot          INT, /* какой слот занимает итем */
  item_id       INT, /* ид итема определяет конкретный итем тип + ид*/
  quantity      INT, /* количество предметов в слоте */
  hp            INT, /* сколько осталось хп у эквипа, до поломки*/
  place_user_id INT -- ид игрока который туда положить предмет(обновли последним), необходимо для публичных ящиков
);