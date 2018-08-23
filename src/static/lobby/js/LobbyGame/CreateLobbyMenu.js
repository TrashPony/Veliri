function CreateLobbyMenu(id, error, hoster) {
    if (error === "") {

        DelElements("NotGameLobby");

        let gameInfo = document.createElement('table');
        gameInfo.width = "400px";
        gameInfo.className = "table";
        gameInfo.id = "gameInfo";

        let parentElem = document.getElementById("lobby");
        parentElem.appendChild(gameInfo);

        CreateSquadBlock(parentElem);

        let br = document.createElement("p");
        parentElem.appendChild(br);

        let cancel = document.createElement("input");
        cancel.type = "button";
        cancel.className = "lobbyButton";
        cancel.value = "Отменить";
        cancel.onclick = ReturnLobby;
        parentElem.appendChild(cancel);

        let inventory = document.createElement("input");
        inventory.type = "button";
        inventory.style.marginLeft = "0px";
        inventory.style.left = "95px";
        inventory.className = "lobbyButton";
        inventory.value = "Инвентарь";
        inventory.onclick = () => {
            InitInventoryMenu(
                () => {
                    lobby.send(
                        JSON.stringify({
                            event: "GetSquad"
                        })
                    )
                }
            );
        };
        parentElem.appendChild(inventory);

        let ready = document.createElement("input");
        ready.type = "button";
        ready.style.marginLeft = "95px";
        ready.style.left = "95px";
        ready.className = "lobbyButton";
        ready.value = "Готов";

        ready.id = id + ":ready";
        ready.onclick = function () {
            let readyID = this.id.split(":");
            sendReady(readyID[0]);
        };

        parentElem.appendChild(ready);

        if (hoster) {
            let button = document.createElement("input");
            button.type = "button";
            button.style.right = "10px";
            button.style.marginLeft = "120px";
            button.className = "lobbyButton";
            button.value = "Начать";
            button.onclick = sendStartNewGame;
            parentElem.appendChild(button);
        }

        createGame = true;

        let tr = document.createElement('tr');
        gameInfo.appendChild(tr);

        let th1 = document.createElement('th');
        th1.className = "h";
        th1.appendChild(document.createTextNode("Игроки"));

        let th2 = document.createElement('th');
        th2.className = "h";
        th2.appendChild(document.createTextNode("Готовность"));

        let th3 = document.createElement('th');
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