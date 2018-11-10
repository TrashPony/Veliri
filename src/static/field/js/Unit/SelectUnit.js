function SelectUnit() {

    RemoveSelect();

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

    if (game.Phase === "targeting") {
        field.send(JSON.stringify({
            event: "SelectWeapon",
            q: Number(this.q),
            r: Number(this.r)
        }));
    }

    field.send(JSON.stringify({
        event: "SelectUnit",
        q: Number(this.q),
        r: Number(this.r)
    }));
}