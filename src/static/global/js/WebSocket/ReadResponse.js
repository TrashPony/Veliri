function ReadResponse(jsonData) {
    if (jsonData.event === "InitGame") {
        Game(jsonData);
    }

    if (jsonData.event === "Error") {
        alert(jsonData.error);
    }

    if (jsonData.event === "PreviewPath") {
        PreviewPath(jsonData);
    }

    if (jsonData.event === "MoveTo") {
        MoveTo(jsonData);
    }

    if (jsonData.event === "MoveOtherUser") {
        MoveOther(jsonData);
    }

    if (jsonData.event === "DisconnectUser") {
        DisconnectUser(jsonData);
    }

    if (jsonData.event === "ConnectNewUser") {
        CreateOtherUser(jsonData.other_user);
    }

    if (jsonData.event === "openBox") {
        OpenBox(jsonData.inventory, jsonData.box_id, jsonData.size)
    }

    if (jsonData.event === "startMoveEvacuation"){
        startMoveEvacuation(jsonData)
    }

    if (jsonData.event === "MoveEvacuation") {
        evacuationMove(jsonData, null)
    }

    if (jsonData.event === "placeEvacuation") {
        placeEvacuation(jsonData)
    }

    if (jsonData.event === "ReturnEvacuation") {
        evacuationMove(jsonData, game.squad.sprite)
    }

    if (jsonData.event === "stopEvacuation") {
        stopEvacuation(jsonData)
    }

    if (jsonData.event === "UpdateBox") { // что бы не откывался у тех у кого окно не открыто
        if (document.getElementById("openBox" + jsonData.box_id)) {
            OpenBox(jsonData.inventory, jsonData.box_id, jsonData.size)
        }
    }

    if (jsonData.event === "NewBox") {
        CreateBox(jsonData.box)
    }

    if (jsonData.event === "UpdateInventory") { // говорит клиенту обновлить инвентарь
        if (document.getElementById("Inventory")) {
            if (inventorySocket) inventorySocket.send(JSON.stringify({event: "openInventory"}))
        }
    }

    if (jsonData.event === "IntoToBase") {
        location.href = "http://" + window.location.host + "/lobby";
    }
}