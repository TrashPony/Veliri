function SelectUnit(unitStat, focus) {

    let testWeapon = {};
    testWeapon.type = "laser";
    testWeapon.artillery = false;
    testWeapon.name = "big_laser";

    // Fire(unitStat, game.map.OneLayerMap[1][2], "coordinate").then(
    //     function () {
    //         Fire(unitStat, game.map.OneLayerMap[1][59], "coordinate").then(
    //             function () {
    //                 Fire(unitStat, game.map.OneLayerMap[59][1], "coordinate").then(
    //                     function () {
    //                         Fire(unitStat, game.map.OneLayerMap[59][59], "coordinate");
    //                     }
    //                 );
    //             }
    //         );
    //     }
    // );

    if (focus) {
        game.camera.focusOnXY(unitStat.sprite.x * game.camera.scale.x, unitStat.sprite.y * game.camera.scale.y);
    }

    if (!unitStat.body.mother_ship && unitStat.owner === game.user.name) {
        field.send(JSON.stringify({
            event: "GetAmmoZone"
        }));
    }

    if (game.Phase === "targeting" && unitStat.owner === game.user.name) {
        field.send(JSON.stringify({
            event: "SelectWeapon",
            q: Number(unitStat.q),
            r: Number(unitStat.r)
        }));
        removeUnitInput();
        RemoveSelect(true, true);
    } else {
        RemoveSelect(false);
    }

    field.send(JSON.stringify({
        event: "SelectUnit",
        q: Number(unitStat.q),
        r: Number(unitStat.r)
    }));

    CreateUnitSubMenu(unitStat);
    MarkUnitSelect(unitStat, 1);
}

function MarkUnitSelect(unit, frame, onclickFunc) {
    unit.sprite.frame = frame;

    if (onclickFunc) {
        unit.sprite.events.onInputDown.add(onclickFunc);
        unit.sprite.input.priorityID = 1;
    }
}

function RemoveUnitMarks() {
    for (let x in game.units) {
        if (game.units.hasOwnProperty(x)) {
            for (let y in game.units[x]) {
                if (game.units[x].hasOwnProperty(y) && game.units[x][y].sprite) {
                    let unit = game.units[x][y];
                    unit.sprite.frame = 0;
                    unit.sprite.events.onInputDown.removeAll();
                    unit.sprite.input.priorityID = 0;
                }
            }
        }
    }
}