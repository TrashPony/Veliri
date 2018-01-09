function InitMoveUnit(jsonMessage) {

    DelMoveCoordinate(); // убираем выделение ячек

    //TODO {"event":"MoveUnit","user_name":"user","path_nodes":[{"Type":"","Texture":"","X":2,"Y":4,"State":0,"H":0,"G":0,"F":0,"Parent":null}]}
    var patchNodes = JSON.parse(jsonMessage).path_nodes;     // берем масив данных очереди перемещения юнита
    //var action = JSON.parse(jsonMessage).unit_action;
    //var errorMove = JSON.parse(jsonMessage).error;

    var x = JSON.parse(jsonMessage).unit_x;     // берем координаты юнита
    var y = JSON.parse(jsonMessage).unit_y;

    var unit = units[x + ":" + y];    // берем юнита
    unit.patchNodes = patchNodes;     // добавляем юниту путь
    var firstNode = unit.patchNodes.shift();     // берем первый пункт назначения
    unit.movePoint = cells[firstNode.x + ":" + firstNode.y];   // берем спрайт пункта назначения и кладем в текущую ноду к которой идет юнит

    unit.rotation = game.physics.arcade.angleToXY(unit, unit.movePoint.x + tileWidth / 2, unit.movePoint.y + tileWidth / 2); // поворачиваем юнита
    unit.body.velocity = game.physics.arcade.velocityFromAngle(unit.angle, UNIT_SPEED); // устанавливаем скорость
    /*if (action === "false") { // выделяем походившего юнита
        unit.tint = 0x757575;
    }*/
}

function MoveUnit() {
    for (var unitID in units) {
        if (units.hasOwnProperty(unitID)) {
            var unit = units[unitID];

            if (units[unitID].movePoint == null) { // если у юнита больше нет цели перемещения выставляем ему скорость движения и поворота 0
                unit.body.angularVelocity = 0;
                unit.body.velocity.x = 0;
                unit.body.velocity.y = 0;
            } else {

                var xTarget = Math.round(unit.movePoint.x + tileWidth / 2),
                    yTarget = Math.round(unit.movePoint.y + tileWidth / 2);
                var xUnit = Math.round(unit.x),
                    yUnit = Math.round(unit.y);

                if ((xUnit - TARGET_MOVE_RANGE < xTarget && xTarget < xUnit + TARGET_MOVE_RANGE) &&
                    (yUnit - TARGET_MOVE_RANGE < yTarget && yTarget < yUnit + TARGET_MOVE_RANGE)) { // если юнит стоит рядом с целью в приемлемом диапазоне то считаем что он достиг цели

                    unit.id = unit.movePoint.id; // присваиваем юниту новый ид соотвевующей ячейки на поле
                    delete units.unitID;         // удаляем его из масива под старым ид
                    units[unit.id] = unit;       // создаем под новым

                    if (unit.patchNodes.length > 0) {
                        var nextNode = unit.patchNodes.shift();     // берем следующий пункт назначения
                        unit.movePoint = cells[nextNode.x + ":" + nextNode.y];     // берем следующий спрайт пункта назначения и кладем в текущую ноду к которой идет юнит
                        unit.rotation = game.physics.arcade.angleToXY(unit, unit.movePoint.x + tileWidth / 2, unit.movePoint.y + tileWidth / 2); // поворачиваем юнита
                        unit.body.velocity = game.physics.arcade.velocityFromAngle(unit.angle, UNIT_SPEED); // устанавливаем скорость
                    } else {
                        unit.movePoint = null;
                    }
                }
            }
        }
    }
}