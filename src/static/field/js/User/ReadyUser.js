function ReadyUser(jsonMessage) {
    if (JSON.parse(jsonMessage).error === null || JSON.parse(jsonMessage).error === undefined) {
        var ready = document.getElementById("Ready");

        if (JSON.parse(jsonMessage).ready) {
            ready.value = "Ты готов!";
            ready.className = "button noActive";
            ready.onclick = null;
        } else {
            ready.value = "Готов!";
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
    RemoveSelect();

    field.send(JSON.stringify({
        event: "Ready"
    }));
}