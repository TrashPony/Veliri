function AddNewSquad() {
    var mask = document.getElementById("mask");
    mask.style.display = "block";

    var lobbyMenu = document.getElementById("lobby");

    var ChoiceNameSquad = document.getElementById("ChoiceNameSquad");

    if (!ChoiceNameSquad) {
        ChoiceNameSquad = document.createElement("div");
        ChoiceNameSquad.id = "ChoiceNameSquad";
        lobbyMenu.appendChild(ChoiceNameSquad);

        var headSpan = document.createElement("span");
        headSpan.className = "value";
        headSpan.innerHTML = "Назовите отряд";

        var inputName = document.createElement("input");

        var acceptButton = document.createElement("button");
        acceptButton.className = "button";
        acceptButton.style.margin = "5px";
        acceptButton.innerHTML = "Принять";
        acceptButton.onclick = function () {
            lobby.send(JSON.stringify({ // запрашиваем список имеющийся отрядов
                event: "AddNewSquad",
                squad_name: inputName.value
            }));

            inputName.value = "";

            var mask = document.getElementById("mask");
            mask.style.display = "none";

            var ChoiceNameSquad = document.getElementById("ChoiceNameSquad");
            ChoiceNameSquad.style.display = "none";

        };

        var cancelButton = document.createElement("button");
        cancelButton.className = "button";
        cancelButton.style.margin = "5px";
        cancelButton.innerHTML = "Отмена";
        cancelButton.onclick = function () {
            inputName.value = "";

            var mask = document.getElementById("mask");
            mask.style.display = "none";

            var ChoiceNameSquad = document.getElementById("ChoiceNameSquad");
            ChoiceNameSquad.style.display = "none";
        };

        ChoiceNameSquad.appendChild(headSpan);
        ChoiceNameSquad.appendChild(inputName);
        ChoiceNameSquad.appendChild(cancelButton);
        ChoiceNameSquad.appendChild(acceptButton);
    } else {
        ChoiceNameSquad.style.display = "block";
    }
}