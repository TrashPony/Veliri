function InitPlayer(jsonMessage) {

    var user = JSON.parse(jsonMessage).user;
    User = user;

    if (user.ready) {
        var ready = document.getElementById("Ready");
        ready.value = "Ты готов!";
        ready.className = "button noActive";
        ready.onclick = null
    }
}