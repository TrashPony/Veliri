function animateCoordinate(coordinate) {
    coordinate.animations.add('select');
    coordinate.animations.play('select', 5, true);

    if (game.Phase === "move") {
        field.send(JSON.stringify({
            event: "GetTargetZone",
            x: Number(coordinate.unitX),
            y: Number(coordinate.unitY),
            to_x: Number(coordinate.MoveX),
            to_y: Number(coordinate.MoveY)
        }));
        game.SelectLineLayer.visible = false;
    }
}

function stopAnimateCoordinate(coordinate) {
    coordinate.animations.getAnimation('select').stop(true);

    if (game.Phase === "move") {
        game.SelectLineLayer.visible = true;
        RemoveTargetLine();
    }
}