function InitMoveUnit(jsonMessage) {

    var error = JSON.parse(jsonMessage).error;

    if (error === null || error === "") {
        var patchNodes = JSON.parse(jsonMessage).path_nodes;     // берем масив данных очереди перемещения юнита
        var watchNodes = JSON.parse(jsonMessage).watch_nodes;    // берем карту с видимости на каждой точке пути
        var unitInfo = JSON.parse(jsonMessage).unit;

        MoveMyUnit(patchNodes, watchNodes, unitInfo)
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

function MoveMyUnit(patchNodes, watchNodes, unitInfo) {

    var unit = GetGameUnit(unitInfo.id);         // берем юнита\

    if (unit !== null) {
        var lastCell = patchNodes[patchNodes.length - 1];
        MarkLastPathCell(unit, lastCell);                        // помечаем ячейку куда идет моб

        unit.patchNodes = patchNodes;                            // добавляем юниту путь
        unit.watchNodes = watchNodes;                            // кладем туда видимост на каждую клетку для отрисовки
        console.log(unit.watchNodes);

        CheckPath(unit);
    }
}

function MarkLastPathCell(unit, cellState) {
    unit.lastCell = cellState;

    var x = cellState.x;
    var y = cellState.y;

    var mark = game.add.sprite(0, 0, 'MarkMoveLastCell'); // создаем метку
    mark.scale.set(.32);
    mark.alpha = 0.8;
    mark.z = 1;

    if (game.map.OneLayerMap[x][y].sprite) {
        game.map.OneLayerMap[x][y].sprite.addChild(mark);
    }
}

function DeleteMarkLastPathCell(cellState) {
    var x = cellState.x;
    var y = cellState.y;
    var mark = game.map.OneLayerMap[x][y].sprite.getChildAt(0);
    mark.destroy();
}

function CheckPath(unit) {
    var firstNode = unit.patchNodes.shift();   // берем первый пункт назначения
    var watchNode = unit.watchNodes[firstNode.x + ":" + firstNode.y];

    if (unit.patchNodes.length === 0) {
        DeleteMarkLastPathCell(unit.lastCell); // удаляем метку
        unit.lastCell = null;
    }

    if (unit.sprite) {
        MoveToCell(unit, firstNode);
        UpdateWatchZone(watchNode);
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

function MoveToCell(unit, targetID) {
    unit.sprite.movePoint = targetID;   //todo берем спрайт пункта назначения и кладем в текущую ноду к которой идет юнит
    unit.shadow.movePoint = targetID;   //todo я не знаю как правильно разместить тень

    unit.sprite.body.velocity = game.physics.arcade.velocityFromAngle(unit.sprite.angle, 100); // устанавливаем скорость
    unit.shadow.body.velocity = game.physics.arcade.velocityFromAngle(unit.sprite.angle, 100); // устанавливаем скорость
}

function MoveUnit() {

    for (var x in game.units) {
        if (game.units.hasOwnProperty(x)) {
            for (var y in game.units[x]) {
                if (game.units[x].hasOwnProperty(y) && game.units[x][y].sprite && game.units[x][y].shadow) {
                    var unit = game.units[x][y].sprite;
                    var shadow = game.units[x][y].shadow;

                    if (unit.movePoint == null) { // если у юнита больше нет цели перемещения выставляем ему скорость движения и поворота 0
                        unit.body.angularVelocity = 0;
                        unit.body.velocity.x = 0;
                        unit.body.velocity.y = 0;

                        shadow.body.angularVelocity = 0;
                        shadow.body.velocity.x = 0;
                        shadow.body.velocity.y = 0;
                    } else {

                        var xTarget = unit.movePoint.x + 100 / 2,
                            yTarget = unit.movePoint.y + 100 / 2;
                        var xUnit = unit.x,
                            yUnit = unit.y;

                        if ((xUnit - 50 < xTarget && xTarget < xUnit + 50) &&
                            (yUnit - 50 < yTarget && yTarget < yUnit + 50)) { // если юнит стоит рядом с целью в приемлемом диапазоне то считаем что он достиг цели

                            //unit.id = unit.movePoint.id;  // присваиваем юниту новый ид соотвевующей ячейки на поле
                            //delete units[unitID];         // удаляем его из масива под старым ид
                            //units[unit.id] = unit;        // создаем под новым

                            if (unit.patchNodes.length > 0) {
                                CheckPath(game.units[x][y]);
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
    }
}