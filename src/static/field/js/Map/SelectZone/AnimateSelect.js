function animateMoveCoordinate(coordinate) {
    coordinate.animations.add('select');
    coordinate.animations.play('select', 5, true);

    field.send(JSON.stringify({
        event: "GetTargetZone",
        q: Number(coordinate.unitQ),
        r: Number(coordinate.unitR),
        to_q: Number(coordinate.MoveQ),
        to_r: Number(coordinate.MoveR)
    }));
    game.SelectLineLayer.visible = false;

    if (coordinate.UnitMS) {

    }
}

function animatePlaceCoordinate(coordinate) {
    coordinate.animations.add('select');
    coordinate.animations.play('select', 5, true);
}

function animateTargetCoordinate(coordinate) {
    coordinate.animations.add('select', [1,2]);
    coordinate.animations.play('select', 3, true);
}

function stopAnimateCoordinate(coordinate) {
    coordinate.animations.getAnimation('select').stop(false);
    coordinate.animations.frame = 0;

    if (game.Phase === "move") {
        game.SelectLineLayer.visible = true;
        RemoveTargetLine();
    }
}