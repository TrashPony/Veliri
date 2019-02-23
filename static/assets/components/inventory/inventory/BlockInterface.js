function RemoveActionConstructorMenu() { // отключает ивенты меню с возможностью востановления

    if (document.getElementById("ConstructorUnit")) document.getElementById("ConstructorUnit").remove();
    if (document.getElementById("repairMenu")) document.getElementById("repairMenu").remove();

    if (document.getElementById("repairMenu")) {
        document.getElementById("repairButton").onclick = null;
        document.getElementById("repairButton").className = "repairButton";
    }

    let unitCells = document.getElementsByClassName("inventoryUnit");

    for (let i = 0; i < unitCells.length; i++) {

        if (unitCells[i].className === "inventoryUnit select") {
            unitCells[i].className = "inventoryUnit active";
        }

        unitCells[i].onmouseoutBack = unitCells[i].onmouseout;
        unitCells[i].onmouseoverBack = unitCells[i].onmouseover;
        unitCells[i].onclickBack = unitCells[i].onclick;
        // :D
        unitCells[i].onmouseout = null;
        unitCells[i].onmouseover = null;
        unitCells[i].onclick = null;
    }

    let equippingCells = document.getElementsByClassName("inventoryEquipping");

    for (let i = 0; i < equippingCells.length; i++) {
        equippingCells[i].onmouseoutBack = equippingCells[i].onmouseout;
        equippingCells[i].onmouseoverBack = equippingCells[i].onmouseover;
        equippingCells[i].onclickBack = equippingCells[i].onclick;

        equippingCells[i].onmouseout = null;
        equippingCells[i].onmouseover = null;
        equippingCells[i].onclick = null;
    }

    let unitIcon = document.getElementById("MSIcon");
    if (unitIcon) {
        unitIcon.onclickBack = unitIcon.onclick;
        unitIcon.onmousemoveBack = unitIcon.onmousemove;
        unitIcon.onmouseoutBack = unitIcon.onmouseout;

        unitIcon.onclick = null;
        unitIcon.onmousemove = null;
        unitIcon.onmouseout = null;
    }
}

function ActionConstructorMenu() {

    if (document.getElementById("repairMenu")) document.getElementById("repairButton").onclick = CreateRepairMenu;

    let unitCells = document.getElementsByClassName("inventoryUnit");

    for (let i = 0; i < unitCells.length; i++) {
        unitCells[i].onmouseout = unitCells[i].onmouseoutBack;
        unitCells[i].onmouseover = unitCells[i].onmouseoverBack;
        unitCells[i].onclick = unitCells[i].onclickBack;

        unitCells[i].onmouseoutBack = null;
        unitCells[i].onmouseoverBack = null;
        unitCells[i].onclickBack = null;
    }

    let equippingCells = document.getElementsByClassName("inventoryEquipping");

    for (let i = 0; i < equippingCells.length; i++) {
        equippingCells[i].onmouseout = equippingCells[i].onmouseoutBack;
        equippingCells[i].onmouseover = equippingCells[i].onmouseoverBack;
        equippingCells[i].onclick = equippingCells[i].onclickBack;

        equippingCells[i].onmouseoutBack = null;
        equippingCells[i].onmouseoverBack = null;
        equippingCells[i].onclickBack = null;
    }

    let unitIcon = document.getElementById("MSIcon");
    if (unitIcon) {
        unitIcon.onclick = unitIcon.onclickBack;
        unitIcon.onmousemove = unitIcon.onmousemoveBack;
        unitIcon.onmouseout = unitIcon.onmouseoutBack;

        unitIcon.onclickBack = null;
        unitIcon.onmousemoveBack = null;
        unitIcon.onmouseoutBack = null;
    }
}