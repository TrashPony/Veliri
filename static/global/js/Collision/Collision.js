function CreateCollision(graphics, height, width, angle, squad) {
    if (!graphics) {
        squad.colision = game.add.graphics(0, 0);
        graphics = squad.colision;
    }

    graphics.clear();

    if (!squad.rectDebag) squad.rectDebag = game.add.graphics(0, 0);

    squad.rectDebag.clear();
    squad.rectDebag.lineStyle(1, 0xFF0000, 0.8);

    if (!squad.sprite) {
        return
    }

    let x = squad.sprite.x;
    let y = squad.sprite.y;

    let bodyRec = {
        sides: [
            {x1: x - width, y1: y - height, x2: x - width, y2: y + height},
            {x1: x - width, y1: y + height, x2: x + width, y2: y + height},
            {x1: x + width, y1: y + height, x2: x + width, y2: y - height},
            {x1: x + width, y1: y - height, x2: x - width, y2: y - height}
        ],
        centerX: squad.sprite.x,
        centerY: squad.sprite.y,
    };

    // поворачиваем квадрат по формуле (x0:y0 - центр)
    //X = (x — x0) * cos(alpha) — (y — y0) * sin(alpha) + x0;
    //Y = (x — x0) * sin(alpha) + (y — y0) * cos(alpha) + y0;

    let rotatePoint = function (x, y, x0, y0, rotate) {
        let alpha = rotate * Math.PI / 180;
        let newX = (x - x0) * Math.cos(alpha) - (y - y0) * Math.sin(alpha) + x0;
        let newY = (x - x0) * Math.sin(alpha) + (y - y0) * Math.cos(alpha) + y0;
        return {x: newX, y: newY}
    };

    let rotateSide = function (side, x0, y0, rotate) {
        let a = rotatePoint(side.x1, side.y1, x0, y0, rotate);
        let b = rotatePoint(side.x2, side.y2, x0, y0, rotate);

        side.x1 = a.x;
        side.y1 = a.y;

        side.x2 = b.x;
        side.y2 = b.y;
    };

    for (let i in bodyRec.sides) {
        rotateSide(bodyRec.sides[i], bodyRec.centerX, bodyRec.centerY, angle)
    }

    squad.rectDebag.moveTo(bodyRec.sides[0].x1, bodyRec.sides[0].y1);
    squad.rectDebag.lineTo(bodyRec.sides[0].x2, bodyRec.sides[0].y2);

    squad.rectDebag.lineTo(bodyRec.sides[1].x1, bodyRec.sides[1].y1);
    squad.rectDebag.lineTo(bodyRec.sides[1].x2, bodyRec.sides[1].y2);

    squad.rectDebag.lineTo(bodyRec.sides[2].x1, bodyRec.sides[2].y1);
    squad.rectDebag.lineTo(bodyRec.sides[2].x2, bodyRec.sides[2].y2);

    squad.rectDebag.lineTo(bodyRec.sides[3].x1, bodyRec.sides[3].y1);
    squad.rectDebag.lineTo(bodyRec.sides[3].x2, bodyRec.sides[3].y2);
    squad.rectDebag.endFill();
}