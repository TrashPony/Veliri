function InitMoveUnit(jsonMessage) {

    DelMoveCoordinate(); // убираем выделение ячек
    var errorMove = JSON.parse(jsonMessage).error;
    if (errorMove === "") {

        var patchNodes = JSON.parse(jsonMessage).path_nodes;     // берем масив данных очереди перемещения юнита
        var watchNode = JSON.parse(jsonMessage).watch_node;      // берем карту с видимости на каждой точке пути
        var action = JSON.parse(jsonMessage).unit_action;
        var x = JSON.parse(jsonMessage).unit_x;                  // берем координаты юнита
        var y = JSON.parse(jsonMessage).unit_y;

        var unit = units[x + ":" + y];                           // берем юнита
        unit.patchNodes = patchNodes;                            // добавляем юниту путь
        unit.watchNode = watchNode;                              // кладем туда видимост на каждую клетку для отрисовки
        unit.action = action;                                    // докидываем в него статус готовности

        MoveCell(unit);
    }
}

function MoveCell(unit) {
    var firstNode = unit.patchNodes.shift();                        // берем первый пункт назначения
    var targetID = firstNode.x + ":" + firstNode.y;
    unit.movePoint = cells[targetID];        // берем спрайт пункта назначения и кладем в текущую ноду к которой идет юнит
    unit.rotation = game.physics.arcade.angleToXY(unit, unit.movePoint.x + tileWidth / 2, unit.movePoint.y + tileWidth / 2); // поворачиваем юнита
    unit.body.velocity = game.physics.arcade.velocityFromAngle(unit.angle, UNIT_SPEED); // устанавливаем скорость

    UpdateWachZone(unit, targetID);
}

function UpdateWachZone(unit, targetID) {
    if (unit.watchNode.hasOwnProperty(targetID)) {

        var watch = unit.watchNode;

        var closeCoordinate = watch[targetID].close_coordinate;
        var openCoordinate = watch[targetID].open_coordinate;
        var openUnit = watch[targetID].open_unit;
        var openStructure = watch[targetID].open_structure;

        if (closeCoordinate) {
            DeleteCoordinates(closeCoordinate);
        }

        if (openCoordinate) {
            OpenCoordinates(openCoordinate);
        }

        if (openUnit) {

        }

        if (openStructure) {

        }
    }
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

                var xTarget = unit.movePoint.x + tileWidth / 2,
                    yTarget = unit.movePoint.y + tileWidth / 2;
                var xUnit = unit.x,
                    yUnit = unit.y;

                if ((xUnit - TARGET_MOVE_RANGE < xTarget && xTarget < xUnit + TARGET_MOVE_RANGE) &&
                    (yUnit - TARGET_MOVE_RANGE < yTarget && yTarget < yUnit + TARGET_MOVE_RANGE)) { // если юнит стоит рядом с целью в приемлемом диапазоне то считаем что он достиг цели

                    unit.id = unit.movePoint.id; // присваиваем юниту новый ид соотвевующей ячейки на поле
                    delete units[unitID];         // удаляем его из масива под старым ид
                    units[unit.id] = unit;       // создаем под новым

                    if (unit.patchNodes.length > 0) {
                        MoveCell(unit);
                    } else {
                        unit.movePoint = null;
                    }
                }
            }
        }
    }
}