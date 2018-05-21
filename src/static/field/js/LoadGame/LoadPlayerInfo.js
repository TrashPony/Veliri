function InitPlayer() {
    var ready = document.getElementById("Ready");

    if (game.user.ready) {
        ready.value = "Ты готов!";
        ready.className = "button noActive";
        ready.onclick = null
    } else {
        ready.value = "Готов!";
        ready.className = "button";
        ready.onclick = function () {
            Ready();
        };
    }
}