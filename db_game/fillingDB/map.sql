/*  ---------------------------- Map -------------------------------------------------- */
INSERT INTO maps (name, x_size, y_size, id_type, level, specification) VALUES ('Test', 10, 10, 6, 2, 'тестовая карта на 2 человека');
INSERT INTO maps (name, x_size, y_size, id_type, level, specification) VALUES ('Test2', 10, 10, 6, 2, 'тестовая карта на 4 человека');

/*  ---------------------------- RESPAWN -------------------------------------------------- */
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 1, 1, 1, 2);
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 1, 8, 8, 2);
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (2, 1, 1, 1, 2);
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (2, 1, 8, 1, 2);
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (2, 1, 1, 8, 2);
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (2, 1, 8, 8, 2);

/*  ---------------------------- Obstacles -------------------------------------------------- */
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 8, 4, 4, 2);
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 9, 5, 4, 2);
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 10, 6, 4, 2);
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 11, 4, 5, 2);
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 12, 4, 6, 2);

/*  ---------------------------- Terrain -------------------------------------------------- */
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 3, 5, 2, 2);
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 3, 0, 4, 2);
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 4, 1, 3, 2);
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 3, 8, 2, 2);
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 4, 7, 7, 2);
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 3, 9, 6, 2);
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 3, 5, 1, 2);
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 7, 9, 0, 2);
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 7, 9, 1, 2);
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 7, 8, 0, 2);

/* Craters */
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 5, 2, 1, 2);
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 5, 7, 3, 2);
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 5, 6, 0, 2);

/* Rocks */
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 6, 2, 7, 4);
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 6, 2, 6, 3);
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 6, 1, 7, 4);
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 6, 3, 7, 4);
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 6, 2, 8, 4);
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 6, 1, 8, 4);
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 6, 3, 8, 4);
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 6, 1, 6, 3);
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 6, 3, 6, 4);

INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 6, 5, 8, 4);
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 6, 5, 7, 4);
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 6, 5, 9, 4);
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 6, 4, 9, 4);
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 6, 6, 9, 4);
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 6, 4, 8, 4);
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 6, 6, 8, 3);
INSERT INTO map_constructor (id_map, id_type, x, y, level) VALUES (1, 6, 4, 7, 4);

/* Coordinate Type */
INSERT INTO coordinate_type (id, type, texture_flore, texture_object, move, view, attack, passable_edges) VALUES (1, 'respawn' , 'terrain', '', true, true, true, true);
INSERT INTO coordinate_type (id, type, texture_flore, texture_object, move, view, attack, passable_edges) VALUES (2, '', 'desert', 'wall', false, false, false, false);
INSERT INTO coordinate_type (id, type, texture_flore, texture_object, move, view, attack, passable_edges) VALUES (3, '', 'desert', 'terrain_1', false, true, true, true);
INSERT INTO coordinate_type (id, type, texture_flore, texture_object, move, view, attack, passable_edges) VALUES (4, '', 'desert', 'terrain_2', false, true, true, true);
INSERT INTO coordinate_type (id, type, texture_flore, texture_object, move, view, attack, passable_edges) VALUES (5, '', 'desert', 'crater', true, true, true, true);
INSERT INTO coordinate_type (id, type, texture_flore, texture_object, move, view, attack, passable_edges) VALUES (6, '', 'desert', '', true, true, true, true);
INSERT INTO coordinate_type (id, type, texture_flore, texture_object, move, view, attack, passable_edges) VALUES (7, '', 'desert', 'terrain_3', false, false, false, false);
INSERT INTO coordinate_type (id, type, texture_flore, texture_object, move, view, attack, passable_edges) VALUES (8, '', 'desert', 'sand_stone_04', false, false, false, false);
INSERT INTO coordinate_type (id, type, texture_flore, texture_object, move, view, attack, passable_edges) VALUES (9, '', 'desert', 'sand_stone_05', false, false, false, false);
INSERT INTO coordinate_type (id, type, texture_flore, texture_object, move, view, attack, passable_edges) VALUES (10, '', 'desert', 'sand_stone_06', false, false, false, false);
INSERT INTO coordinate_type (id, type, texture_flore, texture_object, move, view, attack, passable_edges) VALUES (11, '', 'desert', 'sand_stone_07', false, false, false, false);
INSERT INTO coordinate_type (id, type, texture_flore, texture_object, move, view, attack, passable_edges) VALUES (12, '', 'desert', 'sand_stone_08', false, false, false, false);

/* Coordinate Effects */
INSERT INTO coordinate_type_effect (id_type, id_effect) VALUES (5, 21);