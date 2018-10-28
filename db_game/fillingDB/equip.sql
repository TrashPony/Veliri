/*  ---------------------------- Equipping_type -------------------------------------------------- */
INSERT INTO equipping_type (id, name, active, specification, applicable, region, radius, type_slot, reload, power, use_power, max_hp, steps_time, size)
VALUES (1, 'repair_kit', true, 'ремонтный дроид который ремонтирует выбраного юнита на 10% каждых ход в течение 3х ходов', 'my_units', 0, 2, 2, 2, 5, 10, 100, 1, 0.5);

INSERT INTO equipping_type (id, name, active, specification, applicable, region, radius, type_slot, reload, power, use_power, max_hp, steps_time, size)
VALUES (2, 'energy_shield', true, 'Поглащает 10% всего урона выбраного юнита в течение 5ти ходов', 'my_units', 0, 2, 2, 3, 5, 10, 100, 1, 1);

INSERT INTO equipping_type (id, name, active, specification, applicable, region, radius, type_slot, reload, power, use_power, max_hp, steps_time, size)
VALUES (3, 'small_bomb', true, 'наносит урон в радиусе 1й клетки', 'map', 1, 3, 3, 4, 5, 10, 100, 3, 1.5);

INSERT INTO equipping_type (id, name, active, specification, applicable, region, radius, type_slot, reload, power, use_power, max_hp, steps_time, size)
VALUES (4, 'zone_energy_shield', true, 'Поглащает 33% всего урона выбраного юнита в течение 5ти ходов в радиусе 1й клетки', 'map', 1, 2, 3, 5, 5, 10, 100, 2, 2.5);

INSERT INTO equipping_type (id, name, active, specification, applicable, region, radius, type_slot, reload, power, use_power, max_hp, steps_time, size)
VALUES (5, 'armored', false, 'Дополнительная защита на уязвимые участки корпуса, увеличиевает броню на 10', '', 0, 0, 1, 0, 5, 0, 100, 0, 1.5);

/*  ---------------------------- Anchor_effects_to_equip -------------------------------------------------- */
INSERT INTO equip_effects (id_equip, id_effect) VALUES (1, 1);

INSERT INTO equip_effects (id_equip, id_effect) VALUES (2, 6);
INSERT INTO equip_effects (id_equip, id_effect) VALUES (2, 17);

INSERT INTO equip_effects (id_equip, id_effect) VALUES (3, 11);
INSERT INTO equip_effects (id_equip, id_effect) VALUES (3, 16);
INSERT INTO equip_effects (id_equip, id_effect) VALUES (3, 20);

INSERT INTO equip_effects (id_equip, id_effect) VALUES (4, 6);
INSERT INTO equip_effects (id_equip, id_effect) VALUES (4, 18);
INSERT INTO equip_effects (id_equip, id_effect) VALUES (4, 19);

INSERT INTO equip_effects (id_equip, id_effect) VALUES (5, 21);