function SelectMoveCoordinateCreate(jsonMessage) {

    let moveCoordinate = JSON.parse(jsonMessage).move;

    let unitX = JSON.parse(jsonMessage).unit.x;
    let unitY = JSON.parse(jsonMessage).unit.y;
    let unitID = JSON.parse(jsonMessage).unit.id;

    game.SelectLineLayer.visible = true;

    for (let x in moveCoordinate) {
        if (moveCoordinate.hasOwnProperty(x)) {
            for (let y in moveCoordinate[x]) {
                if (moveCoordinate[x].hasOwnProperty(y)) {

                    let cellSprite = game.map.OneLayerMap[moveCoordinate[x][y].x][moveCoordinate[x][y].y].sprite;
                    let selectSprite = MarkZone(cellSprite, moveCoordinate, x, y, 'Move', true, game.SelectLineLayer, "move");

                    selectSprite.MoveX = moveCoordinate[x][y].x;
                    selectSprite.MoveY = moveCoordinate[x][y].y;

                    selectSprite.unitX = unitX;
                    selectSprite.unitY = unitY;
                    selectSprite.UnitID = unitID;

                    selectSprite.inputEnabled = true;
                    selectSprite.events.onInputDown.add(SelectMoveCoordinate, selectSprite);
                    selectSprite.events.onInputOver.add(animateMoveCoordinate, selectSprite);
                    selectSprite.events.onInputOut.add(stopAnimateCoordinate, selectSprite);
                    selectSprite.input.priorityID = 1; // утсанавливает повышеный приоритет среди спрайтов на которых мышь

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