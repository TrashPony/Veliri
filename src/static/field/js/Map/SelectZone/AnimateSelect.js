function animateMoveCoordinate(coordinate) {
    coordinate.animations.add('select');
    coordinate.animations.play('select', 5, true);

    RemoveTargetLine();

    field.send(JSON.stringify({
        event: "GetTargetZone",
        unit_id: Number(coordinate.UnitID),
        q: Number(coordinate.unitQ),
        r: Number(coordinate.unitR),
        to_q: Number(coordinate.MoveQ),
        to_r: Number(coordinate.MoveR)
    }));

    field.send(JSON.stringify({
        event: "GetPreviewPath",
        unit_id: Number(coordinate.UnitID),
        q: Number(coordinate.unitQ),
        r: Number(coordinate.unitR),
        to_q: Number(coordinate.MoveQ),
        to_r: Number(coordinate.MoveR)
    }));

    game.SelectLineLayer.visible = false;

    if (coordinate.UnitMS) {
        let centerCoordinate = game.map.OneLayerMap[coordinate.MoveQ][coordinate.MoveR];
        let circleCoordinates = getRadius(centerCoordinate.x, centerCoordinate.y, centerCoordinate.z, 1);

        for (let i in circleCoordinates) {
            let q = circleCoordinates[i].Q;
            let r = circleCoordinates[i].R;
            if (game.map.OneLayerMap.hasOwnProperty(q) && game.map.OneLayerMap[q].hasOwnProperty(r)) {
                let animateCoordinate = game.map.OneLayerMap[q][r];
                let selectSprite = MarkZone(animateCoordinate.sprite, circleCoordinates, q, r, 'Move', true, game.SelectRangeLayer, "move", game.SelectRangeLayer);
                selectSprite.animations.add('select');
                selectSprite.animations.play('select', 5, true);
            }
        }
    }
}

function animatePlaceCoordinate(coordinate) {
    coordinate.animations.add('select');
    coordinate.animations.play('select', 5, true);
}

function animateTargetCoordinate(coordinate) {

    let unit = GetGameUnitXY(coordinate.unitQ,coordinate.unitR);
    let targetUnit = GetGameUnitXY(coordinate.TargetQ,coordinate.TargetR);

    if (targetUnit) {
        let style = {font: "24px Finger Paint", fill: "#C00"};
        let damageText = game.add.text(targetUnit.sprite.x + 20, targetUnit.sprite.y - 50, getMinMaxDamage(unit, targetUnit), style);
        damageText.setShadow(1, -1, 'rgba(0,0,0,0.5)', 0);
        coordinate.damageText = damageText;
    }

    coordinate.animations.add('select', [1, 2]);
    coordinate.animations.play('select', 3, true);
}

function stopAnimateCoordinate(coordinate) {
    coordinate.animations.getAnimation('select').stop(false);
    coordinate.animations.frame = 0;

    if (coordinate.damageText) {
        coordinate.damageText.destroy();
    }

    if (game.Phase === "move") {
        game.SelectLineLayer.visible = true;
        RemoveTargetLine();
        RemoveSelectRangeCoordinate();
    }
}