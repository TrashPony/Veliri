function RespawnInit(jsonMessage) {

    var respawn = JSON.parse(jsonMessage).respawn;
    var select = document.getElementById("RespawnSelect");

    if (select) {
        var option = document.createElement("option");
        option.className = "RespawnOption";
        option.value = respawn.Name;
        option.text = respawn.Name;
        option.id = respawn.Id;
        select.appendChild(option);
    }
}

function CreateSelectRespawn(id) {
    var user = document.getElementById(id).cells[1];
    var selectList = document.createElement("select");
    selectList.id = "RespawnSelect";
    selectList.className = "RespawnSelect";
    user.appendChild(selectList);
}