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
        if (game.bases) startMoveEvacuation(jsonData)
    }

    if (jsonData.event === "MoveEvacuation") {
        if (game.bases) evacuationMove(jsonData, null)
    }

    if (jsonData.event === "placeEvacuation") {
        if (game.bases) placeEvacuation(jsonData)
    }

    if (jsonData.event === "ReturnEvacuation") {
        if (game.bases) evacuationMove(jsonData, true)
    }

    if (jsonData.event === "stopEvacuation") {
        if (game.bases) stopEvacuation(jsonData)
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