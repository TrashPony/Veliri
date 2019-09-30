CREATE TABLE squads
(
  id      SERIAL PRIMARY KEY,
  name    VARCHAR(64),
  active  BOOLEAN, /* у игрока может быть в 1 момент времени только 1 отряд, это переменная говорит какой */
  id_user INT REFERENCES users (id), /* кому принадлежит */
  in_game BOOLEAN, /* отряд в бою */

  /* если отряд неактивен то он храниться на конкретной базе */
  id_base INT REFERENCES bases (id)
);

CREATE TABLE squad_thorium_slots
(
  id       SERIAL PRIMARY KEY,
  id_squad INT REFERENCES squads (id), /* ид отряда к которому принадлежит  */
  slot     INT, /* слот тория в мсе */
  thorium  INT /* сколько тория заряжено в слот */
);


CREATE TABLE squad_units
(
  id             SERIAL PRIMARY KEY,
  id_squad       INT REFERENCES squads (id), /* ид отряда к которому принадлежит юнит */

  /* из чего состоит юнит */
  id_body        INT REFERENCES body_type (id), /* ид тела юнита */
  slot           INT, /* номер слота который занимает юнит в материнской машине */

  /* Позиция */
  q              INT, /* q - колона на которой стоит юнит */
  r              INT, /* r - строка на которой стоит юнит */
  rotate         INT,
  id_map         INT,

  /* игрок в бою */
  on_map         BOOLEAN,
  -- ид боя в котором забыли юнита, удалится как бой закончится, актуально ток для ботов, мс нельзя забыть в бою
  id_game        INT,

  /* Игровая статистика */
  target         VARCHAR(64),
  defend         BOOLEAN, /* означат что пользователь защищается юнитов в фазе атаки */
  mother_ship    BOOLEAN, /* является ли этот юнит мазршипом */
  move           BOOLEAN, /* говорит что сейчас ходит именно этот юнит */

  /* Характиристики */
  hp             INT,
  power          INT,
  action_point   INT, /* очки передвижения юнита */

  /* покраска юнитов */
  body_color_1   text not null default '0x15ccff',
  body_color_2   text not null default '0x000000',
  weapon_color_1 text not null default '0x15ccff',
  weapon_color_2 text not null default '0x000000',

  /* путь к файлу готовой покраске, пока не реализовано */
  body_texture   text not null default '',
  weapon_texture text not null default ''
);

CREATE TABLE squad_units_equipping
(/* таблица снаряжения которое нацеплино на юнита */
  id               SERIAL PRIMARY KEY,
  id_squad         INT REFERENCES squads (id),
  type_slot        INT, /* тип слота */
  type             VARCHAR(64), /* оружие(weapon), снаряжение(equip) или боеприпасы (ammo)*/
  id_squad_unit    INT REFERENCES squad_units, /* ид юнита к которому прикреплено снаряжение */
  id_equipping     INT, /* ид снаряжения или оружия */
  slot_in_body     INT, /* слот который занимает снаряжения */
  quantity         INT, /* количество предметов в слоте */
  used             BOOLEAN, /* если true значит уже использовано и ждет перезарядки */
  steps_for_reload INT, /* сколько шагов осталось до перезарядки */
  hp               INT, /* сколько осталось хп у эквипа, до поломки*/
  target           VARCHAR(64) /* цель снаряжения, говорит куда применять его на фазе атаки */
);

/* Главным инвентарем отряда конечно является МП, однако у других юнитов тож есть инвентари например что руда копать, или в ящиках лазить */
CREATE TABLE squad_units_inventory
(/* инвентарь отряда не боевой параметр */
  id            SERIAL PRIMARY KEY,
  id_unit       INT not null default 0, /* какому юниту в отряде принадлежит инвентарь */

  /* оружие(weapon), снаряжение(equip) или боеприпасы (ammo), корпуса (body), ресурсы (resource),
   переработака (recycle), ящики (boxes), детали (detail), чертеж (blueprints), (trash) */
  item_type     text not null default '',
  slot          INT not null default 0, /* какой слот занимает итем */
  item_id       INT not null default 0, /* ид итема определяет конкретный итем тип + ид*/
  quantity      INT not null default 0, /* количество предметов в слоте */
  hp            INT not null default 0, /* сколько осталось хп у эквипа, до поломки*/
  place_user_id INT not null default 0 -- ид игрока который туда положить предмет(обновли последним), необходимо для публичных ящиков
);
