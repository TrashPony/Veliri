CREATE TABLE effects_type (
  id          SERIAL PRIMARY KEY,
  name        VARCHAR(64),       /* название эффекта */
  level       INT,               /* уровень эффекта */
  type        VARCHAR(64),       /* тип эфекта: Усиливает (enhances), уменьшает (reduced), пополняет (replenishes) за ход, отнимает(takes_away) за ход, специальный (special) */
  steps_time  INT,               /* Время дейсвия эфекта в игровых ходах */
  parameter   VARCHAR(64),       /* параметр на которое влияет эфект */
  quantity    INT,               /* кол-во на которое пополяет, уменьшает и тд. */
  percentages BOOLEAN,           /* кол-во в процентах или абсолютное значение */
  forever     BOOLEAN            /* остаються параметры на всегда или нет*/
);
