-- таблица с итемами которые бесполезны для самой игры но работают как дополнения к заданиям например доставить что либо и это либо описано тут
CREATE TABLE trash_type
(
  id          SERIAL PRIMARY KEY,
  name        varchar(64),
  size        real, /* сколько весит 1 экземляр ресурса */
  /* количественное описание того что из этого может вылупица при 100% выработке */
  description text
);
