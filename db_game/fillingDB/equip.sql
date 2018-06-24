/*  ---------------------------- Equipping_type -------------------------------------------------- */
INSERT INTO equipping_type (id, type, active, specification, applicable, region, radius, type_slot, reload, power)
VALUES (1, 'repair_kit', true, 'ремонтный дроид который ремонтирует выбраного юнита на 10% каждых ход в течение 3х ходов', 'my_units', 0, 0, 2, 2, 10);

INSERT INTO equipping_type (id, type, active, specification, applicable, region, radius, type_slot, reload, power)
VALUES (2, 'energy_shield', true, 'Поглащает 10% всего урона выбраного юнита в течение 5ти ходов', 'my_units', 0, 0, 2, 3, 20);

INSERT INTO equipping_type (id, type, active, specification, applicable, region, radius, type_slot, reload, power)
VALUES (3, 'small_bomb', true, 'наносит урон в радиусе 1й клетки', 'map', 1, 3, 3, 4, 30);

INSERT INTO equipping_type (id, type, active, specification, applicable, region, radius, type_slot, reload, power)
VALUES (4, 'zone_energy_shield', true, 'Поглащает 33% всего урона выбраного юнита в течение 5ти ходов в радиусе 1й клетки', 'map', 1, 2, 3, 5, 40);

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