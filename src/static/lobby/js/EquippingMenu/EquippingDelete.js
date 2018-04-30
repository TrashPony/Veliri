function DeleteEquip(box, slot) {

    var mask = document.getElementById("mask");
    mask.style.display = "block";

    var lobbyMenu = document.getElementById("lobby");

    var DellEquip = document.getElementById("DellEquip");

    if (!DellEquip) {
        DellEquip = document.createElement("div");
        DellEquip.id = "DellEquip";
        lobbyMenu.appendChild(DellEquip);

        var headSpan = document.createElement("span");
        headSpan.className = "value";
        headSpan.innerHTML = "Удалить модуль?";

        var acceptButton = document.createElement("button");
        acceptButton.className = "button";
        acceptButton.style.margin = "5px";
        acceptButton.innerHTML = "Принять";

        acceptButton.onclick = function () {

            lobby.send(JSON.stringify({
                event: "RemoveEquipment",
                equip_slot: Number(slot)
            }));

            var mask = document.getElementById("mask");
            mask.style.display = "none";

            var DellEquip = document.getElementById("DellEquip");
            DellEquip.style.display = "none";
        };

        var cancelButton = document.createElement("button");
        cancelButton.className = "button";
        cancelButton.style.margin = "5px";
        cancelButton.innerHTML = "Отмена";
        cancelButton.onclick = function () {
            var mask = document.getElementById("mask");
            mask.style.display = "none";

            var DellEquip = document.getElementById("DellEquip");
            DellEquip.style.display = "none";
        };

        DellEquip.appendChild(headSpan);
        DellEquip.appendChild(cancelButton);
        DellEquip.appendChild(acceptButton);
    } else {
        DellEquip.style.display = "block";
    }
}