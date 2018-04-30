function DeleteUnit(box, slot) {

    var mask = document.getElementById("mask");
    mask.style.display = "block";

    var lobbyMenu = document.getElementById("lobby");

    var DellUnit = document.getElementById("DellUnit");

    if (!DellUnit) {
        DellUnit = document.createElement("div");
        DellUnit.id = "DellUnit";
        lobbyMenu.appendChild(DellUnit);

        var headSpan = document.createElement("span");
        headSpan.className = "value";
        headSpan.innerHTML = "Удалить юнита?";

        var acceptButton = document.createElement("button");
        acceptButton.className = "button";
        acceptButton.style.margin = "5px";
        acceptButton.innerHTML = "Принять";

        acceptButton.onclick = function () {

            lobby.send(JSON.stringify({
                event: "RemoveUnit",
                slot: Number(slot)
            }));

            var mask = document.getElementById("mask");
            mask.style.display = "none";

            var DellUnit = document.getElementById("DellUnit");
            DellUnit.style.display = "none";
        };

        var cancelButton = document.createElement("button");
        cancelButton.className = "button";
        cancelButton.style.margin = "5px";
        cancelButton.innerHTML = "Отмена";
        cancelButton.onclick = function () {
            var mask = document.getElementById("mask");
            mask.style.display = "none";

            var DellUnit = document.getElementById("DellUnit");
            DellUnit.style.display = "none";
        };

        DellUnit.appendChild(headSpan);
        DellUnit.appendChild(cancelButton);
        DellUnit.appendChild(acceptButton);
    } else {
        DellUnit.style.display = "block";
    }
}