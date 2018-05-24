function SelectMoveCoordinateCreate(jsonMessage) {

    var moveCoordinate = JSON.parse(jsonMessage).move;

    var unitX = JSON.parse(jsonMessage).unit.x;
    var unitY = JSON.parse(jsonMessage).unit.y;
    game.SelectLineLayer.visible = true;

    for (var x in moveCoordinate) {
        if (moveCoordinate.hasOwnProperty(x)) {
            for (var y in moveCoordinate[x]) {
                if (moveCoordinate[x].hasOwnProperty(y)) {

                    var cellSprite = game.map.OneLayerMap[moveCoordinate[x][y].x][moveCoordinate[x][y].y].sprite;
                    var selectSprite = MarkZone(cellSprite, moveCoordinate, x, y, 'Move', true, game.SelectLineLayer);

                    selectSprite.MoveX = moveCoordinate[x][y].x;
                    selectSprite.MoveY = moveCoordinate[x][y].y;

                    selectSprite.unitX = unitX;
                    selectSprite.unitY = unitY;

                    selectSprite.inputEnabled = true;

                    selectSprite.events.onInputDown.add(SelectMoveCoordinate, selectSprite);
                    selectSprite.events.onInputOver.add(animateCoordinate, selectSprite);
                    selectSprite.events.onInputOut.add(stopAnimateCoordinate, selectSprite);

                    game.map.selectSprites.push(selectSprite);
                }
            }
        }
    }
}

function SelectMoveCoordinate(selectSprite) {
    if (game.input.activePointer.leftButton.isDown) {

        field.send(JSON.stringify({
            event: "MoveUnit",
            unit_id: Number(selectSprite.UnitID),
            x: Number(selectSprite.unitX),
            y: Number(selectSprite.unitY),
            to_x: Number(selectSprite.MoveX),
            to_y: Number(selectSprite.MoveY)
        }));

        RemoveSelect()
    }
}