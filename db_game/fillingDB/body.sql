/*  ---------------------------- body_type -------------------------------------------------- */
INSERT INTO body_type (id, name, mother_ship, speed, initiative, max_hp, armor, evasion_critical, vulnerability_to_kinetics,
                       vulnerability_to_thermo, vulnerability_to_explosion, range_view, accuracy, max_power, recovery_power, wall_hack,
                       capacity_size, standard_size, standard_size_small, standard_size_medium, standard_size_big)
VALUES (3, 'light_tank', false, 4, 15, 15, 2, 5, 5, 5, 5, 3, 4, 20, 10, false, 3, 1, true, false, false);

INSERT INTO body_type (id, name, mother_ship, speed, initiative, max_hp, armor, evasion_critical, vulnerability_to_kinetics,
                       vulnerability_to_thermo, vulnerability_to_explosion, range_view, accuracy, max_power, recovery_power, wall_hack,
                       capacity_size, standard_size, standard_size_small, standard_size_medium, standard_size_big)
VALUES (4, 'medium_tank', false, 3, 10, 25, 3, 5, 10, 10, 10, 3, 3, 35, 15, false, 5, 2, true, true, false);

INSERT INTO body_type (id, name, mother_ship, speed, initiative, max_hp, armor, evasion_critical, vulnerability_to_kinetics,
                       vulnerability_to_thermo, vulnerability_to_explosion, range_view, accuracy, max_power, recovery_power, wall_hack,
                       capacity_size, standard_size, standard_size_small, standard_size_medium, standard_size_big)
VALUES (1, 'heavy_tank', false, 2, 5, 40, 5, 5, 15, 15, 15, 2, 3, 50, 20, false, 10, 3, true, true, true);

INSERT INTO body_type (name, mother_ship, speed, initiative, max_hp, armor, evasion_critical, vulnerability_to_kinetics,
                       vulnerability_to_thermo, vulnerability_to_explosion, range_view, accuracy, max_power, recovery_power, wall_hack,
                       capacity_size, standard_size, standard_size_small, standard_size_medium, standard_size_big)
VALUES ('MasherShip_1', true, 4, 25, 100, 5, 15, 15, 15, 15, 3, 5, 100, 20, false, 50, 3, true, true, true);

INSERT INTO body_type (name, mother_ship, speed, initiative, max_hp, armor, evasion_critical, vulnerability_to_kinetics,
                       vulnerability_to_thermo, vulnerability_to_explosion, range_view, accuracy, max_power, recovery_power, wall_hack,
                       capacity_size, standard_size, standard_size_small, standard_size_medium, standard_size_big)
VALUES ('MasherShip_2', true, 4, 25, 100, 5, 15, 15, 15, 15, 3, 5, 100, 20, false, 60, 3, true, true, true);

INSERT INTO body_type (name, mother_ship, speed, initiative, max_hp, armor, evasion_critical, vulnerability_to_kinetics,
                       vulnerability_to_thermo, vulnerability_to_explosion, range_view, accuracy, max_power, recovery_power, wall_hack,
                       capacity_size, standard_size, standard_size_small, standard_size_medium, standard_size_big)
VALUES ('MasherShip_3', true, 2, 25, 100, 5, 15, 15, 15, 15, 3, 5, 100, 20, false, 100, 3, true, true, true);

/* BODY TANK*/
/* HEAVY*/
/* I */
INSERT INTO body_slots (id_body, type_slot, number_slot, weapon, weapon_type)
VALUES (1, 1, 1, false, '');
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

/*medium*/
INSERT INTO body_slots (id_body, type_slot, number_slot, weapon, weapon_type)
VALUES (4, 1, 1, false, '');
/* II */
INSERT INTO body_slots (id_body, type_slot, number_slot, weapon, weapon_type)
VALUES (4, 2, 1, false, '');
/* III */
INSERT INTO body_slots (id_body, type_slot, number_slot, weapon, weapon_type)
VALUES (4, 3, 1, true, '');

/*light*/
INSERT INTO body_slots (id_body, type_slot, number_slot, weapon, weapon_type)
VALUES (3, 1, 1, false, '');
/* II */
/* III */
INSERT INTO body_slots (id_body, type_slot, number_slot, weapon, weapon_type)
VALUES (3, 3, 1, true, '');

/* BODY MOTHER*/
/* I */
INSERT INTO body_slots (id_body, type_slot, number_slot, weapon, weapon_type)
VALUES (5, 1, 1, false, '');
INSERT INTO body_slots (id_body, type_slot, number_slot, weapon, weapon_type)
VALUES (5, 1, 2, false, '');
INSERT INTO body_slots (id_body, type_slot, number_slot, weapon, weapon_type)
VALUES (5, 1, 3, false, '');
/* II */
INSERT INTO body_slots (id_body, type_slot, number_slot, weapon, weapon_type)
VALUES (5, 2, 1, false, '');
INSERT INTO body_slots (id_body, type_slot, number_slot, weapon, weapon_type)
VALUES (5, 2, 2, false, '');
/* III */
INSERT INTO body_slots (id_body, type_slot, number_slot, weapon, weapon_type)
VALUES (5, 3, 1, false, '');
INSERT INTO body_slots (id_body, type_slot, number_slot, weapon, weapon_type)
VALUES (5, 3, 2, false, '');
INSERT INTO body_slots (id_body, type_slot, number_slot, weapon, weapon_type)
VALUES (5, 3, 3, false, '');

/* IV */
INSERT INTO body_slots (id_body, type_slot, number_slot, weapon, weapon_type, standard_size)
VALUES (5, 4, 1, false, '', 1);
INSERT INTO body_slots (id_body, type_slot, number_slot, weapon, weapon_type, standard_size)
VALUES (5, 4, 2, false, '', 1);
INSERT INTO body_slots (id_body, type_slot, number_slot, weapon, weapon_type, standard_size)
VALUES (5, 4, 3, false, '', 2);
INSERT INTO body_slots (id_body, type_slot, number_slot, weapon, weapon_type, standard_size)
VALUES (5, 4, 4, false, '', 2);
INSERT INTO body_slots (id_body, type_slot, number_slot, weapon, weapon_type, standard_size)
VALUES (5, 4, 5, false, '', 3);
INSERT INTO body_slots (id_body, type_slot, number_slot, weapon, weapon_type, standard_size)
VALUES (5, 4, 6, false, '', 3);

/* ячейки с торием */
INSERT INTO body_thorium_slots (id_body, number_slot, max_thorium)
VALUES (5, 1, 120);
INSERT INTO body_thorium_slots (id_body, number_slot, max_thorium)
VALUES (5, 2, 100);