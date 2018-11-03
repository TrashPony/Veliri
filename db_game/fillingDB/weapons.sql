/*  ---------------------------- weapon_type -------------------------------------------------- */
/*firearms*/
INSERT INTO weapon_type (name, min_attack_range, range_attack, accuracy, ammo_capacity, artillery, power, max_hp, size, type, standard_size, initiative)
VALUES ('tank_gun', 0, 3, 5, 5, false, 10, 100, 2, 'firearms', 1, 30);

INSERT INTO weapon_type (name, min_attack_range, range_attack, accuracy, ammo_capacity, artillery, power, max_hp, size, type, standard_size, initiative)
VALUES ('artillery', 2, 5, 5, 5, true, 20, 100, 4, 'firearms', 3, 50);

/*missile*/
INSERT INTO weapon_type (name, min_attack_range, range_attack, accuracy, ammo_capacity, artillery, power, max_hp, size, type, standard_size, initiative)
VALUES ('small_missile', 0, 4, 5, 6, false, 10, 100, 2, 'missile', 2, 40);

INSERT INTO weapon_type (name, min_attack_range, range_attack, accuracy, ammo_capacity, artillery, power, max_hp, size, type, standard_size, initiative)
VALUES ('big_missile', 3, 6, 5, 6, true, 20, 100, 4, 'missile', 3, 60);

/*laser*/
INSERT INTO weapon_type (name, min_attack_range, range_attack, accuracy, ammo_capacity, artillery, power, max_hp, size, type, standard_size, initiative)
VALUES ('small_laser', 0, 3, 5, 4, false, 20, 100, 1, 'laser', 1, 20);

INSERT INTO weapon_type (name, min_attack_range, range_attack, accuracy, ammo_capacity, artillery, power, max_hp, size, type, standard_size, initiative)
VALUES ('big_laser', 0, 4, 5, 4, false, 30, 100, 3, 'laser', 2, 40);

/* AMMO Type*/
/*firearms*/
INSERT INTO ammunition_type (name, type, type_attack, min_damage, max_damage, area_covers, size, standard_size)
VALUES ('piu-piu', 'firearms', 'kinetics', 10, 15, 0, 0.1, 2);

INSERT INTO ammunition_type (name, type, type_attack, min_damage, max_damage, area_covers, size, standard_size)
VALUES ('ba-bah', 'firearms', 'kinetics', 15, 25, 1, 0.2, 3);

/*missile*/
INSERT INTO ammunition_type (name, type, type_attack, min_damage, max_damage, area_covers, size, standard_size)
VALUES ('small_missile', 'missile', 'explosion', 12, 15, 0, 0.4, 2);

INSERT INTO ammunition_type (name, type, type_attack, min_damage, max_damage, area_covers, size, standard_size)
VALUES ('big_missile', 'missile', 'explosion', 22, 25, 1, 0.8, 3);

/*laser*/
INSERT INTO ammunition_type (name, type, type_attack, min_damage, max_damage, area_covers, size, standard_size)
VALUES ('small_lens', 'laser', 'thermo', 5, 10, 0, 0.2, 2);

INSERT INTO ammunition_type (name, type, type_attack, min_damage, max_damage, area_covers, size, standard_size)
VALUES ('big_lens', 'laser', 'thermo', 15, 20, 0, 0.4, 3);
