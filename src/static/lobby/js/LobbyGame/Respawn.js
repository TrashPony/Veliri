function RespawnInit(jsonMessage) {

    let respawns = JSON.parse(jsonMessage).respawns;
    let select = document.getElementById("RespawnSelect");

    if (select) {
        for (let id in respawns) {
            if (respawns.hasOwnProperty(id)) {
                let option = document.createElement("option");
                option.className = "RespawnOption";
                option.value = respawns[id].id;
                option.text = respawns[id].id;
                option.id = respawns[id].id + "respawn";
                select.appendChild(option);
            }
        }
    }
}

function CreateSelectRespawn(id) {
    let user = document.getElementById(id).cells[1];
    let selectList = document.createElement("select");
    selectList.id = "RespawnSelect";
    selectList.className = "RespawnSelect";
    user.appendChild(selectList);
}