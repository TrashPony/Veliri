function SelectTargetCoordinateCreate(jsonMessage, func) {
    let targetCoordinates = JSON.parse(jsonMessage).targets;

    let event = JSON.parse(jsonMessage).event;

    let unitX = JSON.parse(jsonMessage).unit.x;
    let unitY = JSON.parse(jsonMessage).unit.y;
    let unitID = JSON.parse(jsonMessage).unit.id;

    for (let x in targetCoordinates) {
        if (targetCoordinates.hasOwnProperty(x)) {
            for (let y in targetCoordinates[x]) {
                if (targetCoordinates[x].hasOwnProperty(y)) {
                    let cellSprite = game.map.OneLayerMap[targetCoordinates[x][y].x][targetCoordinates[x][y].y].sprite;

                    if (event === "GetFirstTargets") {
                        MarkZone(cellSprite, targetCoordinates, x, y, 'Target', false, game.SelectTargetLineLayer, null);
                    }

                    if (event === "GetTargets" || event === "GetEquipMapTargets") {
                        let selectSprite = MarkZone(cellSprite, targetCoordinates, x, y, 'Target', true, game.SelectTargetLineLayer, "target");

                        selectSprite.TargetX = targetCoordinates[x][y].x;
                        selectSprite.TargetY = targetCoordinates[x][y].y;

                        selectSprite.unitX = unitX;
                        selectSprite.unitY = unitY;
                        selectSprite.UnitID = unitID;

                        selectSprite.inputEnabled = true;
                        // Todo если в клетку куда показывает юзер есть юнит надо показывать сколько примерно отниметься хп
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
                    x: Number(unit.x),
                    y: Number(unit.y),
                    to_x: Number(GetGameUnitID(units[i].id).x),
                    to_y: Number(GetGameUnitID(units[i].id).y),
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
            x: Number(selectSprite.unitX),
            y: Number(selectSprite.unitY),
            to_x: Number(selectSprite.TargetX),
            to_y: Number(selectSprite.TargetY)
        }));
        RemoveSelect();
    }
}
