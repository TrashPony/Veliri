function CreateCollision(graphics, body, angle, squad) {
    if (!graphics){
        squad.colision = game.add.graphics(0, 0);
        graphics = squad.colision;
    }

    graphics.clear();

    if (!squad.rectDebag) squad.rectDebag = game.add.graphics(squad.sprite.x, squad.sprite.y);

    squad.rectDebag.clear();
    squad.rectDebag.lineStyle(1, 0xFF0000, 0.8);

    let height = body.height, width =  body.width;

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
}