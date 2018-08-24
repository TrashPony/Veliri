function SelectTargetCoordinateCreate(jsonMessage) {
    let targetCoordinates = JSON.parse(jsonMessage).targets;

    let event = JSON.parse(jsonMessage).event;

    let unitX = JSON.parse(jsonMessage).unit.x;
    let unitY = JSON.parse(jsonMessage).unit.y;
    let unitID = JSON.parse(jsonMessage).unit.id;

    for (let x in targetCoordinates) {
        if (targetCoordinates.hasOwnProperty(x)) {
            for (let y in targetCoordinates[x]) {
                if (targetCoordinates[x].hasOwnProperty(y)) {
                    let cellSprite = game.map.OneLayerMap[targetCoordinates[x][y].x][targetCoordinates[x][y].y].sprite;

                    if (event === "GetFirstTargets") {
                        MarkZone(cellSprite, targetCoordinates, x, y, 'Target', false, game.SelectTargetLineLayer, null);
                    }

                    if (event === "GetTargets") {
                        let selectSprite = MarkZone(cellSprite, targetCoordinates, x, y, 'Target', true, game.SelectTargetLineLayer, "target");

                        selectSprite.TargetX = targetCoordinates[x][y].x;
                        selectSprite.TargetY = targetCoordinates[x][y].y;

                        selectSprite.unitX = unitX;
                        selectSprite.unitY = unitY;
                        selectSprite.UnitID = unitID;

                        selectSprite.inputEnabled = true;
                        // Todo если в клетку куда показывает юзер есть юнит надо показывать сколько примерно отниметься хп
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