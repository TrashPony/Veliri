/*  ---------------------------- body_type -------------------------------------------------- */
INSERT INTO body_type (id, name, mother_ship, speed, initiative, max_hp, armor, evasion_critical, vulnerability_to_kinetics,
                       vulnerability_to_thermo, vulnerability_to_em, vulnerability_to_explosion, range_view, accuracy, power, wall_hack)
VALUES (1, 'tank', false, 3, 5, 40, 5, 5, 15, 15, 15, 15, 3, 3, 30, false);

INSERT INTO body_type (id, name, mother_ship, speed, initiative, max_hp, armor, evasion_critical, vulnerability_to_kinetics,
                       vulnerability_to_thermo, vulnerability_to_em, vulnerability_to_explosion, range_view, accuracy, power, wall_hack)
VALUES (2, 'Mother', true, 3, 25, 100, 5, 15, 15, 15, 15, 15, 3, 5, 100, false);

/* BODY TANK*/
/* I */
INSERT INTO body_slots (id_body, type_slot, number_slot, weapon, weapon_type)
VALUES (1, 1, 1, true, '');
/* II */
INSERT INTO body_slots (id_body, type_slot, number_slot, weapon, weapon_type)
VALUES (1, 2, 1, false, '');
INSERT INTO body_slots (id_body, type_slot, number_slot, weapon, weapon_type)
VALUES (1, 2, 2, false, '');
/* III */
INSERT INTO body_slots (id_body, type_slot, number_slot, weapon, weapon_type)
VALUES (1, 3, 1, true, '');
INSERT INTO body_slots (id_body, type_slot, number_slot, weapon, weapon_type)
VALUES (1, 3, 2, false, '');

/* BODY MOTHER*/
/* I */
INSERT INTO body_slots (id_body, type_slot, number_slot, weapon, weapon_type)
VALUES (1, 1, 1, true, '');
INSERT INTO body_slots (id_body, type_slot, number_slot, weapon, weapon_type)
VALUES (1, 1, 2, true, '');
/* II */
INSERT INTO body_slots (id_body, type_slot, number_slot, weapon, weapon_type)
VALUES (1, 2, 1, false, '');
INSERT INTO body_slots (id_body, type_slot, number_slot, weapon, weapon_type)
VALUES (1, 2, 2, false, '');
INSERT INTO body_slots (id_body, type_slot, number_slot, weapon, weapon_type)
VALUES (1, 2, 3, false, '');
/* III */
INSERT INTO body_slots (id_body, type_slot, number_slot, weapon, weapon_type)
VALUES (1, 3, 1, true, '');
INSERT INTO body_slots (id_body, type_slot, number_slot, weapon, weapon_type)
VALUES (1, 3, 2, false, '');
INSERT INTO body_slots (id_body, type_slot, number_slot, weapon, weapon_type)
VALUES (1, 3, 3, true, '');