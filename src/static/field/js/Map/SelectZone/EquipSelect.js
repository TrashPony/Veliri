function MarkEquipSelect(jsonMessage) {

    let equipSlot = JSON.parse(jsonMessage).equip_slot;
    let equip = JSON.parse(jsonMessage).equip_slot.equip;
    let targets = JSON.parse(jsonMessage).targets;
    let unit = JSON.parse(jsonMessage).unit;

    let coordinates = {};

    for (let q in targets) {
        if (targets.hasOwnProperty(q) && game.map.OneLayerMap.hasOwnProperty(q)) {
            for (let r in targets[q]) {
                if (targets[q].hasOwnProperty(r) && game.map.OneLayerMap[q].hasOwnProperty(r)) {
                    if (coordinates.hasOwnProperty(q)) {
                        coordinates[q][r] = game.map.OneLayerMap[q][r];
                    } else {
                        coordinates[q] = {};
                        coordinates[q][r] = game.map.OneLayerMap[q][r];
                    }
                }
            }
        }
    }
    MarkEquipZone(coordinates, equip, unit, equipSlot);
}

function MarkEquipZone(coordinates, equip, unit, equipSlot) {
    for (let q in coordinates) {
        if (coordinates.hasOwnProperty(q)) {
            for (let r in coordinates[q]) {
                if (coordinates[q].hasOwnProperty(r)) {

                    let cellSprite = coordinates[q][r].sprite;

                    let selectSprite = MarkZone(cellSprite, coordinates, coordinates[q][r].q, coordinates[q][r].r, 'Place', true, game.SelectTargetLineLayer, "place", game.SelectLayer);

                    game.map.OneLayerMap[q][r].selectSprite = selectSprite;

                    selectSprite.inputEnabled = true;

                    selectSprite.gameCoordinateX = coordinates[q][r].x;
                    selectSprite.gameCoordinateY = coordinates[q][r].y;
                    selectSprite.gameCoordinateZ = coordinates[q][r].z;

                    selectSprite.equipID = equip.id;
                    selectSprite.equipRegion = equip.region;

                    selectSprite.targetQ = q;
                    selectSprite.targetR = r;

                    selectSprite.unitQ = unit.q;
                    selectSprite.unitR = unit.r;
                    selectSprite.typeSlot = equipSlot.type_slot;
                    selectSprite.numberSlot = equipSlot.number_slot;

                    selectSprite.events.onInputDown.add(UsedEquip, selectSprite);
                    selectSprite.events.onInputOver.add(animateEquipCoordinate, selectSprite);
                    selectSprite.events.onInputOut.add(stopAnimateEquipCoordinate, selectSprite);

                    selectSprite.input.priorityID = 1; // утсанавливает повышеный приоритет среди спрайтов на которых мышь

                    game.map.selectSprites.push(selectSprite);
                }
            }
        }
    }
}

function animateEquipCoordinate() {
    let xCenter = this.gameCoordinateX;
    let yCenter = this.gameCoordinateY;
    let zCenter = this.gameCoordinateZ;

    let region = this.equipRegion;

    let circleCoordinates = getRadius(xCenter, yCenter, zCenter, region);

    for (let i in circleCoordinates) {
        let q = circleCoordinates[i].Q;
        let r = circleCoordinates[i].R;
        if (game.map.OneLayerMap.hasOwnProperty(q) && game.map.OneLayerMap[q].hasOwnProperty(r)) {
            if (game.map.OneLayerMap[q][r].selectSprite !== undefined) {
                game.map.OneLayerMap[q][r].selectSprite.animations.add('select');
                game.map.OneLayerMap[q][r].selectSprite.animations.play('select', 5, true);
            }
        }
    }
}

function stopAnimateEquipCoordinate() {
    for (let x in game.map.OneLayerMap) {
        if (game.map.OneLayerMap.hasOwnProperty(x)) {
            for (let y in game.map.OneLayerMap[x]) {
                if (game.map.OneLayerMap[x].hasOwnProperty(y) && game.map.OneLayerMap[x][y].selectSprite) {
                    if (game.map.OneLayerMap[x][y].selectSprite.animations.getAnimation('select') !== null) {
                        game.map.OneLayerMap[x][y].selectSprite.animations.getAnimation('select').stop(false);
                        game.map.OneLayerMap[x][y].selectSprite.animations.frame = 0;
                    }
                }
            }
        }
    }
}