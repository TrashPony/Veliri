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

function MoveHostileUnit(jsonMessage) {

 console.log(jsonMessage);

    var patchNodes = JSON.parse(jsonMessage).path;
    var unit = GetGameUnitID(JSON.parse(jsonMessage).unit.id);

    while (patchNodes.length > 0) {

        var firstNode = patchNodes.shift();

        // ищем первую ячейку где мы видим юнита
        if (firstNode.path_node.type !== "hide") {

            patchNodes.unshift(firstNode); // возвращаем точку

            if (unit) {
                unit.path = patchNodes;
            } else {
                unit = JSON.parse(jsonMessage).unit;
                unit.x = firstNode.path_node.x;
                unit.y = firstNode.path_node.y;
                CreateUnit(unit);

                unit.path = patchNodes;                // добавляем юниту путь
            }

            CheckPath(unit);
            break;
        }
    }
}

function CheckPath(unit) {
    var pathNode = unit.path.shift();   // берем первый пункт назначения

    if (pathNode.path_node.type === "hide") {
        StopUnit(unit);

        if (unit.path.length === 0) {
            unit.destroy = true;
        }

        return
    }

    if (pathNode.path_node.type === "inToFog") {
        if (unit.path.length === 0) {
            unit.destroy = true;
        }
    }

    //game.add.tween(unit.sprite).to({alpha: 0}, 1000, Phaser.Easing.Linear.None, true); // плавно затемняет спрайт хз как работает)
    //game.add.tween(unit.shadow).to({alpha: 0}, 1000, Phaser.Easing.Linear.None, true); // Tween позволяет изменять одно или несколько свойств целевого объекта в течение определенного периода времени

    /*while (unit.path.length > 0) {

        pathNode = unit.path.shift();

        if (pathNode.path_node.type !== "hide") {
            unit.movePoint = pathNode.path_node;
            unit.rotate = pathNode.unit_rotate;
            unit.watch = pathNode.watch_node;

           // unit.sprite.visible = true;
           // unit.shadow.visible = true;
            break
        }
    }*/

    if (unit.path.length === 0) {
        DeleteMarkLastPathCell(unit.lastCell); // удаляем метку
        unit.lastCell = null;
    }

    if (unit.sprite) {
        unit.movePoint = pathNode.path_node;
        unit.rotate = pathNode.unit_rotate;
        unit.watch = pathNode.watch_node;
    }
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

                            if (unit.movePoint.type === "inToFog") {
                                game.add.tween(unit.sprite).to({alpha: 0}, 1000, Phaser.Easing.Linear.None, true);
                                game.add.tween(unit.shadow).to({alpha: 0}, 1000, Phaser.Easing.Linear.None, true);
                            }

                            if (unit.movePoint.type === "outFog") {
                                game.add.tween(unit.sprite).to({alpha: 1}, 1000, Phaser.Easing.Linear.None, true);
                                game.add.tween(unit.shadow).to({alpha: 1}, 1000, Phaser.Easing.Linear.None, true);
                            }

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