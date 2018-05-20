function ReadyReader(jsonMessage) {
    var error = JSON.parse(jsonMessage).error;
    phase = JSON.parse(jsonMessage).phase;

    if (error === "") {
        var ready = document.getElementById("Ready");
        var phaseBlock = document.getElementById("phase");

        if (phase === "") {
            ready.value = "Ты готов!";
            ready.className = "button noActive";
            ready.onclick = null;
        } else {
            ready.value = "Готов!";
            ready.className = "button";
            ready.onclick = function () { Ready(); };

            phaseBlock.innerHTML = JSON.parse(jsonMessage).phase;
            var cells = document.getElementsByClassName("fieldUnit create");
            while (0 < cells.length) {
                if (cells[0]) {
                    cells[0].className = "fieldUnit open";
                }
            }
        }
    } else {
        if (error === "not units") {
            alert("У вас нет юнитов")
        }
    }
}

function Ready(){
    RemoveSelect();

    field.send(JSON.stringify({
        event: "Ready"
    }));
}