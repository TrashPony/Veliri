function DeleteSquad() {
    var mask = document.getElementById("mask");
    mask.style.display = "block";

    var lobbyMenu = document.getElementById("lobby");

    var DellSquad = document.getElementById("DellSquad");

    if (!DellSquad) {
        DellSquad = document.createElement("div");
        DellSquad.id = "DellSquad";
        lobbyMenu.appendChild(DellSquad);

        var headSpan = document.createElement("span");
        headSpan.className = "value";
        headSpan.innerHTML = "Удалить отряд?";

        var acceptButton = document.createElement("button");
        acceptButton.className = "button";
        acceptButton.style.margin = "5px";
        acceptButton.innerHTML = "Принять";

        acceptButton.onclick = function () {

            lobby.send(JSON.stringify({
                event: "DeleteSquad"
            }));

            var mask = document.getElementById("mask");
            mask.style.display = "none";

            var DellSquad = document.getElementById("DellSquad");
            DellSquad.style.display = "none";
        };

        var cancelButton = document.createElement("button");
        cancelButton.className = "button";
        cancelButton.style.margin = "5px";
        cancelButton.innerHTML = "Отмена";
        cancelButton.onclick = function () {
            var mask = document.getElementById("mask");
            mask.style.display = "none";

            var DellSquad = document.getElementById("DellSquad");
            DellSquad.style.display = "none";
        };

        DellSquad.appendChild(headSpan);
        DellSquad.appendChild(cancelButton);
        DellSquad.appendChild(acceptButton);
    } else {
        DellSquad.style.display = "block";
    }
}

function RemoveSquad(jsonMessage) {

    var squadID = JSON.parse(jsonMessage).squad_id;
    var squadOption = document.getElementById(squadID + ":squad");
    squadOption.remove();

    DeleteInfoSquad();

    var selectSquad = document.getElementById("listSquad");
    selectSquad.value = "Не выбрано";
}