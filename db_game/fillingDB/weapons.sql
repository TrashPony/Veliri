/*  ---------------------------- weapon_type -------------------------------------------------- */
INSERT INTO weapon_type (name, min_attack_range, range_attack, accuracy, ammo_capacity, artillery, power, max_hp, size, type, standard_size)
VALUES ('tank_gun', 0, 3, 5, 5, false, 10, 100, 4, 'firearms', 2);

INSERT INTO weapon_type (name, min_attack_range, range_attack, accuracy, ammo_capacity, artillery, power, max_hp, size, type, standard_size)
VALUES ('artillery', 2, 5, 5, 5, true, 20, 100, 8, 'firearms', 3);

/* AMMO Type*/
INSERT INTO ammunition_type (name, type, type_attack, min_damage, max_damage, area_covers, size, standard_size)
VALUES ('piu-piu', 'firearms', 'kinetics', 10, 15, 0, 0.1, 2);

INSERT INTO ammunition_type (name, type, type_attack, min_damage, max_damage, area_covers, size, standard_size)
VALUES ('ba-bah', 'firearms', 'kinetics', 15, 25, 1, 0.2, 3);