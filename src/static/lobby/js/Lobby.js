var createGame = false;
var createNameGame = "";
var toField = false;
var respownId;

function ReturnLobby() {
    location.reload();
}

function CreateLobbyGame(mapId) {
    var gameName = document.querySelector('input[name="NameGame"]').value;
    sendCreateLobbyGame(mapId, gameName);
}

function CreateNewGame() {
    if(createNameGame !== "") {
        sendStartNewGame(createNameGame);
    } else {
        location.href = "../../login";
    }
}

function JoinToGame(idGame) {
    toField = true;
    document.cookie = "idGame=" + idGame + "; path=/;";
    location.href = "http://" + window.location.host + "/field";
}

function MapSelection() {

    DelElements("Select SubMenu");

    var SubMenu = document.getElementById("SubMenu");
    var inputs = SubMenu.getElementsByTagName("th");
    var tableH = document.getElementById("NotEndGame");

    tableH.innerHTML = "Выберети карту";

    var i = inputs.length;
    while (i--) {
        var input = inputs[i];
        if (input) {
            var th = input.parentNode.parentNode;
            SubMenu.deleteRow(th.rowIndex);
        }
    }

    var tr = document.createElement('tr');
    SubMenu.appendChild(tr);

    var thName = document.createElement('th');
    thName.innerHTML = "Название карты";
    var thPlayers =  document.createElement('th');
    thPlayers.innerHTML = "Максимум игроков";
    tr.appendChild(thName);
    tr.appendChild(thPlayers);

    lobby.send(JSON.stringify({
        event: "MapView"
    }));

    var div = document.getElementById("cancel");
    var cancel = document.getElementById("cancelButton");

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
    var SelectMap = document.getElementsByClassName(ClassElements);
    while (SelectMap.length > 0) {
        SelectMap[0].parentNode.removeChild(SelectMap[0]);
    }
}