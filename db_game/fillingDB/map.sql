/*  ---------------------------- Map -------------------------------------------------- */
INSERT INTO maps (id, name, q_size, r_size, id_type, level, specification) VALUES (1, 'Test', 15, 15, 6, 2, 'тестовая карта на 2 человека');
INSERT INTO maps (id, name, q_size, r_size, id_type, level, specification) VALUES (2, 'Test2', 10, 10, 6, 2, 'тестовая карта на 4 человека');

/*  ---------------------------- RESPAWN -------------------------------------------------- */
INSERT INTO map_constructor (id_map, id_type, q, r, level, impact) VALUES (1, 1, 1, 1, 2, '');
INSERT INTO map_constructor (id_map, id_type, q, r, level, impact) VALUES (1, 1, 8, 8, 2, '');
INSERT INTO map_constructor (id_map, id_type, q, r, level, impact) VALUES (2, 1, 1, 1, 2, '');
INSERT INTO map_constructor (id_map, id_type, q, r, level, impact) VALUES (2, 1, 8, 1, 2, '');
INSERT INTO map_constructor (id_map, id_type, q, r, level, impact) VALUES (2, 1, 1, 8, 2, '');
INSERT INTO map_constructor (id_map, id_type, q, r, level, impact) VALUES (2, 1, 8, 8, 2, '');

/*  ---------------------------- Obstacles -------------------------------------------------- */
INSERT INTO map_constructor (id_map, id_type, q, r, level, impact) VALUES (1, 8, 7, 13, 2, '');
INSERT INTO map_constructor (id_map, id_type, q, r, level, impact) VALUES (1, 9, 2, 4, 2, '');
INSERT INTO map_constructor (id_map, id_type, q, r, level, impact) VALUES (1, 10, 9, 3, 2, '');
INSERT INTO map_constructor (id_map, id_type, q, r, level, impact) VALUES (1, 11, 12, 6, 2, '');
INSERT INTO map_constructor (id_map, id_type, q, r, level, impact) VALUES (1, 12, 4, 6, 2, '');

/*  ---------------------------- Terrain -------------------------------------------------- */
INSERT INTO map_constructor (id_map, id_type, q, r, level, impact) VALUES (1, 3, 5, 2, 2, '');
INSERT INTO map_constructor (id_map, id_type, q, r, level, impact) VALUES (1, 3, 0, 4, 2, '');
INSERT INTO map_constructor (id_map, id_type, q, r, level, impact) VALUES (1, 4, 1, 8, 4, '');
INSERT INTO map_constructor (id_map, id_type, q, r, level, impact) VALUES (1, 3, 8, 2, 2, '');
INSERT INTO map_constructor (id_map, id_type, q, r, level, impact) VALUES (1, 4, 7, 7, 2, '');
INSERT INTO map_constructor (id_map, id_type, q, r, level, impact) VALUES (1, 3, 9, 6, 2, '');
INSERT INTO map_constructor (id_map, id_type, q, r, level, impact) VALUES (1, 3, 5, 1, 2, '');
INSERT INTO map_constructor (id_map, id_type, q, r, level, impact) VALUES (1, 7, 9, 0, 2, '');
INSERT INTO map_constructor (id_map, id_type, q, r, level, impact) VALUES (1, 7, 9, 1, 2, '');
INSERT INTO map_constructor (id_map, id_type, q, r, level, impact) VALUES (1, 7, 8, 0, 2, '');

/* Craters */
INSERT INTO map_constructor (id_map, id_type, q, r, level, impact) VALUES (1, 5, 2, 1, 2, '');
INSERT INTO map_constructor (id_map, id_type, q, r, level, impact) VALUES (1, 5, 7, 3, 2, '');
INSERT INTO map_constructor (id_map, id_type, q, r, level, impact) VALUES (1, 5, 6, 0, 2, '');

/* Rocks */
  /* fallen */
  INSERT INTO map_constructor (id_map, id_type, q, r, level, impact) VALUES (1, 14, 3, 8, 2, '');
  INSERT INTO map_constructor (id_map, id_type, q, r, level, impact) VALUES (1, 13, 3, 6, 2, '');
INSERT INTO map_constructor (id_map, id_type, q, r, level, impact) VALUES (1, 6, 2, 7, 4, '');
INSERT INTO map_constructor (id_map, id_type, q, r, level, impact) VALUES (1, 6, 2, 6, 3, '');
INSERT INTO map_constructor (id_map, id_type, q, r, level, impact) VALUES (1, 6, 1, 7, 4, '');
INSERT INTO map_constructor (id_map, id_type, q, r, level, impact) VALUES (1, 6, 3, 7, 4, '');
INSERT INTO map_constructor (id_map, id_type, q, r, level, impact) VALUES (1, 6, 2, 8, 4, '');
INSERT INTO map_constructor (id_map, id_type, q, r, level, impact) VALUES (1, 6, 1, 6, 3, '');

