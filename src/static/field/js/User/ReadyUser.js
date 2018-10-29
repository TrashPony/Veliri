function ReadyUser(jsonMessage) {
    if (JSON.parse(jsonMessage).error === null || JSON.parse(jsonMessage).error === undefined) {
        var ready = document.getElementById("Ready");

        if (JSON.parse(jsonMessage).ready) {
            ready.value = "Ты готов!";
            ready.className = "button noActive";
            ready.onclick = null;
        } else {
            ready.value = "Завершить ход";
            ready.className = "button";
            ready.onclick = function () {
                Ready();
            };
        }
    } else {
        alert(JSON.parse(jsonMessage).error)
    }
}

function Ready() {
    // TODO смотреть есть еще очки действий у отряда или нет
    let confirmReady = document.createElement("div");
    confirmReady.id = "confirmReady";

    let head = document.createElement("h3");
    head.innerHTML = "Завершить ход?";
    confirmReady.appendChild(head);

    let text = document.createElement("p");
    text.innerHTML = "Вы уверены что хотите завершить ход? У вас еще остались не использованые очки движения.";
    confirmReady.appendChild(text);

    let cancel = document.createElement("input");
    cancel.type = "submit";
    cancel.value = "Отмена";
    cancel.onclick = function () {
        confirmReady.remove();
    };
    confirmReady.appendChild(cancel);

    let button = document.createElement("input");
    button.type = "submit";
    button.value = "OK";
    button.onclick = function () {

        RemoveSelect();
        field.send(JSON.stringify({
            event: "Ready"
        }));

        confirmReady.remove();
    };
    confirmReady.appendChild(button);

    document.body.appendChild(confirmReady)
}