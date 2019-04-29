CREATE TABLE missions
(
  id                 SERIAL PRIMARY KEY,
  name               VARCHAR(64) not null default '',
  start_dialog_id    INT         not null default 0,
  reward_cr          INT         not null default 0,
  end_dialog_id      INT         not null default 0,
  end_base_id_dialog INT         not null default 0,
  fraction           text        not null default '',
  start_base_id      INT         not null default 0
);

alter table missions
  add column fraction text not null default '';
alter table missions
  add column start_base_id INT not null default 0;

-- таблица наград в виде предметов за квест
CREATE TABLE reward_items
(
  id         SERIAL PRIMARY KEY,
  id_mission INT REFERENCES missions (id),
  /* оружие(weapon), снаряжение(equip) или боеприпасы (ammo), корпуса (body), ресурсы (resource),
   переработака (recycle), ящики (boxes), детали (detail), чертеж (blueprints) */
  item_type  VARCHAR(64),
  slot       INT, /* какой слот занимает итем */
  item_id    INT, /* ид итема определяет конкретный итем тип + ид*/
  quantity   INT, /* количество предметов в слоте */
  hp         INT /* сколько осталось хп у эквипа, до поломки*/
);

CREATE TABLE actions
(
  id                SERIAL PRIMARY KEY,
  id_mission        INT REFERENCES missions (id),

  type_monitor      text    not null,
  complete          BOOLEAN not null default false,
  description       text    not null,
  short_description text    not null,
  base_id           INT     not null default 0,
  Q                 INT     not null default 0,
  R                 INT     not null default 0,
  count             int     not null default 0,
  current_count     int     not null default 0,
  player_id         int     not null default 0,
  bot_id            text    not null default 0,
  dialog_id         INT     not null default 0,

  -- опция говорит порядок выполнения задания
  number            INT     not null default 0,
  -- опция говорит что этот экшон можно выполнять паралельно а не после предыдущего
  async             BOOLEAN not null default false
);

CREATE TABLE need_action_items
(
  id         SERIAL PRIMARY KEY,
  id_actions INT REFERENCES actions (id),
  /* оружие(weapon), снаряжение(equip) или боеприпасы (ammo), корпуса (body), ресурсы (resource),
   переработака (recycle), ящики (boxes), детали (detail), чертеж (blueprints) */
  item_type  VARCHAR(64),
  slot       INT, /* какой слот занимает итем */
  item_id    INT, /* ид итема определяет конкретный итем тип + ид*/
  quantity   INT, /* количество предметов в слоте */
  hp         INT /* сколько осталось хп у эквипа, до поломки*/
);
