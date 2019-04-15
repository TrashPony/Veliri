function MoveNotification(jsonMessage) {
    
    let unitStat = JSON.parse(jsonMessage).unit;

    for (let i in game.unitStorage) {
        if (game.unitStorage.hasOwnProperty(i)) {
            if (unitStat.id === game.unitStorage[i].id) {
                game.unitStorage[i].action_points = unitStat.action_points;
                game.unitStorage[i].move = unitStat.move;
            }
        }
    }

    for (let q in game.units) {
        if (game.units.hasOwnProperty(q)) {
            for (let r in game.units[q]) {
                if (game.units[q].hasOwnProperty(r)) {
                    if (unitStat.id === game.units[q][r].id) {
                        game.units[q][r].action_points = unitStat.action_points;
                        game.units[q][r].move = unitStat.move;
                    }
                }
            }
        }
    }

    if (unitStat.move && game.Phase === "move") {
        notification("Движение", "Твоя очередь двигать юнита")
    }
}

function notification(headText, textText) {
    let queueBlock = document.getElementById("queue");

    let notificationBlock = document.createElement("div");
    notificationBlock.className = "notificationBlock";

    let head = document.createElement("h3");
    head.innerHTML = headText;
    notificationBlock.appendChild(head);

    let text = document.createElement("p");
    text.innerHTML = textText;
    notificationBlock.appendChild(text);

    if (queueBlock) {
        queueBlock.appendChild(notificationBlock);
    }

    let timeNotification = 500;
    setTimeout(function () {
        notificationBlock.style.animation = "notification "+ timeNotification +"ms 1";
    }, 4000);

    setTimeout(function () {
        notificationBlock.style.display = "none";
    }, 4000 + timeNotification)
}