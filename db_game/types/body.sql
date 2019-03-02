CREATE TABLE body_type (
  id                         SERIAL PRIMARY KEY,
  name                       VARCHAR(64),
  mother_ship                BOOLEAN,
  speed                      INT,
  initiative                 INT,      /*  */
  max_hp                     INT,
  armor                      INT,      /* блокируемый урон в абсолюте */
  evasion_critical           INT,
  vulnerability_to_kinetics  INT,
  vulnerability_to_thermo    INT,
  vulnerability_to_explosion INT,
  range_view                 INT,
  accuracy                   INT,
  max_power                  INT,      /* макс колво энергии */
  recovery_power             INT,      /* востановление энергии за ход */
  recovery_hp                INT,      /* востановление hp например за счет дронов */
  wall_hack                  BOOLEAN,
  capacity_size              REAL,     /* вместимость корпуса к кубо-метрах а так же его вес */
  standard_size              INT,      /* small - 1, medium - 2, big - 3, размер корпуса (если корпус мс то неучитывается)*/
  standard_size_small        BOOLEAN,  /* оружие которое может использовать корпус small, medium, big */
  standard_size_medium       BOOLEAN,  /* оружие которое может использовать корпус small, medium, big */
  standard_size_big          BOOLEAN,  /* оружие которое может использовать корпус small, medium, big */

  /*методанные для детектора колизий*/
  body_front_radius          INT,   /* радиус передней сферы мешины */
  body_left_front_angle      INT,   /* градус отклонения для детектора колизий */
  body_right_front_angle     INT,   /* градус отклонения для детектора колизий */

  body_back_radius          INT,   /* радиус задней сферы мешины */
  body_left_back_angle      INT,   /* градус отклонения для детектора колизий */
  body_right_back_angle     INT,   /* градус отклонения для детектора колизий */

  body_side_radius          INT    /* радиус боковой сферы мешины */
);

CREATE TABLE body_thorium_slots (
  id_body       INT REFERENCES body_type (id), /* ид корпуса которому принадлежит слот*/
  number_slot   INT,                           /* номер слота в корпусе */
  max_thorium   INT                            /* сколько макс, тория помещается в ячейке */
);

CREATE TABLE body_slots (
  id_body       INT REFERENCES body_type (id), /* ид корпуса которому принадлежит слот*/
  type_slot     INT,                           /* тип слота куда встаривается оборудование I (1) , II (2), III (3), IV (4), V (5) */
  number_slot   INT,                           /* номер слота в корпусе */
  weapon        BOOLEAN,                       /* труе если слот приналдлежит оружию */
  weapon_type   varchar(64),                   /* тип оружия например "artillery" */
  standard_size INT                            /* определяет тип вмещаемого юнита если это ангар */
);