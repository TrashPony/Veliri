function ConstructorTable(shipBody) {

    UpdateCells(1, "inventoryEquip", shipBody.equippingI, "inventoryEquipping");
    UpdateCells(2, "inventoryEquip", shipBody.equippingII, "inventoryEquipping");
    UpdateCells(3, "inventoryEquip", shipBody.equippingIII, "inventoryEquipping");
    UpdateCells(5, "inventoryEquip", shipBody.equippingV, "inventoryEquipping");

    UpdateCells(3, "inventoryEquip", shipBody.weapons, "inventoryEquipping");
    /* вепоны надо делать отдельно т.к. храняться отдельно*/

    UpdateShipIcon(shipBody)
}

function UpdateShipIcon(shipBody) {
    let unitIcon = document.getElementById("MSIcon");
    unitIcon.shipBody = unitIcon;
    unitIcon.style.backgroundImage = "url(/assets/" + shipBody.name + ".png)";
    unitIcon.onclick = BodyMSMenu;
}