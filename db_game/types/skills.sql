CREATE TABLE skills
(
  id               SERIAL PRIMARY KEY,
  name             VARCHAR(64),
  specification    VARCHAR(255),
  experience_point INT,  -- необходимых опыт для изучения 1 лвл скила, дальше необходимо (lvl + 1) * experience_point */
  type             text, -- scientific, attack, production
  icon             text  -- картинка в base64
);

CREATE TABLE user_skills
(
  id       SERIAL PRIMARY KEY,
  lvl      INT,
  id_skill INT REFERENCES skills (id),
  id_user  INT REFERENCES users (id)
);