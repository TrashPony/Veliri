function FillingInventory(jsonData) {
    let event = JSON.parse(jsonData).event;

    if (event === "openInventory" || event === "UpdateSquad") {
        console.log(jsonData);
        let squad = JSON.parse(jsonData).squad;
        FillingInventoryTable(squad.inventory);

        if (squad.mather_ship.body != null) {
            FillingConstructorTable(squad.mather_ship.body)
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
    }

    console.log(cells);

    let unitIcon = document.getElementById("UnitIcon");
    unitIcon.style.backgroundImage = null;
    unitIcon.onclick = null;
    unitIcon.shipBody = null;

    let inventoryAmmoCell = document.getElementsByClassName("inventoryAmmoCell");
    for (let i = 0; i < inventoryAmmoCell.length; i++) {
        inventoryAmmoCell[i].remove();
    }
}



function FillingSquadTable() {

}

