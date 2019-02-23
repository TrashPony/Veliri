function LoadHoldUnits() {
    let holdTable = document.getElementById("holdUnits");
    let tr;

    if (game.unitStorage !== null && game.unitStorage.length !== 0) {
        let count = 0;
        for (let unit in game.unitStorage) {

            if (count === 0) {
                tr = document.createElement("tr");
                tr.className = "UnitRow";
            }
            if (game.unitStorage.hasOwnProperty(unit)) {
                let td = document.createElement("td");
                let boxUnit = document.createElement("div");
                boxUnit.className = "boxUnit";

                boxUnit.unit = {};
                boxUnit.unit.info = game.unitStorage[unit];
                boxUnit.id = game.unitStorage[unit].id;
                boxUnit.style.backgroundImage = "url(/assets/units/body/" + game.unitStorage[unit].body.name + ".png)";

                let weapon = document.createElement("div");
                weapon.className = "weaponHoldUnit";
                for (let i in game.unitStorage[unit].body.weapons) {
                    if (game.unitStorage[unit].body.weapons.hasOwnProperty(i) && game.unitStorage[unit].body.weapons[i].weapon) {
                        weapon.style.backgroundImage = "url(/assets/units/weapon/" + game.unitStorage[unit].body.weapons[i].weapon.name + ".png)";
                    }
                }
                boxUnit.appendChild(weapon);

                boxUnit.onclick = function () {

                    CreateUnitSubMenu(game.unitStorage[unit]);

                    field.send(JSON.stringify({
                        event: "SelectStorageUnit",
                        unit_id: Number(this.unit.info.id)
                    }));

                    if (document.getElementsByClassName("boxUnit").length === 0) { // todo не работает :(
                        document.getElementById("holdUnits").style.display = "none";
                    }
                };

                boxUnit.onmouseover = function () {
                    unitTip(game.unitStorage[unit]);
                };

                boxUnit.onmouseout = function () {
                    TipOff()
                };

                td.appendChild(boxUnit);
                tr.appendChild(td);
                holdTable.appendChild(tr);
            }

            count++;
            if (count === 3) {
                count = 0
            }
        }
    } else {
        holdTable.style.display = "none";
    }
}