function InitEquippingMenu() {
    var mask = document.getElementById("mask");
    mask.style.display = "block";

    var lobbyMenu = document.getElementById("lobby");

    var equippingMenu = document.getElementById("equippingMenu");

    if (!equippingMenu) {
        equippingMenu = document.createElement("div");
        equippingMenu.id = "equippingMenu";
        lobbyMenu.appendChild(equippingMenu);

        var tableEquip = CreateTableEquip();

        equippingMenu.appendChild(tableEquip);

        var acceptButton = document.createElement("input");
        acceptButton.type = "button";
        acceptButton.value = "Back";
        acceptButton.className = "lobbyButton";
        acceptButton.id = "EquipBackButton";
        acceptButton.onclick = EquipBackToLobby;
        equippingMenu.appendChild(acceptButton);

        lobby.send(JSON.stringify({
            event: "GetEquipping"
        }));

    } else {
        equippingMenu.style.display = "block";
    }
}

function CreateTableEquip() {
    var tableEquip = document.createElement("table");
    tableEquip.className = "table";
    tableEquip.id = "tableEquip";

    var headRow = document.createElement("tr");
    var tdHead = document.createElement("td");
    tdHead.innerHTML = "<span class='Value'>Модули</span>";
    tdHead.colSpan = 2;

    headRow.appendChild(tdHead);
    tableEquip.appendChild(headRow);

    return tableEquip;
}

function EquipBackToLobby() {
    var mask = document.getElementById("mask");
    mask.style.display = "none";

    var unitConstructor = document.getElementById("equippingMenu");
    unitConstructor.style.display = "none";
}