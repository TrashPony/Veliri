function SelectUnit(unitStat, focus) {
    // let testWeapon = {};
    // testWeapon.type = "laser";
    // testWeapon.artillery = false;
    // testWeapon.name = "big_laser";
    //
    // OutFogFire(game.map.OneLayerMap[10][2], game.map.OneLayerMap[1][2], testWeapon, "coordinate").then(function () {
    //    console.log("dfdfd")
    // });
    //
    // let unit = this;
    //
    // Fire(unit, game.map.OneLayerMap[1][2], "coordinate").then(
    //     function () {
    //         Fire(unit, game.map.OneLayerMap[1][9], "coordinate").then(
    //             function () {
    //                 Fire(unit, game.map.OneLayerMap[10][2], "coordinate").then(
    //                     function () {
    //                         Fire(unit, game.map.OneLayerMap[10][9], "coordinate");
    //                     }
    //                 );
    //             }
    //         );
    //     }
    // );

    //Fire(unitStat, GetGameUnitID(31));
    if (focus) {
        game.camera.focusOnXY(unitStat.sprite.x * game.camera.scale.x, unitStat.sprite.y * game.camera.scale.y);
    }

    if(!unitStat.body.mother_ship && unitStat.owner === game.user.name) {
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