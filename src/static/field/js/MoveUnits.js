function InitMoveUnit(jsonMessage) {

    DelMoveCoordinate(); // убираем выделение ячек
    var errorMove = JSON.parse(jsonMessage).error;
    if (errorMove === null || errorMove === "") {
        var owner = JSON.parse(jsonMessage).unit.owner;

        if (owner === MY_ID) {
            MoveMyUnit(jsonMessage, owner)
        } else {
            MoveHostileUnit(jsonMessage, owner)
        }
    }
}

function MoveHostileUnit(jsonMessage, owner) {
    var patchNodes = JSON.parse(jsonMessage).path_nodes;     // берем масив данных очереди перемещения юнита

    while (patchNodes.length > 0) {
        var firstNode = patchNodes.shift();
        var unit;
        if (firstNode.type === "visible") {                  // ищем первую ячейку где мы видим юнита
            var lastCell = patchNodes[patchNodes.length - 1];
            unit = units[firstNode.x + ":" + firstNode.y];
            if (unit) {
                unit.patchNodes = patchNodes;                    // добавляем юниту путь
                unit.owner = owner;
            } else {

                var newUnit = JSON.parse(jsonMessage).unit;
                newUnit.x = firstNode.x;
                newUnit.y = firstNode.y;

                CreateUnit(newUnit);

                unit = units[firstNode.x + ":" + firstNode.y];
                unit.patchNodes = patchNodes;                    // добавляем юниту путь
                unit.owner = owner;
            }
            MarkLastPathCell(unit, lastCell);
            CheckPath(unit);
            break;
        }
    }
}

function MoveMyUnit(jsonMessage, owner) {
    var patchNodes = JSON.parse(jsonMessage).path_nodes;     // берем масив данных очереди перемещения юнита
    var watchNode = JSON.parse(jsonMessage).watch_node;      // берем карту с видимости на каждой точке пути
    var action = JSON.parse(jsonMessage).unit.action;
    var unitNode = patchNodes.shift();                       // первая ячейка пути это ячейка моба
    var unit = units[unitNode.x + ":" + unitNode.y];         // берем юнита

    var lastCell = patchNodes[patchNodes.length - 1];
    MarkLastPathCell(unit, lastCell);                              // помечаем ячейку куда идет моб

    unit.patchNodes = patchNodes;                            // добавляем юниту путь
    unit.watchNode = watchNode;                              // кладем туда видимост на каждую клетку для отрисовки
    unit.action = action;                                    // докидываем в него статус готовности
    unit.owner = owner;

    CheckPath(unit);
}

function MarkLastPathCell(unit, cellState) {
    unit.lastCell = cellState;

    var x = cellState.x;
    var y = cellState.y;

    var mark = game.add.sprite(0, 0, 'MarkMoveLastCell'); // создаем метку
    mark.scale.set(.32);
    mark.alpha = 0.8;
    mark.z = 1;

    if (cells[x + ":" + y]) {
        cells[x + ":" + y].addChild(mark);
    }
}

function DeleteMarkLastPathCell(cellState) {
    var x = cellState.x;
    var y = cellState.y;
    var mark = cells[x + ":" + y].getChildAt(0);

    mark.destroy();
}

function CheckPath(unit) {
    var firstNode = unit.patchNodes.shift();                      // берем первый пункт назначения
    var targetID;

    if (unit.patchNodes.length === 0) {
        DeleteMarkLastPathCell(unit.lastCell); // удаляем метку
        unit.lastCell = null;
    }

    if (unit.owner === MY_ID) {
        targetID = firstNode.x + ":" + firstNode.y;
        MoveToCell(unit, targetID);
        UpdateWachZone(unit, targetID);
    } else {
        if (firstNode.type === "hide") {

            if (firstNode.type === "hide") {
                unit.visible = false;
                if (unit.patchNodes.length === 0) {
                    delete units[unit.id];
                    unit.destroy();
                }
            }

            while (unit.patchNodes.length > 0) {
                firstNode = unit.patchNodes.shift();
                console.log(unit.patchNodes.length);

                if (firstNode.type === "visible") {
                    targetID = firstNode.x + ":" + firstNode.y;
                    MoveToCell(unit, targetID);
                    break
                }
            }
        } else {
            unit.visible = true;
            targetID = firstNode.x + ":" + firstNode.y;
            MoveToCell(unit, targetID);
        }
    }
}

function MoveToCell(unit, targetID) {
    unit.movePoint = cells[targetID];                             // берем спрайт пункта назначения и кладем в текущую ноду к которой идет юнит
    unit.rotation = game.physics.arcade.angleToXY(unit, unit.movePoint.x + tileWidth / 2, unit.movePoint.y + tileWidth / 2); // поворачиваем юнита
    console.log(unit.angle);
    console.log(unit.rotation);
    unit.body.velocity = game.physics.arcade.velocityFromAngle(unit.angle, UNIT_SPEED); // устанавливаем скорость
}

function UpdateWachZone(unit, targetID) {
    if (unit.watchNode.hasOwnProperty(targetID)) {

        var watch = unit.watchNode;

        var closeCoordinate = watch[targetID].close_coordinate;
        var openCoordinate = watch[targetID].open_coordinate;
        var openUnits = watch[targetID].open_unit;
        var openStructure = watch[targetID].open_structure;

        if (closeCoordinate) {
            DeleteCoordinates(closeCoordinate);
        }

        if (openCoordinate) {
            OpenCoordinates(openCoordinate);
        }

        if (openUnits) {
            while (openUnits.length > 0) {
                var openUnit = openUnits.shift();
                CreateUnit(openUnit)
            }
        }

        if (openStructure) {
            // TODO добавить структуры
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

                    unit.id = unit.movePoint.id;  // присваиваем юниту новый ид соотвевующей ячейки на поле
                    delete units[unitID];         // удаляем его из масива под старым ид
                    units[unit.id] = unit;        // создаем под новым

                    if (unit.patchNodes.length > 0) {
                        CheckPath(unit);
                    } else {
                        unit.movePoint = null;
                        if (!unit.action) {
                            unit.tint = 0x757575;
                        }
                    }
                }
            }
        }
    }
}