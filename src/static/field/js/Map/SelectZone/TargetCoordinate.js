function SelectTargetCoordinateCreate(jsonMessage) {

    var targetCoordinates = JSON.parse(jsonMessage).targets;

    for (var x in targetCoordinates) {
        if (targetCoordinates.hasOwnProperty(x)) {
            for (var y in targetCoordinates[x]) {
                if (targetCoordinates[x].hasOwnProperty(y)) {
                    var cellSprite = game.map.OneLayerMap[targetCoordinates[x][y].x][targetCoordinates[x][y].y].sprite;
                    MarkTarget(cellSprite, targetCoordinates, x, y, 'Target');
                }
            }
        }
    }
}

function MarkTarget(cellSprite, placeCoordinate, x, y, selectClass) {

    var left = false;
    var right = false;
    var top = false;
    var bot = false;
    var line;

    if (placeCoordinate.hasOwnProperty(Number(x) + 1)) {
        if (placeCoordinate[Number(x) + 1].hasOwnProperty(y)) {
            right = true;
        }
    }

    if (placeCoordinate.hasOwnProperty(Number(x) - 1)) {
        if (placeCoordinate[Number(x) - 1].hasOwnProperty(y)) {
            left = true;
        }
    }

    if (placeCoordinate[x].hasOwnProperty(Number(y) - 1)) {
        top = true;
    }

    if (placeCoordinate[x].hasOwnProperty(Number(y) + 1)) {
        bot = true;
    }

    if (right && left && !top && bot) {
        line = game.SelectTargetLineLayer.create(cellSprite.x, cellSprite.y, 'select' + selectClass + '_1');
    }

    if (right && !left && top && bot) {
        line = game.SelectTargetLineLayer.create(cellSprite.x, cellSprite.y, 'select' + selectClass + '_1');
        line.anchor.setTo(1, 0);
        line.angle = -90;
    }

    if (!right && left && top && bot) {
        line = game.SelectTargetLineLayer.create(cellSprite.x, cellSprite.y, 'select' + selectClass + '_1');
        line.anchor.setTo(0, 1);
        line.angle = 90;
    }

    if (right && left && top && !bot) {
        line = game.SelectTargetLineLayer.create(cellSprite.x, cellSprite.y, 'select' + selectClass + '_1');
        line.anchor.setTo(0, 1);
        line.scale.y *= -1;
    }

    if (!right && left && top && !bot) {
        line = game.SelectTargetLineLayer.create(cellSprite.x, cellSprite.y, 'select' + selectClass + '_2');
        line.anchor.setTo(0, 1);
        line.scale.y *= -1;
    }

    if (right && !left && !top && bot) {
        line = game.SelectTargetLineLayer.create(cellSprite.x, cellSprite.y, 'select' + selectClass + '_2');
        line.anchor.setTo(1, 0);
        line.angle = -90;
    }

    if (right && !left && top && !bot) {
        line = game.SelectTargetLineLayer.create(cellSprite.x, cellSprite.y, 'select' + selectClass + '_2');
        line.anchor.setTo(1, 1);
        line.scale.x *= -1;
        line.scale.y *= -1;
    }

    if (!right && left && !top && bot) {
        game.SelectTargetLineLayer.create(cellSprite.x, cellSprite.y, 'select' + selectClass + '_2');
    }

    if (right && left && !top && !bot) {
        game.SelectTargetLineLayer.create(cellSprite.x, cellSprite.y, 'select' + selectClass + '_4');
    }

    if (!right && !left && top && bot) {
        line = game.SelectTargetLineLayer.create(cellSprite.x, cellSprite.y, 'select' + selectClass + '_4');
        line.anchor.setTo(1, 0);
        line.angle = -90;
    }

    if (!right && !left && top && !bot) {
        line = game.SelectTargetLineLayer.create(cellSprite.x, cellSprite.y, 'select' + selectClass + '_3');
        line.anchor.setTo(1, 1);
        line.angle = -180;
    }

    if (!right && left && !top && !bot) {
        line = game.SelectTargetLineLayer.create(cellSprite.x, cellSprite.y, 'select' + selectClass + '_3');
        line.anchor.setTo(0, 1);
        line.angle = 90;
    }

    if (right && !left && !top && !bot) {
        line = game.SelectTargetLineLayer.create(cellSprite.x, cellSprite.y, 'select' + selectClass + '_3');
        line.anchor.setTo(1, 0);
        line.angle = -90;
    }

    if (!right && !left && !top && bot) {
        game.SelectTargetLineLayer.create(cellSprite.x, cellSprite.y, 'select' + selectClass + '_3');
    }

    if (!right && !left && !top && !bot) {
        game.SelectTargetLineLayer.create(cellSprite.x, cellSprite.y, 'select' + selectClass + '_5');
    }
}