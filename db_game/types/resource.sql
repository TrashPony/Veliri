/* ОПИСАНИЕ РЕСУРСОВ */

CREATE TABLE resource_type (
  id               SERIAL PRIMARY KEY,
  name             varchar(64),
  size             real, /* сколько весит 1 экземляр ресурса */
  /* количественное описание того что из этого может вылупица при 100% выработке */
  enriched_thorium int,
  iron             int,
  copper           int,
  titanium         int,
  silicon          int,
  plastic          int
);


CREATE TABLE recycled_resource_type (
  id   SERIAL PRIMARY KEY,
  name varchar(64),
  size real /* сколько весит 1 экземляр ресурса */
);

CREATE TABLE craft_detail (
  id               SERIAL PRIMARY KEY,
  name             varchar(64),
  size             real not null default 0, /* сколько весит 1 экземляр ресурса */

  /* количественное описание требуемых ресурсов для создания 1 штуки */

  /* примитивы */
  enriched_thorium int  not null default 0,
  iron             int  not null default 0,
  copper           int  not null default 0,
  titanium         int  not null default 0,
  silicon          int  not null default 0,
  plastic          int  not null default 0,

  /* детали тоесть фактически сылки на самих себя */
  steel            int  not null default 0,
  wire             int  not null default 0
);

/* РЕСУРСЫ НА КАРТЕ */
CREATE TABLE map_type_resource (
  id   SERIAL PRIMARY KEY,
  name varchar(64),
  type varchar(64)               /* ore - добываются экстракторами, oil водокачкой)*/
);

CREATE TABLE map_type_resource_count (/* немного ебана названа таблица, но она говорит какие руда лежат в залежах */
  id                   SERIAL PRIMARY KEY,
  id_map_resource_type INT REFERENCES map_type_resource (id),
  id_base_resource     INT REFERENCES resource_type (id),
  max_count            INT, /* количество в конкретной жиле выберает рандомом ПЕРЕД СОЗДАНИЕМ ё*/
  min_count            INT
);

-- CREATE TABLE map_resource (
--   id           SERIAL PRIMARY KEY,
--   id_map       INT REFERENCES maps (id),
--   id_type      INT REFERENCES map_type_resource (id),
--   name         varchar(64),
--   q            INT,
--   r            INT,
--   rotate       INT,
--   destroy_time timestamp /* время когда залежи самоликвидируется */
-- );
--
-- CREATE TABLE map_resource_count (/* таблица говорит в каких жилах еще остались ресурсы */
--   id               SERIAL PRIMARY KEY,
--   id_map_resource  INT REFERENCES map_resource (id),
--   id_base_resource INT REFERENCES resource_type (id),
--   resource_count   INT
-- );