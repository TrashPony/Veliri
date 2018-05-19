function MarkZone(cellSprite, placeCoordinate, x, y) {
    var left = false;
    var right = false;
    var top = false;
    var bot = false;
    var sprite;

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

    /*console.log(x + ":" + y);

    console.log(right);
    console.log(left);
    console.log(top);
    console.log(bot);

    console.log("----------------------");*/

    if (right && left && top && bot) {
        return game.SelectLayer.create(cellSprite.x, cellSprite.y, 'selectEmpty');
    }

    if (right && left && !top && bot) {
        return game.SelectLayer.create(cellSprite.x, cellSprite.y, 'selectPlace_1');
    }

    if (right && !left && top && bot) {
        sprite = game.SelectLayer.create(cellSprite.x, cellSprite.y, 'selectPlace_1');
        sprite.anchor.setTo(1, 0);
        sprite.angle = -90;
        return sprite
    }

    if (!right && left && top && bot) {
        sprite = game.SelectLayer.create(cellSprite.x, cellSprite.y, 'selectPlace_1');
        sprite.anchor.setTo(0, 1);
        sprite.angle = 90;
        return sprite
    }

    if (right && left && top && !bot) {
        sprite = game.SelectLayer.create(cellSprite.x, cellSprite.y, 'selectPlace_1');
        sprite.anchor.setTo(0, 1);
        sprite.scale.y *= -1;
        return sprite
    }

    if (!right && left && top && !bot) {
        sprite = game.SelectLayer.create(cellSprite.x, cellSprite.y, 'selectPlace_2');
        sprite.anchor.setTo(0, 1);
        sprite.scale.y *= -1;
        return sprite
    }

    if (right && !left && !top && bot) {
        sprite = game.SelectLayer.create(cellSprite.x, cellSprite.y, 'selectPlace_2');
        sprite.anchor.setTo(1, 0);
        sprite.angle = -90;
        return sprite
    }

    if (right && !left && top && !bot) {
        sprite = game.SelectLayer.create(cellSprite.x, cellSprite.y, 'selectPlace_2');
        sprite.anchor.setTo(1, 1);
        sprite.scale.x *= -1;
        sprite.scale.y *= -1;
        return sprite
    }

    if (!right && left && !top && bot) {
        return game.SelectLayer.create(cellSprite.x, cellSprite.y, 'selectPlace_2');
    } else {
        return game.SelectLayer.create(cellSprite.x, cellSprite.y, 'selectEmpty');
    }
}