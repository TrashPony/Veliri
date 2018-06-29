function sendReady(gameName) {

    let selectResp = document.getElementById("RespawnSelect");

    if (selectResp) {
        respownId = selectResp.options[selectResp.selectedIndex];
    }

    lobby.send(JSON.stringify({
        event: "Ready",
        game_name: gameName,
        respawn_id: Number(respownId.id)
    }));
}

function Ready(jsonMessage) {
    console.log(jsonMessage);

    let error = JSON.parse(jsonMessage).error;
    let ownedName = JSON.parse(jsonMessage).user_name;
    let respawn = JSON.parse(jsonMessage).respawn;

    let user = Object();
    user.Name = JSON.parse(jsonMessage).game_user;

    if (error === "") {
        let userRespawnCell = document.getElementById(user.Name).cells[1];
        let userReadyCell = document.getElementById(user.Name).cells[2];

        if (user.Ready = JSON.parse(jsonMessage).ready === "false") {

            user.Ready = " Не готов";
            userReadyCell.innerHTML = user.Ready;
            userReadyCell.className = "Failed";
            userRespawnCell.innerHTML = "";

            if (ownedName === user.Name) {
                CreateSelectRespawn(user.Name);
            }

        } else {

            if (ownedName === user.Name) {
                DelElements("RespawnSelect");
            }

            user.Ready = " Готов";
            userReadyCell.innerHTML = user.Ready;
            userReadyCell.className = "Success";

            user.Respawn = respawn;
            userRespawnCell.innerHTML = user.Respawn.Name;
        }
        Respawn();
    } else {
        alert(error)
    }
}