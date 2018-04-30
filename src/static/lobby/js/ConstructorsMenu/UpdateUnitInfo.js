function UpdateUnitInfo(unit) {


    for (var parameter in unit) {
        if(unit.hasOwnProperty(parameter)) {
            var row = document.getElementById(parameter);
            if (row) {
                row.innerHTML = unit[parameter];
            }
        }
    }
}