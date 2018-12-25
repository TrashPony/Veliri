CREATE TABLE resource_type (
  id               SERIAL PRIMARY KEY,
  name             varchar(64),
  size             real, /* сколько весит 1 экземляр ресурса */
  /* todo количественное описание того что из этого может вылупица при 100% выработке*/
  enriched_thorium int
);

CREATE TABLE recycled_resource_type (
  id   SERIAL PRIMARY KEY,
  name varchar(64),
  size real /* сколько весит 1 экземляр ресурса */
);