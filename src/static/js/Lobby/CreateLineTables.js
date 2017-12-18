function CreateLobbyLine(gameContent, menu, className, id, func, funcMouse, funcOutMouse, text, owned) {
    var list = document.getElementById(menu);
    var tr = document.createElement('tr');

    var tdName = document.createElement('td');
    var tdID = document.createElement('td');
    var tdStep = document.createElement('td');
    var tdPhase = document.createElement('td');
    var tdMyStep = document.createElement('td');

    tr.style.wordWrap = 'break-word';
    tr.className = className;
    tr.align = "center";
    tr.id = id;
    tr.onclick = func;
    tr.onmouseover = funcMouse;
    tr.onmouseout = funcOutMouse;


    if (gameContent === "NotEndGame") {

        tdName.appendChild(document.createTextNode(text.Name));
        tdID.appendChild(document.createTextNode(text.Id));
        tdStep.appendChild(document.createTextNode(text.Step));
        tdPhase.appendChild(document.createTextNode(text.Phase));
        tdMyStep.appendChild(document.createTextNode(text.Ready));

        tr.appendChild(tdName);
        tr.appendChild(tdID);
        tr.appendChild(tdStep);
        tr.appendChild(tdPhase);
        tr.appendChild(tdMyStep);

        list.appendChild(tr);
    }

    if (gameContent === "Game") {
        tdName.appendChild(document.createTextNode(text.Name));
        tdID.appendChild(document.createTextNode(text.Map));
        tdStep.appendChild(document.createTextNode(text.Creator));
        tdPhase.appendChild(document.createTextNode(text.Copasity));
        tdMyStep.appendChild(document.createTextNode(text.Players));

        tr.appendChild(tdName);
        tr.appendChild(tdID);
        tr.appendChild(tdStep);
        tr.appendChild(tdPhase);
        tr.appendChild(tdMyStep);

        list.appendChild(tr);
    }

    if (gameContent === "Map") {
        tdName.appendChild(document.createTextNode(text.Name));
        tdPhase.appendChild(document.createTextNode(text.Copasity));

        tr.appendChild(tdName);
        tr.appendChild(tdPhase);

        list.appendChild(tr);
    }




    if (id === owned) {
        CreateSelectRespawn(id, text)
    }
}

function CreateLobbyMenu(textButton, funcButton, id, error, hoster) {
    if (error === "") {
        DelElements("NotGameLobby");
        var gameInfo = document.createElement('div');
        gameInfo.className = "gameInfo";
        gameInfo.id = "gameInfo";
        var parentElem = document.getElementById("lobby");
        parentElem.appendChild(gameInfo);
        var cancel = document.createElement("input");
        cancel.type = "button";
        cancel.value = "Отменить";
        cancel.onclick = ReturnLobby;
        gameInfo.appendChild(cancel);
        var button = document.createElement("input");
        button.type = "button";
        button.value = textButton;
        button.onclick = funcButton;
        button.id = id;
        gameInfo.appendChild(button);
        if (hoster) {
            var ready = document.createElement("input");
            ready.type = "button";
            ready.value = "Готов";
            ready.onclick = function () {
                sendReady(this.id)
            };
            ready.id = id;
            gameInfo.appendChild(ready);
        }
        createGame = true;
        var parentElemDiv = document.getElementsByClassName("gameInfo");
        var div3 = document.createElement('div');
        div3.className = "User";
        div3.appendChild(document.createTextNode("Подключенные Игроки"));
        parentElemDiv[0].appendChild(div3);
    } else {
        if (error === "lobby is full") {
            console.log("Игра полная");
            alert("Игра полная")
        }
        if (error === "unknown error") {
            console.log("unknown error");
            alert("unknown error")
        }
    }
}

function CreateSelectRespawn(id, msg) {
    var user = document.getElementById(id);
    user.innerHTML = msg + " точка респауна: ";
    var selectList = document.createElement("select");
    selectList.id = "RespawnSelect";
    selectList.className = "RespawnSelect";
    user.appendChild(selectList);
}