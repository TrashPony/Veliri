function SelectTargetCoordinateCreate(jsonMessage, func) {
    let targetCoordinates = JSON.parse(jsonMessage).targets;

    let event = JSON.parse(jsonMessage).event;

    let unitQ = JSON.parse(jsonMessage).unit.q;
    let unitR = JSON.parse(jsonMessage).unit.r;
    let unitID = JSON.parse(jsonMessage).unit.id;

    for (let q in targetCoordinates) {
        if (targetCoordinates.hasOwnProperty(q)) {
            for (let r in targetCoordinates[q]) {
                if (targetCoordinates[q].hasOwnProperty(r)) {
                    let cellSprite = game.map.OneLayerMap[targetCoordinates[q][r].q][targetCoordinates[q][r].r].sprite;

                    if (event === "GetFirstTargets") {
                        MarkZone(cellSprite, targetCoordinates, q, r, 'Target', false, game.SelectTargetLineLayer, null, game.SelectLayer, true);
                    }

                    if (event === "GetTargets" || event === "GetEquipMapTargets") {
                        let selectSprite = MarkZone(cellSprite, targetCoordinates, q, r, 'Target', true, game.SelectTargetLineLayer, "target", game.SelectLayer, false);

                        selectSprite.TargetQ = targetCoordinates[q][r].q;
                        selectSprite.TargetR = targetCoordinates[q][r].r;

                        selectSprite.unitQ = unitQ;
                        selectSprite.unitR = unitR;
                        selectSprite.UnitID = unitID;

                        selectSprite.inputEnabled = true;

                        selectSprite.events.onInputDown.add(func, selectSprite);
                        selectSprite.events.onInputOver.add(getTargetZone, selectSprite);
                        selectSprite.events.onInputOut.add(removeTargetZone, selectSprite);

                        selectSprite.input.priorityID = 0; // утсанавливает повышеный приоритет среди спрайтов на которых мышь

                        game.map.OneLayerMap[q][r].targetSelectSprite = selectSprite;

                        game.map.selectSprites.push(selectSprite);
                    }
                }
            }
        }
    }
}

function removeTargetZone(coordinate) {

    let unit = GetGameUnitXY(coordinate.unitQ, coordinate.unitR);

    let areaCovers;
    for (let weaponSlot in unit.body.weapons) { // оружие может быть только 1 под диз доке, масив это обман
        if (unit.body.weapons.hasOwnProperty(weaponSlot) && unit.body.weapons[weaponSlot].ammo) {
            areaCovers = unit.body.weapons[weaponSlot].ammo.area_covers;
        }
    }
    let targetCoordinate = game.map.OneLayerMap[coordinate.TargetQ][coordinate.TargetR];

    let damageZone = getRadius(targetCoordinate.x, targetCoordinate.y, targetCoordinate.z, areaCovers);

    for (let i in damageZone) {
        let q = damageZone[i].Q;
        let r = damageZone[i].R;
        if (game.map.OneLayerMap.hasOwnProperty(q) && game.map.OneLayerMap[q].hasOwnProperty(r)) {
            let animateCoordinate = game.map.OneLayerMap[q][r];

            if (animateCoordinate.targetSelectSprite) {
                stopAnimateCoordinate(animateCoordinate.targetSelectSprite);
            }

            if (animateCoordinate.fastTargetSprite) {
                stopAnimateCoordinate(animateCoordinate.fastTargetSprite);
                animateCoordinate.fastTargetSprite.destroy();
                animateCoordinate.fastTargetSprite = null;
            }

            if (animateCoordinate.damageText) {
                animateCoordinate.damageText.destroy();
            }

            if (GetGameUnitXY(q, r)) {
                HideUnitStatus(GetGameUnitXY(q, r))
            }
        }
    }
}

