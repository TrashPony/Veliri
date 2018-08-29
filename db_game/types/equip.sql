CREATE TABLE equipping_type (
  id            SERIAL PRIMARY KEY,
  name          VARCHAR(64),
  active        BOOLEAN,             /* активное или пасивное оборудование */
  specification VARCHAR(255),        /* описание снаряжения */
  applicable    VARCHAR(64),         /* прменимо к своим "my_units", вражеским "hostile_units", всем "all",  карте "map", на себя "myself" */
  region        INT,                 /* регион которе покрывает эфектом на карте относительно применения */
  radius        INT,                 /* радиус на которое работает снаряжение */
  type_slot     INT,                 /* тип слота куда встаривается оборудование I (1) , II (2), III (3), IV (4), V (5) */
  reload        INT,                 /* кол-во ходов на перезарядку */
  power         INT,                 /* кол-во потребляемой энергии */
  use_power     INT,                 /* кол-во расходуемой энергии при активации */
  max_hp        INT,                 /* кол-во хп до половки оборудования */
  steps_time    INT                  /* Время дейсвия эфектов в игровых ходах */
);

CREATE TABLE equip_effects (
  id_equip  INT REFERENCES equipping_type (id), /* ид снаряжения */
  id_effect INT REFERENCES effects_type (id)    /* ид эфекта которое накладывается снаряжением */
);