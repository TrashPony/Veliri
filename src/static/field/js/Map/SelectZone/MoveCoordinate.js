function SelectMoveCoordinateCreate(jsonMessage) {

    let moveCoordinate = JSON.parse(jsonMessage).move;

    let unitQ = JSON.parse(jsonMessage).unit.q;
    let unitR = JSON.parse(jsonMessage).unit.r;
    let unitID = JSON.parse(jsonMessage).unit.id;

    game.SelectLineLayer.visible = true;

    for (let q in moveCoordinate) {
        if (moveCoordinate.hasOwnProperty(q)) {
            for (let r in moveCoordinate[q]) {
                if (moveCoordinate[q].hasOwnProperty(r)) {

                    let cellSprite = game.map.OneLayerMap[q][r].sprite;

                    let selectSprite = MarkZone(cellSprite, moveCoordinate, q, r, 'Move', true, game.SelectLineLayer, "move");

                    selectSprite.MoveQ = q;
                    selectSprite.MoveR = r;

                    selectSprite.unitQ = unitQ;
                    selectSprite.unitR = unitR;
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
            q: Number(selectSprite.unitQ),
            r: Number(selectSprite.unitR),
            to_q: Number(selectSprite.MoveQ),
            to_r: Number(selectSprite.MoveR)
        }));

        RemoveSelect()
    }
}