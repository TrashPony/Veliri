function ChangePhaseNotification(jsonMessage) {

    let notificationBlock = document.createElement("div");
    notificationBlock.className = "notificationBlock";

    let head = document.createElement("h3");
    head.innerHTML = "Смена фазы";
    notificationBlock.appendChild(head);

    let text = document.createElement("p");
    text.innerHTML = "Начата фаза " + JSON.parse(jsonMessage).game_phase;
    notificationBlock.appendChild(text);

    let button = document.createElement("input");
    button.type = "submit";
    button.value = "OK";
    button.onclick = function () {
        notificationBlock.remove();
    };
    notificationBlock.appendChild(button);

    document.body.appendChild(notificationBlock)
}