CREATE TABLE maps (
  id            SERIAL PRIMARY KEY,
  name          VARCHAR(64),
  q_size        INT, /* размер карты по Х */
  r_size        INT, /* размер карты по Y */
  id_type       INT REFERENCES coordinate_type (id), /* определяет основной тип тесктур если они явно не указаны в конструкторе */
  level         INT REFERENCES coordinate_type (id), /* определяет основной уровень координат на карте еще он не перепределен конструктором */
  specification VARCHAR(255)                         /* описание карты */
);

CREATE TABLE map_constructor (
  id      SERIAL PRIMARY KEY,
  id_map  INT REFERENCES maps (id), /* ид карты к которой принадлежит координата */
  id_type INT REFERENCES coordinate_type (id), /* ид типа координаты */
  q       INT,
  r       INT,
  level   INT, /* определяет уровень координаты ""примечание 1"" */

  /* impact тут показана координата которая влияет на текущую координату, параметр дается автоматически через редактор карт
   поэтому нельзя влиять на карту из бд на живую */
  impact  VARCHAR(64)
);

CREATE TABLE coordinate_type (
  id                    SERIAL PRIMARY KEY,
  type                  VARCHAR(64),
  texture_flore         VARCHAR(64), /* имя текстуры земли, воды, лавы и тд. */

  /* одновременно может быть либо статичный обьект, либо анимация */
  texture_object        VARCHAR(64), /* имя текстуры обьекта (камень, дерево, стена и тд) если он есть на зоне */
  animate_sprite_sheets VARCHAR(64), /* имя файла анимации */

  animate_loop          BOOLEAN, /* если координата анимирована говорит что анимация будет всегда по кругу иначе анимацию должно что то активировать */
  move                  BOOLEAN, /* определяет можно ли ходить через эту координату */
  view                  BOOLEAN, /* определяет можно ли видить через эту координату */
  attack                BOOLEAN, /* определяет можно ли атаковать через эту координату */

  /* impact_radius свойство которое говорит что обьект стоит не на 1 координате а занимает все координаты
  в какомто радиусе, свойство только для редактора карт, из за него теперь нельзя править карту на живую в бд т.к.
  могут быть не соотвествия */
  impact_radius         INT
);

CREATE TABLE coordinate_type_effect (
  id_type   INT REFERENCES coordinate_type (id), /* ид координаты на которую кладем эфекты */
  id_effect INT REFERENCES effects_type (id)     /* ид эфекта которое накладывается когда юнит находиться на зоне */
);