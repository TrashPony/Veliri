function ReadyUser(jsonMessage) {
    if (JSON.parse(jsonMessage).error === null || JSON.parse(jsonMessage).error === undefined) {
        game.user.ready = JSON.parse(jsonMessage).ready;
        InitPlayer()
    } else {
        alert(JSON.parse(jsonMessage).error)
    }
}

function Ready() {
    let passMove = false;
    let passTarget = false;

    for (let q in game.units) {
        for (let r in game.units[q]) {
            if (game.units[q][r].action_points > 0 && game.units[q][r].move && game.units[q][r].owner === game.user.name) {
                passMove = true;
            }
            if (!game.units[q][r].target && game.units[q][r].owner === game.user.name) {
                passTarget = true;
            }
        }
    }

    if (passMove && game.Phase === "move") {
        ViewAlert("У вас еще остались не использованые очки движения.");
    } else if (passTarget && game.Phase === "targeting") {
        ViewAlert("Не у всех юнитов есть цель.");
    } else {
        RemoveSelect();
        field.send(JSON.stringify({
            event: "Ready"
        }));
    }
}

function ViewAlert(text) {

    let page = {
        text: text,
        picture: "base.png",
        asc: [],
    };

    let dialogBlock = CreatePageDialog("LeaveBlock", page, null, false, true);
    dialogBlock.style.right = "calc(50% - 125px)";
    dialogBlock.style.top = "calc(50% - 300px)";
    dialogBlock.style.bottom = "unset";
    dialogBlock.style.left = "unset";

    let ask = document.createElement("div");
    ask.className = "asks";
    ask.innerHTML = "<div class='wrapperAsk'>Завершить фазу</div>";
    ask.onclick = function () {
        RemoveSelect();
        field.send(JSON.stringify({
            event: "Ready"
        }));
        dialogBlock.remove();
    };

    let ask2 = document.createElement("div");
    ask2.className = "asks";
    ask2.innerHTML = "<div class='wrapperAsk'>Отмена</div>";
    ask2.onclick = function () {
        dialogBlock.remove();
    };

    dialogBlock.appendChild(ask);
    dialogBlock.appendChild(ask2);
}