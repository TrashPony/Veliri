function CreatePathToUnit(jsonMessage) {

    let error = JSON.parse(jsonMessage).error;
    if (error === null || error === "") {

        let path = JSON.parse(jsonMessage).path;                           // берем масив данных очереди перемещения юнита
        let unitStat = JSON.parse(jsonMessage).unit;

        let unit = GetGameUnitID(unitStat.id);         // берем юнита

        if (!unit) {

            let boxUnit = document.getElementById(JSON.parse(jsonMessage).unit.id);
            if (boxUnit) {
                boxUnit.remove();
            }

            unitStat.q = path[0].path_node.q;
            unitStat.r = path[0].path_node.r;
            unitStat.rotate = path[0].unit_rotate;

            unit = CreateLocalUnit(unitStat, false)         // создаем юнита
        }

        unit.action_points = unitStat.action_points;

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
                // если юнита не существует даем ему координаты первой видимой клетки
                unit = JSON.parse(jsonMessage).unit;
                unit.q = firstNode.path_node.q;
                unit.r = firstNode.path_node.r;
                unit.rotate = firstNode.unit_rotate;

                CreateLocalUnit(unit, true);
                unit.path = patchNodes;                // добавляем юниту путь
                UncoverUnit(unit);                     // откраем юнита в игре
            }

            CheckPath(unit);
            break
        }
    }

    if (patchNodes.length > 0 && patchNodes[patchNodes.length - 1].path_node.type !== "hide" &&
        patchNodes[patchNodes.length - 1].path_node.type !== "inToFog") {
        MarkLastPathCell(unit, patchNodes[patchNodes.length - 1].path_node);        // помечаем ячейку куда идет моб
    }
}

function CheckPath(unit) {
    let pathNode = unit.path.shift();   // берем первый пункт назначения

    if (pathNode.to_mc) {
        UnitHide(unit);
        DeleteMarkLastPathCell(unit.lastCell);
        UpdateWatchZone(pathNode.watch_node);
        unit.lastCell = null;
        return
    }

    if (pathNode.path_node.type === "hide") {
        StopUnit(unit);

        if (unit.path.length === 0) {
            UnitHide(unit);
        }

        return
    }

    if (pathNode.path_node.type === "inToFog") {
        if (unit.path.length === 0) {
            UnitHide(unit);
        }
    }

    if (unit.sprite) {
        unit.rotate = pathNode.unit_rotate; // тут приходят фактическое положение на карте
        unit.watch = pathNode.watch_node;
        unit.movePoint = pathNode.path_node;
    }

    if (unit.path.length === 0) {
        DeleteMarkLastPathCell(unit.lastCell); // удаляем метку куда идет моб
        unit.lastCell = null;
    }
}

function MoveToCell(unit, x, y) {
    game.physics.arcade.moveToXY(unit.sprite, x, y, 100);
}

function StopUnit(unit) {
    unit.sprite.body.angularVelocity = 0;
    unit.sprite.body.velocity.x = 0;
    unit.sprite.body.velocity.y = 0;
}

function HideUnit(unit) {
    game.add.tween(unit.sprite).to({alpha: 0}, 1000, Phaser.Easing.Linear.None, true);
}

function UncoverUnit(unit) {
    game.add.tween(unit.sprite).to({alpha: 1}, 1000, Phaser.Easing.Linear.None, true);
    game.add.tween(unit.sprite.heal).to({alpha: 1}, 1000, Phaser.Easing.Linear.None, true);
}

function MoveUnit() {
    for (let q in game.units) {
        if (game.units.hasOwnProperty(q)) {
            for (let r in game.units[q]) {
                if (game.units[q].hasOwnProperty(r) && game.units[q][r].sprite) {
                    let unit = game.units[q][r];

                    if (unit.movePoint == null) { // если у юнита больше нет цели перемещения выставляем ему скорость движения и поворота 0
                        StopUnit(unit);
                    } else {
                        let xy = GetXYCenterHex(unit.movePoint.q, unit.movePoint.r);

                        let dist = game.physics.arcade.distanceToXY(unit.sprite, xy.x, xy.y);

                        if (Math.round(dist) >= -10 && Math.round(dist) <= 10) { // если юнит стоит рядом с целью в приемлемом диапазоне то считаем что он достиг цели
                            reachPoint(unit);
                        } else {

                            let spriteRotate = Math.round(unit.sprite.angle);
                            let needRotate = Math.round(Math.atan2(xy.y - unit.sprite.y, xy.x - unit.sprite.x) * 180 / Math.PI);

                            let diffRotate;
                            if (spriteRotate > needRotate) {
                                diffRotate = spriteRotate - needRotate
                            } else {
                                diffRotate = needRotate - spriteRotate
                            }

                            if (diffRotate < 0) {
                                diffRotate += 360;
                            }
                            if (diffRotate > 180) diffRotate -= 180;

                            // кстыль
                            if (needRotate === 180 && spriteRotate === -180) {
                                needRotate = -180
                            }
                            // кстыль

                            if (needRotate === spriteRotate) {
                                MoveToCell(unit, xy.x, xy.y);
                            } else {
                                SetAngle(unit, needRotate, diffRotate * 30, false);
                            }
                        }
                    }
                }
            }
        }
    }
}

function reachPoint(unit) {

    delete game.units[unit.q][unit.r];

    unit.q = unit.movePoint.q;
    unit.r = unit.movePoint.r;

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