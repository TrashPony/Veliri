CREATE TABLE maps (
  id            SERIAL PRIMARY KEY,
  name          VARCHAR(64),
  x_size        INT, /* размер карты по Х */
  y_size        INT, /* размер карты по Y */
  id_type       INT REFERENCES coordinate_type (id), /* определяет основной тип тесктур если они явно не указаны в конструкторе */
  level         INT REFERENCES coordinate_type (id), /* определяет основной уровень координат на карте еще он не перепределен конструктором */
  specification VARCHAR(255)                         /* описание карты */
);

CREATE TABLE map_constructor (
  id      SERIAL PRIMARY KEY,
  id_map  INT REFERENCES maps (id), /* ид карты к которой принадлежит координата */
  id_type INT REFERENCES coordinate_type (id), /* ид типа координаты */
  x       INT,
  y       INT,
  level   INT           /* определяет уровень координаты ""примечание 1"" */
);

CREATE TABLE coordinate_type (
  id             SERIAL PRIMARY KEY,
  type           VARCHAR(64),
  texture_flore  VARCHAR(64), /* имя текстуры земли, воды, лавы и тд. */
  texture_object VARCHAR(64), /* имя текстуры обьекта (камень, дерево, стена и тд) если он есть на зоне */
  move           BOOLEAN, /* определяет можно ли ходить через эту координату */
  view           BOOLEAN, /* определяет можно ли видить через эту координату */
  attack         BOOLEAN, /* определяет можно ли атаковать через эту координату */
  passable_edges BOOLEAN      /* определяет можно ли проходить на искосок от коорднаты */
);

CREATE TABLE coordinate_type_effect (
  id_type   INT REFERENCES coordinate_type (id), /* ид координаты на которую кладем эфекты */
  id_effect INT REFERENCES effects_type (id)     /* ид эфекта которое накладывается когда юнит находиться на зоне */
);


/* ""примечание 1"".
    5 уровней высот.
    Если родитеская координата 4 а дочерняя 3 то это спуск.
    Если родитеская 3 а дочерняя 4 то это подьем.
    Если разница уровней 2 и более то проезд не возможен т.к. слишком высоко или резкий обвал
    Если разница уровней 2 снизу вверх то просмотр и стрельба невозможна
    Если разница уровней 2+ сверху в них то просмотр и стрельба с бонусом +1 за кажный уровень разницы
*/