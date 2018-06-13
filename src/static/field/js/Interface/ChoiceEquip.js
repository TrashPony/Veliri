function ChoiceEquip() {

    var inventory = document.getElementById("inventory");

    if (inventory) {
        inventory.remove();
    }

    inventory = document.createElement("div");
    inventory.id = "inventory";

    var table = CreateTableInventory();
    inventory.appendChild(table);

    document.getElementById("inventoryBlock").appendChild(inventory);

    FillingCellInventory();
}

function CreateTableInventory() {
    var table = document.createElement("table");
    table.id = "TableInventory";
    var trHead = document.createElement("tr");
    trHead.height = "10px";
    var thHead = document.createElement("th");
    thHead.colSpan = 4;
    thHead.className = "h";
    thHead.innerHTML = "Инвертарь";
    thHead.appendChild(trHead);
    table.appendChild(thHead);

    for (var i = 0; i < 3; i++) {
        var rowInventory = document.createElement("tr");
        rowInventory.className = "rowInventory";
        for (var j = 0; j < 4; j++) {
            var cellInventory = document.createElement("td");
            cellInventory.className = "cellInventory";
            rowInventory.appendChild(cellInventory);
        }
        table.appendChild(rowInventory);
    }

    return table
}

function FillingCellInventory() {
    var cells = document.getElementsByClassName('cellInventory');

    // todo будут проблемы если эквипом больше чем ячеек
    for (var i = 0; i < game.user.equip.length; i++) {
        cells[i].equip = game.user.equip[i];
        cells[i].style.backgroundImage = "url(/assets/" + cells[i].equip.type + ".png)";

        //cells[i].onclick = function () {
        //    UsedEquip(this.equip);
        //};
        cells[i].onmouseover = function () {
            TipEquipOn(this.equip);
        };
        cells[i].onmouseout = function () {
            TipEquipOff();
        };
    }
}