function SelectTargetCoordinateCreate(jsonMessage) {
    var targetCoordinates = JSON.parse(jsonMessage).targets;

    var event = JSON.parse(jsonMessage).event;

    var unitX = JSON.parse(jsonMessage).unit.x;
    var unitY = JSON.parse(jsonMessage).unit.y;
    var unitID = JSON.parse(jsonMessage).unit.id;

    for (var x in targetCoordinates) {
        if (targetCoordinates.hasOwnProperty(x)) {
            for (var y in targetCoordinates[x]) {
                if (targetCoordinates[x].hasOwnProperty(y)) {
                    var cellSprite = game.map.OneLayerMap[targetCoordinates[x][y].x][targetCoordinates[x][y].y].sprite;

                    if (event === "GetFirstTargets") {
                        MarkZone(cellSprite, targetCoordinates, x, y, 'Target', false, game.SelectTargetLineLayer, null);
                    }

                    if (event === "GetTargets") {
                        var selectSprite = MarkZone(cellSprite, targetCoordinates, x, y, 'Target', true, game.SelectTargetLineLayer, "target");

                        selectSprite.TargetX = targetCoordinates[x][y].x;
                        selectSprite.TargetY = targetCoordinates[x][y].y;

                        selectSprite.unitX = unitX;
                        selectSprite.unitY = unitY;
                        selectSprite.UnitID = unitID;

                        selectSprite.inputEnabled = true;

                        selectSprite.events.onInputDown.add(SelectTarget, selectSprite);
                        selectSprite.events.onInputOver.add(animateTargetCoordinate, selectSprite);
                        selectSprite.events.onInputOut.add(stopAnimateCoordinate, selectSprite);

                        selectSprite.input.priorityID = 1; // утсанавливает повышеный приоритет среди спрайтов на которых мышь

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
    }
}