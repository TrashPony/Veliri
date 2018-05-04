function RespawnInit(jsonMessage) {

    var respawns = JSON.parse(jsonMessage).respawns;
    var select = document.getElementById("RespawnSelect");

    if (select) {
        for (var i = 0; i < respawns.length; i++) {
            if (respawns[i].UserName === "") {
                var option = document.createElement("option");
                option.className = "RespawnOption";
                option.value = respawns[i].Name;
                option.text = respawns[i].Name;
                option.id = respawns[i].Id;
                select.appendChild(option);
            }
        }
    }
}

function CreateSelectRespawn(id) {
    var user = document.getElementById(id).cells[1];
    var selectList = document.createElement("select");
    selectList.id = "RespawnSelect";
    selectList.className = "RespawnSelect";
    user.appendChild(selectList);
}