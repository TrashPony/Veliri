function LoadHoldUnits() {
    var holdTable = document.getElementById("holdUnits");
    var tr;

    if (game.unitStorage !== null) {
        var count = 0;
        for (var unit in game.unitStorage) {

            if (count === 0) {
                tr = document.createElement("tr");
                tr.className = "UnitRow";
            }
            if (game.unitStorage.hasOwnProperty(unit)) {
                var td = document.createElement("td");
                var boxUnit = document.createElement("div");
                boxUnit.className = "boxUnit";

                boxUnit.unit = {};
                boxUnit.unit.info = game.unitStorage[unit];
                boxUnit.id = game.unitStorage[unit].id;

                boxUnit.onclick = function () {
                    field.send(JSON.stringify({
                        event: "SelectStorageUnit",
                        unit_id: Number(this.unit.info.id)
                    }));
                };

                boxUnit.onmouseover = function () {
                    unitTip(this.unit);
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
    }
}