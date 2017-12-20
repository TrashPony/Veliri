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


    if (list && gameContent === "NotEndGame") {

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

    if (list && gameContent === "Game") {
        tdName.appendChild(document.createTextNode(text.Name));
        tdID.appendChild(document.createTextNode(text.Map));
        tdStep.appendChild(document.createTextNode(text.Creator));
        tdPhase.appendChild(document.createTextNode(text.Players));
        tdMyStep.appendChild(document.createTextNode(text.Copasity));

        tr.appendChild(tdName);
        tr.appendChild(tdID);
        tr.appendChild(tdStep);
        tr.appendChild(tdPhase);
        tr.appendChild(tdMyStep);

        list.appendChild(tr);
    }

    if (list && gameContent === "Map") {
        tdName.appendChild(document.createTextNode(text.Name));
        tdPhase.appendChild(document.createTextNode(text.Copasity));

        tr.appendChild(tdName);
        tr.appendChild(tdPhase);

        list.appendChild(tr);
    }

    if (list && gameContent === "User") {

        tdName.appendChild(document.createTextNode(text.Name));
        tdStep.appendChild(document.createTextNode(text.Respawn));
        tdPhase.appendChild(document.createTextNode(text.Ready));

        if (text.Ready !== " Готов") {
            tdPhase.className = "Failed";
        } else {
            tdPhase.className = "Success";
        }

        tr.appendChild(tdName);
        tr.appendChild(tdStep);
        tr.appendChild(tdPhase);

        list.appendChild(tr);

        if (id === owned && text.Respawn === "") {
            CreateSelectRespawn(id);
            Respawn();
        }
    }
}

function CreateLobbyMenu(id, error, hoster) {
    if (error === "") {

        DelElements("NotGameLobby");

        var gameInfo = document.createElement('table');
        gameInfo.width = "400px";
        gameInfo.className = "table";
        gameInfo.id = "gameInfo";

        var parentElem = document.getElementById("lobby");
        parentElem.appendChild(gameInfo);

        var br = document.createElement("p");
        parentElem.appendChild(br);

        var cancel = document.createElement("input");
        cancel.type = "button";
        cancel.className = "button";
        cancel.value = "Отменить";
        cancel.onclick = ReturnLobby;
        parentElem.appendChild(cancel);

        var ready = document.createElement("input");
        ready.type = "button";
        ready.style.marginLeft = "10px";
        ready.className = "button";
        ready.value = "Готов";
        ready.onclick = function () {
            sendReady(this.id)
        };

        ready.id = id;
        parentElem.appendChild(ready);

        if (hoster) {
            var button = document.createElement("input");
            button.type = "button";
            button.style.marginLeft = "120px";
            button.className = "button";
            button.value = "Начать";
            button.onclick = CreateNewGame;
            button.id = id;
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
            console.log("Игра полная");
            alert("Игра полная")
        }
        if (error === "unknown error") {
            console.log("unknown error");
            alert("unknown error")
        }
    }
}

function CreateSelectRespawn(id) {
    var user = document.getElementById(id).cells[1];
    var selectList = document.createElement("select");
    selectList.id = "RespawnSelect";
    selectList.className = "RespawnSelect";
    user.appendChild(selectList);
}