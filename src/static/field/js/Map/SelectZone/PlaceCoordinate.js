function SelectCoordinateUnitCreate(jsonMessage) {

    var placeCoordinate = JSON.parse(jsonMessage).place_coordinate;
    var unitID = JSON.parse(jsonMessage).unit.id;

    for (var x in placeCoordinate) {
        if (placeCoordinate.hasOwnProperty(x)) {
            for (var y in placeCoordinate[x]) {
                if (placeCoordinate[x].hasOwnProperty(y)) {

                    var cellSprite = game.map.OneLayerMap[placeCoordinate[x][y].x][placeCoordinate[x][y].y].sprite;

                    var selectSprite = MarkZone(cellSprite, placeCoordinate, x, y, 'Place');

                    selectSprite.PlaceX = placeCoordinate[x][y].x;
                    selectSprite.PlaceY = placeCoordinate[x][y].y;
                    selectSprite.UnitID = unitID;

                    selectSprite.inputEnabled = true;
                    selectSprite.events.onInputDown.add(SelectPlaceCoordinate, selectSprite);
                    selectSprite.events.onInputOver.add(animateCoordinate, selectSprite);
                    selectSprite.events.onInputOut.add(stopAnimateCoordinate, selectSprite);

                    game.map.selectSprites.push(selectSprite);
                }
            }
        }
    }
}

function SelectPlaceCoordinate(selectSprite) {
    field.send(JSON.stringify({
        event: "PlaceUnit",
        unit_id: Number(selectSprite.UnitID),
        x: Number(selectSprite.PlaceX),
        y: Number(selectSprite.PlaceY)
    }));

    RemoveSelect()
}

function animateCoordinate(coordinate) {
    coordinate.animations.add('select');
    coordinate.animations.play('select', 5, true);

}

function stopAnimateCoordinate(coordinate) {
    coordinate.animations.getAnimation('select').stop(true);
}