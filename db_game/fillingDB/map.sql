/*  ---------------------------- Map -------------------------------------------------- */
INSERT INTO maps (name, x_size, y_size, id_type, level, specification) VALUES ('Test', 10, 10, 6, 3, 'тестовая карта на 2 человека');
INSERT INTO maps (name, x_size, y_size, id_type, level, specification) VALUES ('Test2', 10, 10, 6, 3, 'тестовая карта на 4 человека');

/*  ---------------------------- RESPAWN -------------------------------------------------- */
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 1, 1, 1, 3);
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 1, 8, 8, 3);
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (2, 1, 1, 1, 3);
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (2, 1, 8, 1, 3);
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (2, 1, 1, 8, 3);
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (2, 1, 8, 8, 3);

/*  ---------------------------- Obstacles -------------------------------------------------- */
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 2, 4, 4, 3);
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 2, 5, 4, 3);
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 2, 6, 4, 3);
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 2, 4, 5, 3);
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 2, 4, 6, 3);

/*  ---------------------------- Terrain -------------------------------------------------- */
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 3, 5, 2, 3);
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 3, 0, 4, 3);
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 4, 1, 3, 3);
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 3, 8, 2, 3);
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 4, 7, 7, 3);
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 3, 9, 6, 3);
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 3, 5, 1, 3);


/* Craters */
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 5, 2, 1, 3);
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 5, 7, 3, 3);
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 5, 6, 0, 3);

/* Coordinate Type */
INSERT INTO coordinate_type (id, type, texture_flore, texture_object, move, view, attack, passable_edges) VALUES (1, 'respawn' , 'terrain', '', true, true, true, true);
INSERT INTO coordinate_type (id, type, texture_flore, texture_object, move, view, attack, passable_edges) VALUES (2, '', 'desert', 'wall', false, false, false, false);
INSERT INTO coordinate_type (id, type, texture_flore, texture_object, move, view, attack, passable_edges) VALUES (3, '', 'desert', 'terrain_1', false, true, true, true);
INSERT INTO coordinate_type (id, type, texture_flore, texture_object, move, view, attack, passable_edges) VALUES (4, '', 'desert', 'terrain_2', false, true, true, true);
INSERT INTO coordinate_type (id, type, texture_flore, texture_object, move, view, attack, passable_edges) VALUES (5, '', 'desert', 'crater', true, true, true, true);
INSERT INTO coordinate_type (id, type, texture_flore, texture_object, move, view, attack, passable_edges) VALUES (6, '' , 'terrain', '', true, true, true, true);

/* Coordinate Effects */
INSERT INTO coordinate_type_effect (id_type, id_effect) VALUES (5, 21);