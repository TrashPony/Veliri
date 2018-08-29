/*  ---------------------------- Effects_type -------------------------------------------------- */
INSERT INTO effects_type (id, name, level, type, parameter, quantity, percentages, forever)  VALUES (1, 'repair', 1,'replenishes', 'hp', 10, true, true);
INSERT INTO effects_type (id, name, level, type, parameter, quantity, percentages, forever)  VALUES (2, 'repair', 2,'replenishes', 'hp', 15, true, true);
INSERT INTO effects_type (id, name, level, type, parameter, quantity, percentages, forever)  VALUES (3, 'repair', 3,'replenishes', 'hp', 20, true, true);
INSERT INTO effects_type (id, name, level, type, parameter, quantity, percentages, forever)  VALUES (4, 'repair', 4,'replenishes', 'hp', 25, true, true);
INSERT INTO effects_type (id, name, level, type, parameter, quantity, percentages, forever)  VALUES (5, 'repair', 5,'replenishes', 'hp', 30, true, true);

INSERT INTO effects_type (id, name, level, type, parameter, quantity, percentages, forever)  VALUES (6, 'shield', 1,'enhances', 'armor', 10, false, false);
INSERT INTO effects_type (id, name, level, type, parameter, quantity, percentages, forever)  VALUES (7, 'shield', 2,'enhances', 'armor', 20, false, false);
INSERT INTO effects_type (id, name, level, type, parameter, quantity, percentages, forever)  VALUES (8, 'shield', 3,'enhances', 'armor', 30, false, false);
INSERT INTO effects_type (id, name, level, type, parameter, quantity, percentages, forever)  VALUES (9, 'shield', 4,'enhances', 'armor', 40, false, false);
INSERT INTO effects_type (id, name, level, type, parameter, quantity, percentages, forever)  VALUES (10, 'shield', 5,'enhances', 'armor', 50, false, false);

INSERT INTO effects_type (id, name, level, type, parameter, quantity, percentages, forever)  VALUES (11, 'damage', 1,'takes_away', 'hp', 10, false, true);
INSERT INTO effects_type (id, name, level, type, parameter, quantity, percentages, forever)  VALUES (12, 'damage', 2,'takes_away', 'hp', 20, false, true);
INSERT INTO effects_type (id, name, level, type, parameter, quantity, percentages, forever)  VALUES (13, 'damage', 3,'takes_away', 'hp', 30, false, true);
INSERT INTO effects_type (id, name, level, type, parameter, quantity, percentages, forever)  VALUES (14, 'damage', 4,'takes_away', 'hp', 40, false, true);
INSERT INTO effects_type (id, name, level, type, parameter, quantity, percentages, forever)  VALUES (15, 'damage', 5,'takes_away', 'hp', 50, false, true);

INSERT INTO effects_type (id, name, level, type, parameter, quantity, percentages, forever)  VALUES (16, 'animate_explosion', 1, 'animate', '', 0, false, false);
INSERT INTO effects_type (id, name, level, type, parameter, quantity, percentages, forever)  VALUES (17, 'animate_energy_shield', 1, 'unit_always_animate', '', 0, false, false);
INSERT INTO effects_type (id, name, level, type, parameter, quantity, percentages, forever)  VALUES (18, 'animate_energy_shield', 1, 'zone_always_animate', '', 0, false, false);

INSERT INTO effects_type (id, name, level, type, parameter, quantity, percentages, forever)  VALUES (19, 'energy_shield_zone_anchor', 1, 'anchor', '', 0, false, false);
INSERT INTO effects_type (id, name, level, type, parameter, quantity, percentages, forever)  VALUES (20, 'small_bomb_zone_anchor', 1, 'anchor', '', 0, false, false);

INSERT INTO effects_type (id, name, level, type, parameter, quantity, percentages, forever)  VALUES (21, 'defend', 1, 'enhances', 'armor', 10, false, false);
