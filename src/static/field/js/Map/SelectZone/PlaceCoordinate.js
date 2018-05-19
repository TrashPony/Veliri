function SelectCoordinateUnitCreate(jsonMessage) {

    var placeCoordinate = JSON.parse(jsonMessage).place_coordinate;
    var unitID = JSON.parse(jsonMessage).unit.id;

    console.log(placeCoordinate);

    for (var x in placeCoordinate) {
        if (placeCoordinate.hasOwnProperty(x)) {
            for (var y in placeCoordinate[x]) {
                if (placeCoordinate[x].hasOwnProperty(y)) {

                    var cellSprite = game.map.OneLayerMap[placeCoordinate[x][y].x][placeCoordinate[x][y].y].sprite;

                    var selectSprite = MarkZone(cellSprite, placeCoordinate, x, y);

                    selectSprite.PlaceX = placeCoordinate[x][y].x;
                    selectSprite.PlaceY = placeCoordinate[x][y].y;
                    selectSprite.UnitID = unitID;

                    selectSprite.inputEnabled = true;
                    selectSprite.events.onInputDown.add(SelectPlaceCoordinate, selectSprite);
                    selectSprite.events.onInputOut.add(TipOff);

                    game.map.selectSprites.push(selectSprite);
                }
            }
        }
    }

   /*for (var i = 0; i < placeCoordinate.length; i++) {



    }*/
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