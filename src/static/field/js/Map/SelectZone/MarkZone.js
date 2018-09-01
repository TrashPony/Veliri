function MarkZone(cellSprite, placeCoordinate, x, y, selectClass, addEmpty, typeLine, selector) {
    let left = false;
    let right = false;
    let top = false;
    let bot = false;
    let sprite;
    let line;

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

    if (addEmpty) {
        if (selector === "move" || selector === "place") sprite = game.SelectLayer.create(cellSprite.x, cellSprite.y, 'selectEmpty');
        if (selector === "target") sprite = game.SelectLayer.create(cellSprite.x, cellSprite.y, 'selectTarget');
    }

    if (right && left && !top && bot) {
        line = typeLine.create(cellSprite.x, cellSprite.y, 'select' + selectClass + '_1');
    }

    if (right && !left && top && bot) {
        line = typeLine.create(cellSprite.x, cellSprite.y, 'select' + selectClass + '_1');
        line.anchor.setTo(1, 0);
        line.angle = -90;
    }

    if (!right && left && top && bot) {
        line = typeLine.create(cellSprite.x, cellSprite.y, 'select' + selectClass + '_1');
        line.anchor.setTo(0, 1);
        line.angle = 90;
    }

    if (right && left && top && !bot) {
        line = typeLine.create(cellSprite.x, cellSprite.y, 'select' + selectClass + '_1');
        line.anchor.setTo(0, 1);
        line.scale.y *= -1;
    }

    if (!right && left && top && !bot) {
        line = typeLine.create(cellSprite.x, cellSprite.y, 'select' + selectClass + '_2');
        line.anchor.setTo(0, 1);
        line.scale.y *= -1;
    }

    if (right && !left && !top && bot) {
        line = typeLine.create(cellSprite.x, cellSprite.y, 'select' + selectClass + '_2');
        line.anchor.setTo(1, 0);
        line.angle = -90;
    }

    if (right && !left && top && !bot) {
        line = typeLine.create(cellSprite.x, cellSprite.y, 'select' + selectClass + '_2');
        line.anchor.setTo(1, 1);
        line.scale.x *= -1;
        line.scale.y *= -1;
    }

    if (!right && left && !top && bot) {
        typeLine.create(cellSprite.x, cellSprite.y, 'select' + selectClass + '_2');
    }

    if (right && left && !top && !bot) {
        typeLine.create(cellSprite.x, cellSprite.y, 'select' + selectClass + '_4');
    }

    if (!right && !left && top && bot) {
        line = typeLine.create(cellSprite.x, cellSprite.y, 'select' + selectClass + '_4');
        line.anchor.setTo(1, 0);
        line.angle = -90;
    }

    if (!right && !left && top && !bot) {
        line = typeLine.create(cellSprite.x, cellSprite.y, 'select' + selectClass + '_3');
        line.anchor.setTo(1, 1);
        line.angle = -180;
    }

    if (!right && left && !top && !bot) {
        line = typeLine.create(cellSprite.x, cellSprite.y, 'select' + selectClass + '_3');
        line.anchor.setTo(0, 1);
        line.angle = 90;
    }

    if (right && !left && !top && !bot) {
        line = typeLine.create(cellSprite.x, cellSprite.y, 'select' + selectClass + '_3');
        line.anchor.setTo(1, 0);
        line.angle = -90;
    }

    if (!right && !left && !top && bot) {
        typeLine.create(cellSprite.x, cellSprite.y, 'select' + selectClass + '_3');
    }

    if (!right && !left && !top && !bot) {
        typeLine.create(cellSprite.x, cellSprite.y, 'select' + selectClass + '_5');
    }

    return sprite
}