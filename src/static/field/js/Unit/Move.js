function CreatePathToUnit(jsonMessage) {
    let error = JSON.parse(jsonMessage).error;

    if (error === null || error === "") {

        let path = JSON.parse(jsonMessage).path;     // берем масив данных очереди перемещения юнита
        let unit = GetGameUnitID(JSON.parse(jsonMessage).unit.id);         // берем юнита
        unit.action = JSON.parse(jsonMessage).unit.action;

        if (unit !== null && path) {

            let lastCell = path[path.length - 1].path_node;
            MarkLastPathCell(unit, lastCell);        // помечаем ячейку куда идет моб
            unit.path = path;                        // добавляем юниту путь
            CheckPath(unit);

        } else {

            if (unit !== null && unit.action) {
                DeactivationUnit(unit);
                RemoveSelect();
            }
        }
    }
}

function MoveHostileUnit(jsonMessage) {

    let patchNodes = JSON.parse(jsonMessage).path;
    let unit = GetGameUnitID(JSON.parse(jsonMessage).unit.id);

    while (patchNodes.length > 0) {

        let firstNode = patchNodes.shift();

        // ищем первую ячейку где мы видим юнита
        if (firstNode.path_node.type !== "hide") {

            patchNodes.unshift(firstNode); // возвращаем точку

            if (unit) {
                unit.path = patchNodes;
            } else {
                unit = JSON.parse(jsonMessage).unit;
                unit.x = firstNode.path_node.x;
                unit.y = firstNode.path_node.y;
                CreateUnit(unit, true);

                unit.path = patchNodes;                // добавляем юниту путь
            }

            CheckPath(unit);
            break;
        }
    }

    if (patchNodes[patchNodes.length - 1].path_node.type !== "hide" &&
        patchNodes[patchNodes.length - 1].path_node.type !== "inToFog") {
        MarkLastPathCell(unit, patchNodes[patchNodes.length - 1].path_node);        // помечаем ячейку куда идет моб
    }
}

function CheckPath(unit) {
    let pathNode = unit.path.shift();   // берем первый пункт назначения

    if (pathNode.path_node.type === "hide") {
        StopUnit(unit);

        if (unit.path.length === 0) {
            UnitDestroy(unit);
        }

        return
    }

    if (pathNode.path_node.type === "inToFog") {
        if (unit.path.length === 0) {
            UnitDestroy(unit);
        }
    }

    if (unit.path.length === 0) {
        DeleteMarkLastPathCell(unit.lastCell); // удаляем метку
        unit.lastCell = null;
    }

    if (unit.sprite) {
        unit.rotate = pathNode.unit_rotate;
        unit.watch = pathNode.watch_node;
        unit.movePoint = pathNode.path_node;
    }
}

function MoveToCell(unit) {
    unit.sprite.body.velocity = game.physics.arcade.velocityFromAngle(unit.spriteAngle, 100); // устанавливаем скорость
}

function StopUnit(unit) {
    unit.sprite.body.angularVelocity = 0;
    unit.sprite.body.velocity.x = 0;
    unit.sprite.body.velocity.y = 0;
}

function HideUnit(unit) {
    game.add.tween(unit.sprite).to({alpha: 0}, 1000, Phaser.Easing.Linear.None, true);
    game.add.tween(unit.sprite.unitBody).to({alpha: 0}, 1000, Phaser.Easing.Linear.None, true);
    game.add.tween(unit.sprite.unitShadow).to({alpha: 0}, 1000, Phaser.Easing.Linear.None, true);
    game.add.tween(unit.sprite.healBar).to({alpha: 0}, 1000, Phaser.Easing.Linear.None, true);
    game.add.tween(unit.sprite.heal).to({alpha: 0}, 1000, Phaser.Easing.Linear.None, true);
    // todo если нет спрайта ошибка game.add.tween(unit.sprite.shield).to({alpha: 0}, 1000, Phaser.Easing.Linear.None, true);
    //TODO тут надо сделать `for in unit.sprite` но мне чето лень :D
}

function UncoverUnit(unit) {
    game.add.tween(unit.sprite).to({alpha: 1}, 1000, Phaser.Easing.Linear.None, true);
    game.add.tween(unit.sprite.unitBody).to({alpha: 1}, 1000, Phaser.Easing.Linear.None, true);
    game.add.tween(unit.sprite.unitShadow).to({alpha: 1}, 1000, Phaser.Easing.Linear.None, true);
    game.add.tween(unit.sprite.healBar).to({alpha: 1}, 1000, Phaser.Easing.Linear.None, true);
    game.add.tween(unit.sprite.heal).to({alpha: 1}, 1000, Phaser.Easing.Linear.None, true);
    // todo если нет спрайта ошибкаgame.add.tween(unit.sprite.shield).to({alpha: 1}, 1000, Phaser.Easing.Linear.None, true);
}

function MoveUnit() {
    for (let x in game.units) {
        if (game.units.hasOwnProperty(x)) {
            for (let y in game.units[x]) {
                if (game.units[x].hasOwnProperty(y) && game.units[x][y].sprite) {
                    let unit = game.units[x][y];

                    if (unit.movePoint == null) { // если у юнита больше нет цели перемещения выставляем ему скорость движения и поворота 0
                        StopUnit(unit);
                    } else {

                        if (unit.spriteAngle === unit.rotate) {
                            MoveToCell(unit);
                        } else {
                            StopUnit(unit);
                        }

                        let dist = game.physics.arcade.distanceToXY(unit.sprite, unit.movePoint.x * 100 + 50 , unit.movePoint.y * 100 + 50);

                        if (Math.round(dist) >= -5 && Math.round(dist) <= 5) { // если юнит стоит рядом с целью в приемлемом диапазоне то считаем что он достиг цели

                            delete game.units[unit.x][unit.y];

                            unit.x = unit.movePoint.x;
                            unit.y = unit.movePoint.y;

                            addToGameUnit(unit);

                            if (unit.movePoint.type === "inToFog") {
                                HideUnit(unit);
                            }

                            if (unit.movePoint.type === "outFog") {
                                UncoverUnit(unit);
                            }

                            unit.movePoint = null;
                            UpdateWatchZone(unit.watch);

                            if (unit.path.length > 0) {
                                StopUnit(unit);
                                CheckPath(unit);
                            } else {
                                if (unit.action) {
                                    DeactivationUnit(unit);
                                }
                            }
                        }
                    }
                }
            }
        }
    }
}