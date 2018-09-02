function SelectCoordinateUnitCreate(jsonMessage) {
    console.log(jsonMessage);
    RemoveSelect();

    let placeCoordinate = JSON.parse(jsonMessage).place_coordinate;
    let unitID = JSON.parse(jsonMessage).unit.id;

    for (let q in placeCoordinate) {
        if (placeCoordinate.hasOwnProperty(q)) {
            for (let r in placeCoordinate[q]) {
                if (placeCoordinate[q].hasOwnProperty(r)) {

                    let cellSprite = game.map.OneLayerMap[placeCoordinate[q][r].q][placeCoordinate[q][r].r].sprite;

                    let selectSprite = MarkZone(cellSprite, placeCoordinate, q, r, 'Place', true, game.SelectLineLayer, "place");

                    selectSprite.PlaceQ = placeCoordinate[q][r].q;
                    selectSprite.PlaceR = placeCoordinate[q][r].r;
                    selectSprite.UnitID = unitID;

                    selectSprite.inputEnabled = true;
                    selectSprite.events.onInputDown.add(SelectPlaceCoordinate, selectSprite);
                    selectSprite.events.onInputOver.add(animatePlaceCoordinate, selectSprite);
                    selectSprite.events.onInputOut.add(stopAnimateCoordinate, selectSprite);
                    selectSprite.input.priorityID = 1; // утсанавливает повышеный приоритет среди спрайтов на которых мышь

                    game.map.selectSprites.push(selectSprite);
                }
            }
        }
    }
}

function SelectPlaceCoordinate(selectSprite) {
    if (game.input.activePointer.leftButton.isDown) {
        field.send(JSON.stringify({
            event: "PlaceUnit",
            unit_id: Number(selectSprite.UnitID),
            q: Number(selectSprite.PlaceQ),
            r: Number(selectSprite.PlaceR)
        }));

        RemoveSelect()
    }
}