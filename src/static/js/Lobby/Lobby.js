var createGame = false;
var createNameGame = "";
var toField = false;
var sock;
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

    sock.send(JSON.stringify({
        event: "MapView"
    }));

    var div = document.getElementById("cancel");
    var cancel = document.createElement("input");
    cancel.type = "button";
    cancel.value = "Отменить";
    cancel.className = "button";
    cancel.onclick = ReturnLobby;
    div.appendChild(cancel);

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

function NewChatMessage(jsonMessage) {
    var chatBox = document.getElementById("chatBox");
    var UserName = document.createElement("span");
    UserName.className = "ChatUserName";
    UserName.innerHTML = JSON.parse(jsonMessage).game_user + ":";
    var TextMessage = document.createElement("span");
    TextMessage.className = "ChatText";
    TextMessage.innerHTML = JSON.parse(jsonMessage).message;
    chatBox.appendChild(UserName);
    chatBox.appendChild(TextMessage);
    chatBox.appendChild(document.createElement("br"));
}