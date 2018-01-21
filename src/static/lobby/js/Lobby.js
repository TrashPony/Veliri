var createGame = false;
var createNameGame = "";
var toField = false;
var respownId;

function ReturnLobby() {
    location.reload();
}

function CreateLobbyGame(mapName) {
    var gameName = document.querySelector('input[name="NameGame"]').value;
    sendCreateLobbyGame(mapName, gameName);
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

function MouseOverMap(id) {
    var info = document.getElementById('SelectInfo');
    info.innerHTML = "Имя карты " + id;
    var div = document.createElement('div');
    div.style.wordWrap = 'break-word';
    div.appendChild(document.createTextNode("Невероятная картинка карты! В разработке"));
    div.className = "infoMap";
    div.id = "infoImage";
    info.appendChild(div);
    var div2 = document.createElement('div');
    div2.style.wordWrap = 'break-word';
    div2.className = "infoMap";
    div2.appendChild(document.createTextNode("Описание карты, в разработке"));
    info.appendChild(div2);
}

function MouseOutMap() {
    var info = document.getElementById('SelectInfo');
    info.innerHTML = "";
    DelElements("infoMap");
}

function DelElements(ClassElements) {
    var SelectMap = document.getElementsByClassName(ClassElements);
    while (SelectMap.length > 0) {
        SelectMap[0].parentNode.removeChild(SelectMap[0]);
    }
}