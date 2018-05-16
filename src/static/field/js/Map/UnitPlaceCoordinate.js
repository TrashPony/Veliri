function SelectCoordinateUnitCreate(jsonMessage) {
    console.log(jsonMessage);

    var place_coordinate = JSON.parse(jsonMessage).place_coordinate;
    var unitID = JSON.parse(jsonMessage).unit.id;

    for (var i = 0; i < place_coordinate.length; i++) {
        var cellSprite = game.map.OneLayerMap[place_coordinate[i].x][place_coordinate[i].y].sprite;
        var selectSprite = game.make.sprite(0, 0, 'selectCreate');

        selectSprite.PlaceX = place_coordinate[i].x;
        selectSprite.PlaceY = place_coordinate[i].y;
        selectSprite.UnitID = unitID;

        selectSprite.inputEnabled = true;
        selectSprite.events.onInputDown.add(SelectPlaceCoordinate, selectSprite);
        selectSprite.events.onInputOut.add(TipOff);

        cellSprite.addChild(selectSprite);
        game.map.selectSprites.push(selectSprite);
    }
}

function SelectPlaceCoordinate(selectSprite) {

    field.send(JSON.stringify({
        event: "PlaceUnit",
        unit_id: Number(selectSprite.UnitID),
        x: Number(selectSprite.PlaceX),
        y: Number(selectSprite.PlaceY)
    }));

    RemoveSelectCoordinateUnitCreate()
}

function RemoveSelectCoordinateUnitCreate() {
    while (game.map.selectSprites.length > 0) {
        var selectSprite = game.map.selectSprites.shift();
        selectSprite.destroy();
    }
}