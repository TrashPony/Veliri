function NewLobbyUser(jsonMessage) {
    var UsersInLobbyTable = document.getElementById("UsersInLobby");
    var NewUser = JSON.parse(jsonMessage).game_user;

    var userTr  = document.getElementById(NewUser + "allList");

    if (!userTr) {
        var tr = document.createElement("tr");
        tr.id = NewUser + "allList";
        tr.style.textAlign = "center";
        var td = document.createElement("td");
        td.className = "Value";
        td.innerHTML = NewUser;
        tr.appendChild(td);
        UsersInLobbyTable.appendChild(tr);
    }
}

function DelLobbyUser(jsonMessage) {
    var DelUser = JSON.parse(jsonMessage).game_user;
    var userTr  = document.getElementById(DelUser + "allList");
    if (userTr) {
        userTr.parentNode.removeChild(userTr);
    }
}