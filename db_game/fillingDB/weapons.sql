/*  ---------------------------- weapon_type -------------------------------------------------- */
INSERT INTO weapon_type (name, min_attack_range, range_attack, accuracy, artillery, power)
VALUES ('tank_gun', 0, 3, 5, false, 10);

INSERT INTO weapon_type (name, min_attack_range, range_attack, accuracy, artillery, power)
VALUES ('artillery', 2, 5, 5, true, 20);

/* AMMO Type*/
INSERT INTO ammunition_type (name, type, type_attack, damage, area_covers)
VALUES ('piu-piu', 'em', 'em', 15, 0);

INSERT INTO ammunition_type (name, type, type_attack, damage, area_covers)
VALUES ('ba-bah', 'explosion', 'explosion', 25, 1);