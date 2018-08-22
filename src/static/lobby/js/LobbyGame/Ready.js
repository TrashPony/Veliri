function sendReady(gameName) {

    let selectResp = document.getElementById("RespawnSelect");

    let respID = "";
    if (selectResp) {
        respID = selectResp.value;
    }

    lobby.send(JSON.stringify({
        event: "Ready",
        game_name: gameName,
        respawn_id: Number(respID)
    }));
}

function Ready(jsonMessage) {
    let error = JSON.parse(jsonMessage).error;
    let ownedName = JSON.parse(jsonMessage).user_name;
    let respawn = JSON.parse(jsonMessage).respawn;

    let user = Object();
    user.Name = JSON.parse(jsonMessage).game_user;
    user.Ready = JSON.parse(jsonMessage).ready;

    if (error === "") {
        let userRespawnCell = document.getElementById(user.Name).cells[1];
        let userReadyCell = document.getElementById(user.Name).cells[2];

        if (!user.Ready) {

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
            userRespawnCell.innerHTML = user.Respawn.id;
        }
        Respawn();
    } else {
        alert(error)
    }
}