function Alert(text, headText, okButton, time, alert, id) {

    let notificationWrapper = document.createElement("div");
    notificationWrapper.id = id;
    if (alert){
        notificationWrapper.className = "notificationWrapper alert";
    } else {
        notificationWrapper.className = "notificationWrapper";
    }

    let notificationBlock = document.createElement("div");
    notificationBlock.className = "notificationBlock";

    let head = document.createElement("h3");
    head.innerHTML = headText;
    notificationBlock.appendChild(head);

    let textBlock = document.createElement("p");
    textBlock.innerHTML = text;
    notificationBlock.appendChild(textBlock);

    if (time > 0) {
        let timeBlock = document.createElement("div");
        timeBlock.className = "timeBlock";
        timeBlock.innerHTML = time;
        time--;
        for (let i = 0; i < time; i++){
            setTimeout(function () {
                timeBlock.innerHTML = time;
                time--
            }, i * 1000);

            setTimeout(function () {
                notificationWrapper.remove();
            }, time * 1000)
        }
        notificationBlock.appendChild(timeBlock);
    }

    if (okButton) {
        let button = document.createElement("input");
        button.type = "submit";
        button.value = "OK";
        button.onclick = function () {
            notificationWrapper.remove();
        };
        notificationBlock.appendChild(button);
    }

    notificationWrapper.appendChild(notificationBlock);
    document.body.appendChild(notificationWrapper);
}