INSERT INTO map_constructor (id_map, id_type, q, r, level, impact) VALUES (1, 6, 5, 8, 4, '');
INSERT INTO map_constructor (id_map, id_type, q, r, level, impact) VALUES (1, 6, 5, 7, 4, '');
INSERT INTO map_constructor (id_map, id_type, q, r, level, impact) VALUES (1, 6, 5, 9, 4, '');
INSERT INTO map_constructor (id_map, id_type, q, r, level, impact) VALUES (1, 6, 4, 9, 4, '');
INSERT INTO map_constructor (id_map, id_type, q, r, level, impact) VALUES (1, 6, 6, 9, 3, '');
INSERT INTO map_constructor (id_map, id_type, q, r, level, impact) VALUES (1, 6, 4, 8, 4, '');
INSERT INTO map_constructor (id_map, id_type, q, r, level, impact) VALUES (1, 6, 6, 8, 3, '');
INSERT INTO map_constructor (id_map, id_type, q, r, level, impact) VALUES (1, 6, 4, 7, 4, '');
INSERT INTO map_constructor (id_map, id_type, q, r, level, impact) VALUES (1, 6, 0, 7, 3, '');
INSERT INTO map_constructor (id_map, id_type, q, r, level, impact) VALUES (1, 6, 1, 8, 4, '');

/* Coordinate Type */
INSERT INTO coordinate_type (id, type, texture_flore, texture_object, animate_sprite_sheets, animate_loop, move, view, attack, impact_radius) VALUES (1, 'respawn' , 'desert', '', '', false, true, true, true, 0);
INSERT INTO coordinate_type (id, type, texture_flore, texture_object, animate_sprite_sheets, animate_loop, move, view, attack, impact_radius) VALUES (2, '', 'desert', 'wall', '', false, false, false, false, 0);
INSERT INTO coordinate_type (id, type, texture_flore, texture_object, animate_sprite_sheets, animate_loop, move, view, attack, impact_radius) VALUES (3, '', 'desert', 'terrain_1', '', false,  false, true, true, 0);
INSERT INTO coordinate_type (id, type, texture_flore, texture_object, animate_sprite_sheets, animate_loop, move, view, attack, impact_radius) VALUES (4, '', 'desert', 'terrain_2', '', false, false, true, true, 0);
INSERT INTO coordinate_type (id, type, texture_flore, texture_object, animate_sprite_sheets, animate_loop, move, view, attack, impact_radius) VALUES (5, '', 'desert', 'crater', '', false, true, true, true, 0);
INSERT INTO coordinate_type (id, type, texture_flore, texture_object, animate_sprite_sheets, animate_loop, move, view, attack, impact_radius) VALUES (6, '', 'desert', '', '', false, true, true, true, 0);
INSERT INTO coordinate_type (id, type, texture_flore, texture_object, animate_sprite_sheets, animate_loop, move, view, attack, impact_radius) VALUES (7, '', 'desert', 'terrain_3', '', false, false, false, false, 0);
INSERT INTO coordinate_type (id, type, texture_flore, texture_object, animate_sprite_sheets, animate_loop, move, view, attack, impact_radius) VALUES (8, '', 'desert', 'sand_stone_04', '', false, false, false, false, 0);
INSERT INTO coordinate_type (id, type, texture_flore, texture_object, animate_sprite_sheets, animate_loop, move, view, attack, impact_radius) VALUES (9, '', 'desert', 'sand_stone_05', '', false, false, false, false, 0);
INSERT INTO coordinate_type (id, type, texture_flore, texture_object, animate_sprite_sheets, animate_loop, move, view, attack, impact_radius) VALUES (10, '', 'desert', 'sand_stone_06', '', false, false, false, false, 0);
INSERT INTO coordinate_type (id, type, texture_flore, texture_object, animate_sprite_sheets, animate_loop, move, view, attack, impact_radius) VALUES (11, '', 'desert', 'sand_stone_07', '', false, false, false, false, 0);
INSERT INTO coordinate_type (id, type, texture_flore, texture_object, animate_sprite_sheets, animate_loop, move, view, attack, impact_radius) VALUES (12, '', 'desert', 'sand_stone_08', '', false, false, false, false, 0);
INSERT INTO coordinate_type (id, type, texture_flore, texture_object, animate_sprite_sheets, animate_loop, move, view, attack, impact_radius) VALUES (13, '', 'desert', 'fallen_01', '', false, false, true, true, 0);
INSERT INTO coordinate_type (id, type, texture_flore, texture_object, animate_sprite_sheets, animate_loop, move, view, attack, impact_radius) VALUES (14, '', 'desert', 'fallen_02', '', false, false, true, true, 0);

/* Coordinate Effects */
INSERT INTO coordinate_type_effect (id_type, id_effect) VALUES (5, 21);