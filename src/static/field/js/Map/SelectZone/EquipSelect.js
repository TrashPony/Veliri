function MarkEquipSelect(markCode, equip) {
    var applicable = equip.applicable;

    if (applicable === "map") {

        var coordinates = {};

        for (var x in game.map.OneLayerMap) {
            if (game.map.OneLayerMap.hasOwnProperty(x)) {

                for (var y in game.map.OneLayerMap[x]) {
                    if (game.map.OneLayerMap[x].hasOwnProperty(y) && game.map.OneLayerMap[x][y].fogSprite.hide) {
                        // если скрыт туман войны то клетка видна и значит можно примернить снарягу
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

        MarkEquipZone(coordinates, equip);

    } else {

        for (var x in game.units) {
            if (game.units.hasOwnProperty(x)) {
                for (var y in game.units[x]) {
                    if (game.units[x].hasOwnProperty(y) && game.units[x][y].sprite) {

                        var unit = game.units[x][y];

                        if (applicable === "all") {
                            unit.sprite.frame = markCode;

                            unit.gameCoordinateX = x;
                            unit.gameCoordinateY = y;
                            unit.equipID = equip.id;

                            unit.sprite.events.onInputDown.add(UsedEquip, unit);
                            unit.sprite.input.priorityID = 1; // утсанавливает повышеный приоритет среди спрайтов на которых мышь
                        }

                        if (applicable === "my_units" && game.user.name === unit.owner) {
                            unit.sprite.frame = markCode;

                            unit.gameCoordinateX = x;
                            unit.gameCoordinateY = y;
                            unit.equipID = equip.id;

                            unit.sprite.events.onInputDown.add(UsedEquip, unit);
                            unit.sprite.input.priorityID = 1;
                        }

                        if (applicable === "hostile_units" && game.user.name !== unit.owner) {
                            unit.sprite.frame = markCode;

                            unit.gameCoordinateX = x;
                            unit.gameCoordinateY = y;
                            unit.equipID = equip.id;

                            unit.sprite.events.onInputDown.add(UsedEquip, unit);
                            unit.sprite.input.priorityID = 1;
                        }
                    }
                }
            }
        }
    }
}

function MarkEquipZone(coordinates, equip) {
    for (var x in coordinates) {
        if (coordinates.hasOwnProperty(x)) {
            for (var y in coordinates[x]) {
                if (coordinates[x].hasOwnProperty(y)) {

                    var cellSprite = coordinates[x][y].sprite;

                    var selectSprite = MarkZone(cellSprite, coordinates, coordinates[x][y].x, coordinates[x][y].y, 'Place', true, game.SelectTargetLineLayer, "place");

                    game.map.OneLayerMap[x][y].selectSprite = selectSprite;

                    selectSprite.inputEnabled = true;

                    selectSprite.gameCoordinateX = x;
                    selectSprite.gameCoordinateY = y;
                    selectSprite.equipID = equip.id;
                    selectSprite.equipRegion = equip.region;

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
    var xCenter = this.gameCoordinateX;
    var yCenter = this.gameCoordinateY;
    var region = this.equipRegion;

    var circleCoordinates = getRadius(xCenter, yCenter, region);

    for (var x in circleCoordinates) {
        if (circleCoordinates.hasOwnProperty(x) && game.map.OneLayerMap.hasOwnProperty(x)) {
            for (var y in circleCoordinates[x]) {
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
    for (var x in game.map.OneLayerMap) {
        if (game.map.OneLayerMap.hasOwnProperty(x)) {
            for (var y in game.map.OneLayerMap[x]) {
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