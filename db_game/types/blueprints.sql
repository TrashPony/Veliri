CREATE TABLE blueprints
(
  id               SERIAL PRIMARY KEY,
  name             VARCHAR(64),
  /* оружие(weapon), снаряжение(equip) или боеприпасы (ammo), корпуса (body), ресурсы (resource),
     переработака (recycle), ящики (boxes), детали (detail) */
  item_type        VARCHAR(64),
  /* ид итема определяет конкретный итем (тип + ид), который получит пользователь после крафта*/
  item_id          INT,
  /* имя файла иконки */
  icon             VARCHAR(64) not null default '',
  /* количесво секунд необходимое для создания итема */
  craft_time       INT         not null default 0,
  /* бесконечное количество прогонов */
  original         BOOLEAN     not null default false,
  /* количество прогон которое осталось */
  copies           INT         not null default 0,
  /* количество предметов которое выйдет из чертежа */
  count            INT         not null default 1,

  /* количественное описание требуемых ресурсов для создания 1 штуки */
  ---- примитивы
  enriched_thorium int         not null default 0,
  iron             int         not null default 0,
  copper           int         not null default 0,
  titanium         int         not null default 0,
  silicon          int         not null default 0,
  plastic          int         not null default 0,

  ---- детали
  steel            int         not null default 0,
  wire             int         not null default 0, -- провода
  electronics      int         not null default 0,
  wires            int         not null default 0, -- проволока
  gear             int         not null default 0,
  titanium_plate   int         not null default 0,
  batteries        int         not null default 0,
  armor_items      int         not null default 0
);

CREATE TABLE created_blueprint
(
  id                     SERIAL PRIMARY KEY,
  /* определяет итем который будет на выходе */
  id_blueprint           INT REFERENCES blueprints (id),
  /* база где происходит крафт и куда упасть на склад */
  id_base                INT REFERENCES bases (id),
  /* какому игроку */
  id_user                INT REFERENCES users (id),
  /*время начала*/
  start_time             timestamp,
  /* время окончания */
  finish_time            timestamp,
  /* процент налога миниралов */
  mineral_tax_percentage INT,
  /* процент налога времени */
  time_tax_percentage    INT
);