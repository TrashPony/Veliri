function SelectUnit() {


    CreateUnitSubMenu(this);

    MarkUnitSelect(this, 1);


    //let testWeapon = {};
    //testWeapon.type = "laser";
    //testWeapon.artillery = false;
    //testWeapon.name = "big_laser";

    //OutFogFire(game.map.OneLayerMap[10][2], game.map.OneLayerMap[1][2], testWeapon, "coordinate").then(function () {
    //    console.log("dfdfd")
    //});

    /*let unit = this;

    Fire(unit, game.map.OneLayerMap[1][2], "coordinate").then(
        function () {
            Fire(unit, game.map.OneLayerMap[1][9], "coordinate").then(
                function () {
                    Fire(unit, game.map.OneLayerMap[10][2], "coordinate").then(
                        function () {
                            Fire(unit, game.map.OneLayerMap[10][9], "coordinate");
                        }
                    );
                }
            );
        }
    );*/

    //Fire(this, GetGameUnitID(357));

    if (game.Phase === "targeting" && this.owner === game.user.name) {
        field.send(JSON.stringify({
            event: "SelectWeapon",
            q: Number(this.q),
            r: Number(this.r)
        }));
        removeUnitInput();
        RemoveSelect(true);
    } else {
        RemoveSelect(false);
    }

    field.send(JSON.stringify({
        event: "SelectUnit",
        q: Number(this.q),
        r: Number(this.r)
    }));
}