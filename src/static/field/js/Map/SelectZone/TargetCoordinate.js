function SelectTargetCoordinateCreate(jsonMessage) {

    console.log(jsonMessage);

    var targetCoordinates = JSON.parse(jsonMessage).targets;

    var unitX = JSON.parse(jsonMessage).unit.x;
    var unitY = JSON.parse(jsonMessage).unit.y;
    var unitID = JSON.parse(jsonMessage).unit.id;

    for (var x in targetCoordinates) {
        if (targetCoordinates.hasOwnProperty(x)) {
            for (var y in targetCoordinates[x]) {
                if (targetCoordinates[x].hasOwnProperty(y)) {
                    var cellSprite = game.map.OneLayerMap[targetCoordinates[x][y].x][targetCoordinates[x][y].y].sprite;

                    if (game.Phase === "move") {
                        MarkZone(cellSprite, targetCoordinates, x, y, 'Target', false, game.SelectTargetLineLayer);
                    }

                    if (game.Phase === "targeting") {
                        var selectSprite = MarkZone(cellSprite, targetCoordinates, x, y, 'Target', true, game.SelectTargetLineLayer);

                        selectSprite.TargetX = targetCoordinates[x][y].x;
                        selectSprite.TargetY = targetCoordinates[x][y].y;

                        selectSprite.unitX = unitX;
                        selectSprite.unitY = unitY;
                        selectSprite.UnitID = unitID;

                        selectSprite.inputEnabled = true;
                        selectSprite.events.onInputDown.add(SelectTarget, selectSprite);  // todo
                        selectSprite.events.onInputOver.add(animateCoordinate, selectSprite);
                        selectSprite.events.onInputOut.add(stopAnimateCoordinate, selectSprite);

                        game.map.selectSprites.push(selectSprite);
                    }
                }
            }
        }
    }
}

function SelectTarget(selectSprite) {
    if (game.input.activePointer.leftButton.isDown) {

        field.send(JSON.stringify({
            event: "SetTarget",
            unit_id: Number(selectSprite.UnitID),
            x: Number(selectSprite.unitX),
            y: Number(selectSprite.unitY),
            to_x: Number(selectSprite.TargetX),
            to_y: Number(selectSprite.TargetY)
        }));

        RemoveSelect()
    }
}