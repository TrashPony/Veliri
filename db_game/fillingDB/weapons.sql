/*  ---------------------------- weapon_type -------------------------------------------------- */
INSERT INTO weapon_type (name, min_attack_range, range_attack, accuracy, ammo_capacity, artillery, power, max_hp)
VALUES ('tank_gun', 0, 3, 5, 5, false, 10, 100);

INSERT INTO weapon_type (name, min_attack_range, range_attack, accuracy, ammo_capacity, artillery, power, max_hp)
VALUES ('artillery', 2, 5, 5, 5, true, 20, 100);

/* AMMO Type*/
INSERT INTO ammunition_type (name, type, type_attack, min_damage, max_damage, area_covers)
VALUES ('piu-piu', 'em', 'em', 10, 15, 0);

INSERT INTO ammunition_type (name, type, type_attack, min_damage, max_damage, area_covers)
VALUES ('ba-bah', 'explosion', 'explosion', 15, 25, 1);