function getTargetZone(coordinate) {
    RemoveTargetLine();

    let unit = GetGameUnitXY(coordinate.unitQ, coordinate.unitR);

    let areaCovers;
    for (let weaponSlot in unit.body.weapons) { // оружие может быть только 1 под диз доке, масив это обман
        if (unit.body.weapons.hasOwnProperty(weaponSlot) && unit.body.weapons[weaponSlot].ammo) {
            areaCovers = unit.body.weapons[weaponSlot].ammo.area_covers;
        }
    }
    let targetCoordinate = game.map.OneLayerMap[coordinate.TargetQ][coordinate.TargetR];

    let damageZone = getRadius(targetCoordinate.x, targetCoordinate.y, targetCoordinate.z, areaCovers);

    for (let i in damageZone) {
        let q = damageZone[i].Q;
        let r = damageZone[i].R;
        if (game.map.OneLayerMap.hasOwnProperty(q) && game.map.OneLayerMap[q].hasOwnProperty(r)) {
            let animateCoordinate = game.map.OneLayerMap[q][r];

            if (animateCoordinate.targetSelectSprite) {
                animateTargetCoordinate(animateCoordinate.targetSelectSprite);
                damageText(q, r, animateCoordinate, coordinate, unit)
            } else {
                let selectSprite = MarkZone(animateCoordinate.sprite, damageZone, q, r, 'Target', true, game.SelectTargetLineLayer, "target", game.SelectLayer, false);
                game.map.OneLayerMap[q][r].fastTargetSprite = selectSprite;
                animateTargetCoordinate(selectSprite);
                damageText(q, r, animateCoordinate, coordinate, unit)
            }
        }
    }
}

function damageText(q, r, animateCoordinate, coordinate, unit) {
    let targetUnit = GetGameUnitXY(q, r);

    if (targetUnit) {

        VisibleUnitStatus(targetUnit);

        if (animateCoordinate.damageText) {
            animateCoordinate.damageText.destroy();
        }

        let style = {font: "20px Finger Paint", fill: "#C00"};
        let damageText;

        if (coordinate.TargetQ === q && coordinate.TargetR === r) {
            damageText = game.add.text(targetUnit.sprite.x + 20, targetUnit.sprite.y - 50, getMinMaxDamage(unit, targetUnit, false), style);
        } else {
            damageText = game.add.text(targetUnit.sprite.x + 20, targetUnit.sprite.y - 50, getMinMaxDamage(unit, targetUnit, true), style);
        }

        damageText.setShadow(1, -1, 'rgba(0,0,0,0.5)', 0);
        animateCoordinate.damageText = damageText;
    }
}

function SelectTargetUnit(jsonMessage) {
    let units = JSON.parse(jsonMessage).units;
    let unit = JSON.parse(jsonMessage).unit;
    let equipSlot = JSON.parse(jsonMessage).equip_slot;

    for (let i in units) {
        if (units.hasOwnProperty(i)) {

            let func = () => {
                field.send(JSON.stringify({
                    event: "UseUnitEquip",
                    unit_id: Number(unit.id),
                    q: Number(unit.q),
                    r: Number(unit.r),
                    to_q: Number(GetGameUnitID(units[i].id).q),
                    to_r: Number(GetGameUnitID(units[i].id).r),
                    equip_id: Number(equipSlot.equip.id),
                    equip_type: Number(equipSlot.type_slot),
                    number_slot: Number(equipSlot.number_slot)
                }));
                RemoveSelect();
            };

            MarkUnitSelect(GetGameUnitID(units[i].id), 2, func)
        }
    }
}

function SelectWeaponTarget(selectSprite) {

    removeTargetZone(selectSprite);

    if (game.input.activePointer.leftButton.isDown) {
        field.send(JSON.stringify({
            event: "SetWeaponTarget",
            unit_id: Number(selectSprite.UnitID),
            q: Number(selectSprite.unitQ),
            r: Number(selectSprite.unitR),
            to_q: Number(selectSprite.TargetQ),
            to_r: Number(selectSprite.TargetR)
        }));
        RemoveSelect();
    }
}
