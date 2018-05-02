function UnitMouseOver(unit) {
    if (unit !== undefined) {
        var tipUnit = document.getElementById("tipUnit");
        tipUnit.style.display = "block";
        for (var parameter in unit) {
            if (unit.hasOwnProperty(parameter)) {
                var row = document.getElementById(parameter + ":tip");
                if (row) {
                    row.innerHTML = unit[parameter];
                }
            }
        }
    }
}