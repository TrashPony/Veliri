function SelectMoveCoordinateCreate(jsonMessage) {
    var moveCoordinate = JSON.parse(jsonMessage).move;

    for (var x in moveCoordinate) {
        if (moveCoordinate.hasOwnProperty(x)) {
            for (var y in moveCoordinate[x]) {
                if (moveCoordinate[x].hasOwnProperty(y)) {

                    var cellSprite = game.map.OneLayerMap[moveCoordinate[x][y].x][moveCoordinate[x][y].y].sprite;
                    var selectSprite = MarkZone(cellSprite, moveCoordinate, x, y, 'Move');

                    selectSprite.PlaceX = moveCoordinate[x][y].x;
                    selectSprite.PlaceY = moveCoordinate[x][y].y;

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
    /*field.send(JSON.stringify({
        event: "PlaceUnit",
        unit_id: Number(selectSprite.UnitID),
        x: Number(selectSprite.PlaceX),
        y: Number(selectSprite.PlaceY)
    }));*/

    RemoveSelect()
}