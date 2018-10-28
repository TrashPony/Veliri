function CreateMotherShipParamsMenu() {
    let menu = document.getElementById("MotherShipParams");

    let spanInventory = document.createElement("span");
    spanInventory.className = "InventoryHead";
    spanInventory.innerHTML = "ПАРАМЕТРЫ";
    spanInventory.style.width = "110px";
    menu.appendChild(spanInventory);


    let spanAttack = document.createElement("span");
    spanAttack.className = "Value params";
    spanAttack.innerHTML = " ⇓ Атака ";

    let spanDef = document.createElement("span");
    spanDef.className = "Value params";
    spanDef.innerHTML = " ⇓ Защита ";

    let spanNav = document.createElement("span");
    spanNav.className = "Value params";
    spanNav.innerHTML = " ⇓ Навигация ";

    menu.appendChild(spanAttack);
    menu.appendChild(spanDef);
    menu.appendChild(spanNav);
}