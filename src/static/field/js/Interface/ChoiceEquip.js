function ChoiceEquip() {
    var inventory = document.getElementById("inventory");

    if (!inventory) {
        inventory = document.createElement("div");
        inventory.id = "inventory";

        var table = CreateTableInventory();
        inventory.appendChild(table);
        var cancelButton = CreateCancelButton();
        inventory.appendChild(cancelButton);

    } else {
        inventory.style.display = "inline-block";
    }

    inventory.style.top = stylePositionParams.top - 90 + 'px';
    inventory.style.left = 70 + stylePositionParams.left + 'px';

    document.body.appendChild(inventory);
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
            //todo заполнение итемами + onclick + mouseOver со вспылающей подсказкой
            rowInventory.appendChild(cellInventory);
        }
        table.appendChild(rowInventory);
    }

    return table
}

function CreateCancelButton() {
    var cancelButton = document.createElement("input");
    cancelButton.type = "button";
    cancelButton.value = "Отмена";
    cancelButton.className = "button subMenu";

    cancelButton.onclick = function () {
        if (document.getElementById("inventory")) {
            document.getElementById("inventory").remove()
        }
    };

    return cancelButton;
}