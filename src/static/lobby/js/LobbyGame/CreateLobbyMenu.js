function CreateLobbyMenu(id, error, hoster) {
    if (error === "") {

        DelElements("NotGameLobby");

        var gameInfo = document.createElement('table');
        gameInfo.width = "400px";
        gameInfo.className = "table";
        gameInfo.id = "gameInfo";

        var parentElem = document.getElementById("lobby");
        parentElem.appendChild(gameInfo);

        CreateChoiceSquadBlock(parentElem);

        var br = document.createElement("p");
        parentElem.appendChild(br);

        var cancel = document.createElement("input");
        cancel.type = "button";
        cancel.className = "lobbyButton";
        cancel.value = "Отменить";
        cancel.onclick = ReturnLobby;
        parentElem.appendChild(cancel);

        var ready = document.createElement("input");
        ready.type = "button";
        ready.style.marginLeft = "10px";
        ready.style.left = "95px";
        ready.className = "lobbyButton";
        ready.value = "Готов";

        ready.id = id + ":ready";
        ready.onclick = function () {
            var readyID = this.id.split(":");
            sendReady(readyID[0])
        };

        parentElem.appendChild(ready);

        if (hoster) {
            var button = document.createElement("input");
            button.type = "button";
            button.style.right = "10px";
            button.style.marginLeft = "120px";
            button.className = "lobbyButton";
            button.value = "Начать";
            button.onclick = CreateNewGame;
            parentElem.appendChild(button);
        }

        createGame = true;

        var tr = document.createElement('tr');
        gameInfo.appendChild(tr);

        var th1 = document.createElement('th');
        th1.className = "h";
        th1.appendChild(document.createTextNode("Игроки"));

        var th2 = document.createElement('th');
        th2.className = "h";
        th2.appendChild(document.createTextNode("Готовность"));

        var th3 = document.createElement('th');
        th3.className = "h";
        th3.appendChild(document.createTextNode("Респаун"));

        tr.appendChild(th1);
        tr.appendChild(th3);
        tr.appendChild(th2);

    } else {
        if (error === "lobby is full") {
            alert("Игра полная")
        }
        if (error === "unknown error") {
            alert("unknown error")
        }
    }
}