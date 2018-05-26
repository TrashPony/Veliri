function CreatePathToUnit(jsonMessage) {
    var error = JSON.parse(jsonMessage).error;

    if (error === null || error === "") {
        var path = JSON.parse(jsonMessage).path;     // берем масив данных очереди перемещения юнита

        var unit = GetGameUnitID(JSON.parse(jsonMessage).unit.id);         // берем юнита
        if (unit !== null) {
            var lastCell = path[path.length - 1].path_node;
            MarkLastPathCell(unit, lastCell);        // помечаем ячейку куда идет моб
            unit.path = path;                        // добавляем юниту путь
            CheckPath(unit);
        }
    }
}

/*function MoveHostileUnit(jsonMessage, owner) {
    var patchNodes = JSON.parse(jsonMessage).path_nodes;     // берем масив данных очереди перемещения юнита

    while (patchNodes.length > 0) {
        var firstNode = patchNodes.shift();
        var unit;
        if (firstNode.type === "visible") {                  // ищем первую ячейку где мы видим юнита
            var lastCell = patchNodes[patchNodes.length - 1];
            unit = units[firstNode.x + ":" + firstNode.y];
            if (unit) {
                unit.patchNodes = patchNodes;                // добавляем юниту путь
                unit.owner = owner;
            } else {

                var newUnit = JSON.parse(jsonMessage).unit;
                newUnit.x = firstNode.x;
                newUnit.y = firstNode.y;

                CreateUnit(newUnit);

                unit = units[firstNode.x + ":" + firstNode.y];
                unit.patchNodes = patchNodes;                  // добавляем юниту путь
                unit.owner = owner;
            }
            MarkLastPathCell(unit, lastCell);
            CheckPath(unit);
            break;
        }
    }
}*/

function CheckPath(unit) {
    var pathNode = unit.path.shift();   // берем первый пункт назначения

    if (unit.path.length === 0) {
        DeleteMarkLastPathCell(unit.lastCell); // удаляем метку
        unit.lastCell = null;
    }

    if (unit.sprite) {
        unit.movePoint = pathNode.path_node;
        unit.rotate = pathNode.unit_rotate;
        unit.watch = pathNode.watch_node;
    }

    /* else {
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
    }*/
}

function MoveToCell(unit) {
    unit.sprite.body.velocity = game.physics.arcade.velocityFromAngle(unit.spriteAngle, 100); // устанавливаем скорость
    unit.shadow.body.velocity = game.physics.arcade.velocityFromAngle(unit.spriteAngle, 100); // устанавливаем скорость
}

function StopUnit(unit) {
    unit.sprite.body.angularVelocity = 0;
    unit.sprite.body.velocity.x = 0;
    unit.sprite.body.velocity.y = 0;

    unit.shadow.body.angularVelocity = 0;
    unit.shadow.body.velocity.x = 0;
    unit.shadow.body.velocity.y = 0;
}

function MoveUnit() {

    for (var x in game.units) {
        if (game.units.hasOwnProperty(x)) {
            for (var y in game.units[x]) {
                if (game.units[x].hasOwnProperty(y) && game.units[x][y].sprite && game.units[x][y].shadow) {
                    var unit = game.units[x][y];

                    if (unit.movePoint == null) { // если у юнита больше нет цели перемещения выставляем ему скорость движения и поворота 0
                        StopUnit(unit);
                    } else {
                        if (unit.spriteAngle === unit.rotate) {
                            MoveToCell(unit);
                        }
                        var xTarget = (unit.movePoint.x * 100) + 50,
                            yTarget = (unit.movePoint.y * 100) + 50;

                        var xUnit = unit.sprite.x,
                            yUnit = unit.sprite.y;

                        if ((xUnit - 20 < xTarget && xTarget < xUnit + 20) &&
                            (yUnit - 20 < yTarget && yTarget < yUnit + 20)) { // если юнит стоит рядом с целью в приемлемом диапазоне то считаем что он достиг цели

                            delete game.units[unit.x][unit.y];

                            unit.x = unit.movePoint.x;
                            unit.y = unit.movePoint.y;

                            addToGameUnit(unit);

                            unit.movePoint = null;
                            UpdateWatchZone(unit.watch);

                            if (unit.path.length > 0) {
                                StopUnit(unit);
                                CheckPath(unit);
                            } else {
                                if (unit.action) {
                                    unit.sprite.tint = 0x757575;
                                }
                            }
                        }
                    }
                }
            }
        }
    }
}