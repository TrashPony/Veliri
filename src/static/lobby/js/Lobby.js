let createGame = false;
let toField = false;
let respownId;

function ReturnLobby() {
    location.reload();
}

function CreateLobbyGame(mapId) {
    let gameName = document.querySelector('input[name="NameGame"]').value;
    sendCreateLobbyGame(mapId, gameName);
}

function JoinToGame(idGame) {
    toField = true;
    document.cookie = "idGame=" + idGame + "; path=/;";
    location.href = "http://" + window.location.host + "/field";
}

function MapSelection() {

    DelElements("Select SubMenu");

    let SubMenu = document.getElementById("SubMenu");
    let inputs = SubMenu.getElementsByTagName("th");
    let tableH = document.getElementById("NotEndGame");

    tableH.innerHTML = "Выберети карту";

    let i = inputs.length;
    while (i--) {
        let input = inputs[i];
        if (input) {
            let th = input.parentNode.parentNode;
            SubMenu.deleteRow(th.rowIndex);
        }
    }

    let tr = document.createElement('tr');
    SubMenu.appendChild(tr);

    let thName = document.createElement('th');
    thName.innerHTML = "Название карты";
    let thPlayers =  document.createElement('th');
    thPlayers.innerHTML = "Максимум игроков";
    tr.appendChild(thName);
    tr.appendChild(thPlayers);

    lobby.send(JSON.stringify({
        event: "MapView"
    }));

    let div = document.getElementById("cancel");
    let cancel = document.getElementById("cancelButton");

    if (!cancel) {
        cancel = document.createElement("input");
        cancel.type = "button";
        cancel.value = "Отменить";
        cancel.className = "button";
        cancel.id = "cancelButton";
        cancel.onclick = ReturnLobby;
        div.appendChild(cancel);
    }
}

function DelElements(ClassElements) {
    let SelectMap = document.getElementsByClassName(ClassElements);
    while (SelectMap.length > 0) {
        SelectMap[0].parentNode.removeChild(SelectMap[0]);
    }
}