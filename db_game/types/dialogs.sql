CREATE TABLE dialogs
(
  id          SERIAL PRIMARY KEY,
  name        text not null default '',
  /*
  доступ диалога, base - можно вызвать на базе, object - привязан к какому либо обьекту и только рядом с ним можно
  его вызвать, world - можно вызвать везде и всегда
  */
  access_type text not null default ''
);

CREATE TABLE dialog_pages
(
  id        SERIAL PRIMARY KEY,
  id_dialog INT REFERENCES dialogs (id),
  /* номер страницы */
  number    INT  not null default 0,
  name      text not null default '',
  /* содержание страницы, можно писать HTML */
  text      text not null default '',
  /* имя файла персонажа который ведет диалог */
  picture   text not null default ''
);

CREATE TABLE dialog_asc
(
  id          SERIAL PRIMARY KEY,
  /* id_page отвечает за то где показывать этот ответ */
  id_page     INT REFERENCES dialog_pages (id),
  /* номер страницы на которую ведет ответ, если 0 то закрывает диалог */
  to_page     INT  not null default 1,
  name        text not null default '',
  text        text not null default '',
  /* задает тип функции которая отрботает если нажать отмет */
  type_action text not null default ''
);