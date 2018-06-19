CREATE TABLE effects_type(
    id          serial primary key,
    name        varchar(64), /* название эффекта */
    level       int,         /* уровень эффекта */
    type        varchar(64), /* тип эфекта: Усиливает (enhances), уменьшает (reduced), пополняет (replenishes) за ход, отнимает(takes_away) за ход, специальный (special) */
    steps_time  int,         /* Время дейсвия эфекта в игровых ходах */
    parameter   varchar(64), /* параметр на которое влияет эфект */
    quantity    int,         /* кол-во на которое пополяет, уменьшает и тд. */
    percentages boolean,     /* кол-во в процентах или абсолютное значение */
    forever     boolean      /* остаються параметры на всегда или нет*/
);