function CreateTerrain(coordinate) {

    var floorSprite = CreateSpriteTerrain(coordinate);
    var fogSprite = game.fogOfWar.create(coordinate.x * game.tileSize, coordinate.y * game.tileSize, 'FogOfWar');

    floorSprite.inputEnabled = true; // включаем ивенты на спрайт
    floorSprite.events.onInputOut.add(TipOff, floorSprite);
    floorSprite.events.onInputDown.add(RemoveSelect);
    floorSprite.z = 0;

    coordinate.sprite = floorSprite;
    coordinate.fogSprite = fogSprite;
}

function CreateSpriteTerrain(coordinate) {
    // todo нечитабельный говнокод
    var leftLevel = coordinateLevel(coordinate.x - 1, coordinate.y, coordinate.level);
    var leftTopLevel = coordinateLevel(coordinate.x - 1, coordinate.y - 1, coordinate.level);
    var leftBotLevel = coordinateLevel(coordinate.x - 1, coordinate.y + 1, coordinate.level);

    var rightLevel = coordinateLevel(coordinate.x + 1, coordinate.y, coordinate.level);
    var rightTopLevel = coordinateLevel(coordinate.x + 1, coordinate.y - 1, coordinate.level);
    var rightBotLevel = coordinateLevel(coordinate.x + 1, coordinate.y + 1, coordinate.level);

    var topLevel = coordinateLevel(coordinate.x, coordinate.y - 1, coordinate.level);
    var bottomLevel = coordinateLevel(coordinate.x, coordinate.y + 1, coordinate.level);

    if (leftLevel === rightLevel
        && leftLevel === topLevel
        && leftLevel === bottomLevel
        && leftLevel === rightBotLevel
        && leftLevel === leftTopLevel
        && leftLevel === leftBotLevel
        && leftLevel === rightTopLevel) {

        // если спрайты вокруг 1го уровня то спрайт прямой
        return game.floorLayer.create(coordinate.x * game.tileSize, coordinate.y * game.tileSize, coordinate.texture_flore + "1");

    } else {
        if (leftLevel === coordinate.level - 1 && rightLevel === coordinate.level + 1) { // подьем с лева на право
            return game.floorLayer.create(coordinate.x * game.tileSize, coordinate.y * game.tileSize, coordinate.texture_flore + "9");
        }
        if (leftLevel === coordinate.level + 1 && rightLevel === coordinate.level - 1) { // подьем с право на луво
            return game.floorLayer.create(coordinate.x * game.tileSize, coordinate.y * game.tileSize, coordinate.texture_flore + "5");
        }
        if (topLevel === coordinate.level + 1 && bottomLevel === coordinate.level - 1) { // подьем с верха вниз
            return game.floorLayer.create(coordinate.x * game.tileSize, coordinate.y * game.tileSize, coordinate.texture_flore + "3");
        }
        if (topLevel === coordinate.level - 1 && bottomLevel === coordinate.level + 1) { // подьем с низа вверх
            return game.floorLayer.create(coordinate.x * game.tileSize, coordinate.y * game.tileSize, coordinate.texture_flore + "7");
        }

        if (leftTopLevel === coordinate.level - 1 && rightBotLevel === coordinate.level + 1) { // лево верх -> право низ
            return game.floorLayer.create(coordinate.x * game.tileSize, coordinate.y * game.tileSize, coordinate.texture_flore + "2");
        }

        if (leftTopLevel === coordinate.level + 1 && rightBotLevel === coordinate.level - 1) { // право низ -> лево верх
            return game.floorLayer.create(coordinate.x * game.tileSize, coordinate.y * game.tileSize, coordinate.texture_flore + "6");
        }

        if (leftBotLevel === coordinate.level + 1 && rightTopLevel === coordinate.level - 1) { // право верх -> лево низ
            return game.floorLayer.create(coordinate.x * game.tileSize, coordinate.y * game.tileSize, coordinate.texture_flore + "4");
        }

        if (leftBotLevel === coordinate.level - 1 && rightTopLevel === coordinate.level + 1) { //  лево низ -> право верх
            return game.floorLayer.create(coordinate.x * game.tileSize, coordinate.y * game.tileSize, coordinate.texture_flore + "8");
        }
    }


    return game.floorLayer.create(coordinate.x * game.tileSize, coordinate.y * game.tileSize, coordinate.texture_flore + "1");
}

function coordinateLevel(x, y, defaultLevel) {

    if (game.map.OneLayerMap.hasOwnProperty(x)) {
        if (game.map.OneLayerMap[x].hasOwnProperty(y)) {
            return game.map.OneLayerMap[x][y].level
        } else {
            return defaultLevel
        }
    } else {
        return defaultLevel
    }
}
