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
                        MarkZone(cellSprite, targetCoordinates, q, r, 'Target', false, game.SelectTargetLineLayer, null, game.SelectLayer);
                    }

                    if (event === "GetTargets" || event === "GetEquipMapTargets") {
                        let selectSprite = MarkZone(cellSprite, targetCoordinates, q, r, 'Target', true, game.SelectTargetLineLayer, "target", game.SelectLayer);

                        selectSprite.TargetQ = targetCoordinates[q][r].q;
                        selectSprite.TargetR = targetCoordinates[q][r].r;

                        selectSprite.unitQ = unitQ;
                        selectSprite.unitR = unitR;
                        selectSprite.UnitID = unitID;

                        selectSprite.inputEnabled = true;

                        selectSprite.events.onInputDown.add(func, selectSprite);
                        selectSprite.events.onInputOver.add(animateTargetCoordinate, selectSprite);
                        selectSprite.events.onInputOut.add(stopAnimateCoordinate, selectSprite);

                        selectSprite.input.priorityID = 1; // утсанавливает повышеный приоритет среди спрайтов на которых мышь

                        game.map.selectSprites.push(selectSprite);
                    }
                }
            }
        }
    }
}

function SelectTargetUnit(jsonMessage) {
    let units = JSON.parse(jsonMessage).units;
    let unit = JSON.parse(jsonMessage).unit;
    let equipSlot = JSON.parse(jsonMessage).equip_slot;

    for (let i in units){
        if (units.hasOwnProperty(i)) {

            let func =()=> {
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
