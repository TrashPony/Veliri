function animateMoveCoordinate(coordinate) {
    coordinate.animations.add('select');
    coordinate.animations.play('select', 5, true);

    field.send(JSON.stringify({
        event: "GetTargetZone",
        x: Number(coordinate.unitX),
        y: Number(coordinate.unitY),
        to_x: Number(coordinate.MoveX),
        to_y: Number(coordinate.MoveY)
    }));
    game.SelectLineLayer.visible = false;
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