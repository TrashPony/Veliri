function MarkEquipSelect(jsonMessage) {

    let equipSlot = JSON.parse(jsonMessage).equip_slot;
    let equip = JSON.parse(jsonMessage).equip_slot.equip;
    let targets = JSON.parse(jsonMessage).targets;
    let unit = JSON.parse(jsonMessage).unit;

    let coordinates = {};

    for (let x in targets) {
        if (targets.hasOwnProperty(x) && game.map.OneLayerMap.hasOwnProperty(x)) {
            for (let y in targets[x]) {
                if (targets[x].hasOwnProperty(y) && game.map.OneLayerMap[x].hasOwnProperty(y)) {
                    if (coordinates.hasOwnProperty(x)) {
                        coordinates[x][y] = game.map.OneLayerMap[x][y];
                    } else {
                        coordinates[x] = {};
                        coordinates[x][y] = game.map.OneLayerMap[x][y];
                    }
                }
            }
        }
    }
    MarkEquipZone(coordinates, equip, unit, equipSlot);
}

function MarkEquipZone(coordinates, equip, unit, equipSlot) {
    for (let x in coordinates) {
        if (coordinates.hasOwnProperty(x)) {
            for (let y in coordinates[x]) {
                if (coordinates[x].hasOwnProperty(y)) {

                    let cellSprite = coordinates[x][y].sprite;

                    let selectSprite = MarkZone(cellSprite, coordinates, coordinates[x][y].x, coordinates[x][y].y, 'Place', true, game.SelectTargetLineLayer, "place");

                    game.map.OneLayerMap[x][y].selectSprite = selectSprite;

                    selectSprite.inputEnabled = true;

                    selectSprite.gameCoordinateX = x;
                    selectSprite.gameCoordinateY = y;
                    selectSprite.equipID = equip.id;
                    selectSprite.equipRegion = equip.region;

                    selectSprite.unitX = unit.x;
                    selectSprite.unitY = unit.y;
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
    let region = this.equipRegion;

    let circleCoordinates = getRadius(xCenter, yCenter, region);

    for (let x in circleCoordinates) {
        if (circleCoordinates.hasOwnProperty(x) && game.map.OneLayerMap.hasOwnProperty(x)) {
            for (let y in circleCoordinates[x]) {
                if (circleCoordinates[x].hasOwnProperty(y) && game.map.OneLayerMap[x].hasOwnProperty(y)) {
                    if (game.map.OneLayerMap[x][y].selectSprite !== undefined) {
                        game.map.OneLayerMap[x][y].selectSprite.animations.add('select');
                        game.map.OneLayerMap[x][y].selectSprite.animations.play('select', 5, true);
                    }
                }
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