CREATE TABLE maps
(
  id              SERIAL PRIMARY KEY,
  name            VARCHAR(64),
  x_size          INT, /* размер карты по Х */
  y_size          INT, /* размер карты по Y */
  id_type         INT REFERENCES coordinate_type (id), /* определяет основной тип тесктур если они явно не указаны в конструкторе */
  level           INT, /* определяет основной уровень координат на карте еще он не перепределен конструктором */
  specification   VARCHAR(255), /* описание карты */

  /* если параметр global true то это карта является глобальной тоесть неимзенной картой мира */
  /* если false то это локальная карта */
  global          BOOLEAN,

  /* параметр in_game, может быть только у лоалькой карты, если он true - значит карта принадлежит какой то активной игре */
  /* иначе это шаблон локальной карты например для быстрых боев. */
  in_game         BOOLEAN,

  -- координаты отображения в меню глобальной карты
  x_global        int,
  y_global        int,

  --Replics Explores Reverses
  fraction        text,
  -- переменная отвечает можно захватить этот сектор или нет
  possible_battle boolean
);

CREATE TABLE map_constructor
(
  id                 SERIAL PRIMARY KEY,
  id_map             INT REFERENCES maps (id), /* ид карты к которой принадлежит координата */
  id_type            INT, /* ид типа координаты */
  texture_over_flore VARCHAR(64), /* название текстуры земли */
  /* говорит в какой последовательности отрисовывать текстуры, ид енподходит т.к. координата уже могла быть в бд перед нанесения текстуры */
  texture_priority   INT,
  /* тоже самое но для обьектов */
  object_priority    INT,
  x                  INT,
  y                  INT,
  rotate             INT, /* говорит на сколько повернуться спрайту обьекта в координате если он есть конечно */
  animate_speed      INT, /* если координата анимация говорит с какой скоростью ее вопспроизводить, кадров в секунду */
  x_shadow_offset    INT, /* смещение тени по Х от центра координаты */
  y_shadow_offset    INT, /* смещение тени по Y от центра координаты */
  shadow_intensity   INT, /* сила тени от 0 до 1, (val / 100) */
  scale              INT, /* определяет какой размер должен быть тексуры обьекта, анимации на карте 100 - 100%, 10 - 10%, 200 - 200% от оригинала */
  shadow             BOOLEAN, /* определяет нужна ли обьекту тень */

  /* если тру то с течением времени или по эвенту игрока эвакуируют с этой клетки без его желания */
  transport          BOOLEAN,

  /* если строка не пуста значит эта клетка прослушивается, например вход в базу (base) или переход в другой сектор (sector),
  и когда игрок на ней происходит событие */
  handler            VARCHAR(64),

  /* соотвественно место куда попадает игрок после ивента */
  to_positions       json, -- [ {"x": 1, "y": 1, "resp_rotate":90}, {"x": 2, "y": 2, "resp_rotate":90} ], координаты куда ведет телепорт
  to_base_id         INT,  -- не внешний ключь потому что в го нет налов, а сылки на базу может и не быть)
  to_map_id          INT REFERENCES maps (id)
);

CREATE TABLE coordinate_type
(
  id                    SERIAL PRIMARY KEY,
  type                  VARCHAR(64),
  texture_flore         VARCHAR(64), /* имя текстуры земли, воды, лавы и тд. -- УСТАРЕЛО -- */

  /* одновременно может быть либо статичный обьект, либо анимация */
  texture_object        VARCHAR(64), /* имя текстуры обьекта (камень, дерево, стена и тд) если он есть на зоне */
  animate_sprite_sheets VARCHAR(64), /* имя файла анимации */

  animate_loop          BOOLEAN, /* если координата анимирована говорит что анимация будет всегда по кругу иначе анимацию должно что то активировать */
  /* параметр чисто для отображения, говорит перекроет юнит своим телом этот обьект или нет если надетет на него*/
  unit_overlap          BOOLEAN,

  object_name           text    not null default '',
  object_description    text    not null default '',
  object_inventory      BOOLEAN not null default false, -- для этой координаты по ид создается ящик в таблице ящиков на карте
  object_hp             int     not null default -1,    -- -2 - бесмертный и даже нет хп, -1 - бесмертный с хп, 0 - мертвый

  /* геодата привязаная к обьекту,
   работает как и глобальная дата но х у это смещение от центра обьекта и пропадает если обьекта больше нет.
   x, y смещается если обьект повернут, имеет размер отличный от 100%
   radius изменяется если обьект имеет размер отличный от 100%
   [ {"x": 1, "y": 1, "radius":90}, {"x": 2, "y": 2, "radius":90} ]
   */
  geo_data              json    not null default '{}'
);

CREATE TABLE coordinate_type_effect
(
  id_type   INT REFERENCES coordinate_type (id), /* ид координаты на которую кладем эфекты */
  id_effect INT REFERENCES effects_type (id) /* ид эфекта которое накладывается когда юнит находиться на зоне */
);

CREATE TABLE global_geo_data
(
  id     SERIAL PRIMARY KEY,
  id_map INT REFERENCES maps (id), /* где находится непроходимая точка */
  x      INT,
  y      INT,
  radius INT /* размер непроходимой точки */
);

CREATE TABLE map_danger_anomalies /* опасные аномалии на карте (не обязаельно опасные С: ) */
(
  id     SERIAL PRIMARY KEY,
  id_map INT REFERENCES maps (id), /* где находится непроходимая точка */
  x      INT,
  y      INT,
  radius INT,
  type   text, -- тип аномалии (unknown, mortality)
  power  int   -- сила аномалии, типо mortality 999 отнимает 999 хп в тик проверки аномалии
);

-- таблица лазерных людей на карте, например связи от сдания к здания, энерго забор и тд
CREATE TABLE map_beams
(
  id      SERIAL PRIMARY KEY,
  id_map  INT REFERENCES maps (id),
  x_start INT  not null default 0,
  y_start INT  not null default 0,
  x_end   INT  not null default 0,
  y_end   INT  not null default 0,
  color   text not null default '0x000000'
);

-- таблица источников туманов, дыма, бурь и тд на карте
CREATE TABLE map_emitters
(
  id              SERIAL PRIMARY KEY,
  id_map          INT REFERENCES maps (id),
  x               INT     not null default 0,
  y               INT     not null default 0,
  -- минимальный размер частицы в источнике min_scale/100
  min_scale       INT     not null default 0,
  -- макс размер частицы в источнике max_scale/100
  max_scale       INT     not null default 0,
  -- минимальный скорость частицы
  min_speed       INT     not null default 0,
  -- макс скорость частицы
  max_speed       INT     not null default 0,
  -- время жизни частицы
  ttl             INT     not null default 0,
  -- ширина источника
  width           INT     not null default 0,
  -- высота источника
  height          INT     not null default 0,
  -- цветовой фильтр который применится к частицам
  color           text    not null default '',
  -- скорость пораждения частиц ms
  frequency       INT     not null default 0,
  -- min_alpha/100
  min_alpha       INT     not null default 0,
  -- max_alpha/100
  max_alpha       INT     not null default 0,
  -- включить анимацтю для частицы
  animate         BOOLEAN not null default false,
  -- кадры в секунду
  animate_speed   INT     not null default 0,
  -- имя спрайта
  name_particle   text    not null default '',
  -- время за которое альфа пройдет от мин к макс
  alpha_loop_time INT     not null default 0,
  -- если тру то после того как альфа наберет макс, пойдет обратно плавно назад
  yoyo            BOOLEAN not null default false
);