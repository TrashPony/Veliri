function NoActiveCell() {
    let cells = document.getElementsByClassName("inventoryEquipping");

    for (let i = 0; i < cells.length; i++) {
        cells[i].ammoCell = null;

        cells[i].className = "inventoryEquipping noActive";
        cells[i].style.backgroundImage = "";
        cells[i].style.boxShadow = "0 0 0 0 rgb(0, 0, 0)";

        cells[i].onmouseout = null;
        cells[i].onmouseover = null;
        cells[i].onclick = null;

        for (let child in cells[i].childNodes) {
            if (cells[i].childNodes.hasOwnProperty(child)) {
                cells[i].childNodes[child].remove();
            }
        }
    }

    let unitIcon = document.getElementById("MSIcon");
    unitIcon.style.backgroundImage = null;
    unitIcon.onclick = null;
    unitIcon.shipBody = null;

    let powerPanel = document.getElementById("powerPanel");
    powerPanel.innerHTML = "<span class='Value'> Энергия: <br>" + 0 + "/" + 0 + "</span>";
}

function NoActiveUnitCell(slotData) {
    let constructorUnit = document.getElementById("ConstructorUnit");
    if (constructorUnit) {
        if (JSON.parse(constructorUnit.slotData).number_slot === slotData.number_slot) {
            let cells = document.getElementsByClassName("UnitEquip");
            for (let i = 0; i < cells.length; i++) {
                cells[i].ammoCell = null;

                cells[i].className = "UnitEquip noActive";
                cells[i].style.backgroundImage = "";
                cells[i].style.boxShadow = "0 0 0 0 rgb(0, 0, 0)";

                cells[i].onmouseout = null;
                cells[i].onmouseover = null;
                cells[i].onclick = null;

                for (let child in cells[i].childNodes) {
                    if (cells[i].childNodes.hasOwnProperty(child)) {
                        cells[i].childNodes[child].remove();
                    }
                }
            }

            let unitIcon = document.getElementById("UnitIcon");
            unitIcon.style.backgroundImage = null;
            unitIcon.onclick = null;
            unitIcon.shipBody = null;
        }

        let powerPanel = document.getElementById("unitPowerPanel");
        powerPanel.innerHTML = "<span class='Value'>" + 0 + "/" + 0 + "</span>";
    }
}