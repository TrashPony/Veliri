function FillingInventory(jsonData) {
    let event = JSON.parse(jsonData).event;
    console.log(jsonData);
    if (event === "openInventory" || event === "UpdateSquad") {
        let squad = JSON.parse(jsonData).squad;
        InventoryTable(squad.inventory);
        SquadTable(squad);
        if (squad.mather_ship.body != null) {
            ConstructorTable(squad.mather_ship.body);
        } else {
            NoActiveCell();
        }
    }
}

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
}