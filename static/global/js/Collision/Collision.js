function CreateCollision(graphics, body, angle, squad) {
    graphics.clear();

    if (!squad.rectDebag) squad.rectDebag = game.add.graphics(squad.sprite.x, squad.sprite.y);

    squad.rectDebag.clear();
    squad.rectDebag.lineStyle(1, 0xFF0000, 0.8);

    let height = 25, width = 40;

    squad.rectDebag.moveTo(-width, -height);
    squad.rectDebag.lineTo(-width, +height);
    squad.rectDebag.lineTo(-height, +height);
    squad.rectDebag.lineTo(+width, +height);
    squad.rectDebag.lineTo(+width, +height);
    squad.rectDebag.lineTo(+width, -height);
    squad.rectDebag.lineTo(+height, -height);
    squad.rectDebag.lineTo(-width, -height);
    squad.rectDebag.endFill();

    squad.rectDebag.x = squad.sprite.x;
    squad.rectDebag.y = squad.sprite.y;
    squad.rectDebag.angle = angle;

    graphics.beginFill(0x00FF00, 0.5);
    graphics.arc(
        squad.sprite.x,
        squad.sprite.y,
        body.front_radius,
        game.math.degToRad(angle + body.right_front_angle),
        game.math.degToRad(angle - body.left_front_angle),
        true,
    );
    graphics.endFill();

    graphics.beginFill(0x00FF00, 0.5);
    graphics.arc(
        squad.sprite.x,
        squad.sprite.y,
        body.back_radius,
        game.math.degToRad(angle + (body.right_back_angle - 180)),
        game.math.degToRad(angle - (body.left_back_angle - 180)),
        true,
    );
    graphics.endFill();

    graphics.beginFill(0x0000ff, 0.5);
    graphics.arc(
        squad.sprite.x,
        squad.sprite.y,
        body.side_radius,
        game.math.degToRad(370),
        game.math.degToRad(0),
        true,
    );
    graphics.endFill();

    let rad = angle * Math.PI / 180;
    let bX = 43 * Math.cos(rad) + squad.sprite.x;
    let bY = 43 * Math.sin(rad) + squad.sprite.y;
    graphics.beginFill(0xFF0000, 0.5);
    graphics.drawCircle(bX, bY, 5);
}

function CreateMSGeo(squad) {
    squad.colision = game.add.graphics(0, 0);
    CreateCollision(squad.colision, squad.mather_ship.body, squad.mather_ship.rotate, squad)
}

function CreateOtherMSGeo(squad) {
    squad.colision = game.add.graphics(0, 0);
    CreateCollision(squad.colision, squad.body, squad.rotate, squad);
}