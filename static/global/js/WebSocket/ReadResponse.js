function ReadResponse(jsonData) {
    if (jsonData.event === "InitGame") {
        Game(jsonData);
    }

    if (jsonData.event === "Error") {
        Notification(jsonData.error);
    }

    if (jsonData.event === "PreviewPath") {
        PreviewPath(jsonData);
    }

    if (jsonData.event === "MoveTo") {
        MoveTo(jsonData);
    }

    if (jsonData.event === "WorkOutThorium") {
        ThoriumBar(jsonData.thorium_slots)
    }

    if (jsonData.event === "DisconnectUser") {
        DisconnectUser(jsonData);
    }

    if (jsonData.event === "ConnectNewUser") {
        CreateNewUnit(jsonData.short_unit);
    }

    if (jsonData.event === "openBox") {
        OpenBox(jsonData.inventory, jsonData.box_id, jsonData.size, jsonData.error)
    }

    if (jsonData.event === "startMoveEvacuation") {
        if (game.bases) startMoveEvacuation(jsonData)
    }

    if (jsonData.event === "FreeMoveEvacuation") {
        FreeMoveEvacuation(jsonData)
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

    if (jsonData.event === "AfterburnerToggle") {
        Afterburner(jsonData.afterburner)
    }

    if (jsonData.event === "ChangeGravity") {
        ChangeGravity(jsonData.squad)
    }

    if (jsonData.event === "DestroyBox") {
        DestroyBox(jsonData.box_id)
    }

    if (jsonData.event === "startMining") {
        StartMining(jsonData);
    }

    if (jsonData.event === "destroyReservoir") {
        DestroyReservoir(jsonData);
    }

    if (jsonData.event === "updateReservoir") {
        UpdateReservoir(jsonData)
    }

    if (jsonData.event === "stopMining") {
        StopMining(jsonData)
    }

    if (jsonData.event === "setFreeCoordinate") {
        Alert("Освободите выход с базы. <br> Иначе будете отбуксированы!", "Внимание!", false, jsonData.seconds, true, "setFreeResp");
    }

    if (jsonData.event === "softTransition") {
        Alert("Перемещениече через: ", "Внимание!", false, 5, false, "softTransition");
    }

    if (jsonData.event === "removeSoftTransition") {
        if (document.getElementById("softTransition")) {
            document.getElementById("softTransition").remove();
        }
    }

    if (jsonData.event === "removeNoticeFreeCoordinate") {
        if (document.getElementById("setFreeResp")) document.getElementById("setFreeResp").remove();
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

    if (jsonData.event === "AnomalySignal") {
        VisibleAnomalies(jsonData.anomalies)
    }

    if (jsonData.event === "RemoveAnomalies") {
        RemoveOldAnomaly();
    }

    if (jsonData.event === "SelectDigger") {
        SelectDigger(jsonData.coordinates, jsonData.slot, jsonData.type_slot);
    }

    if (jsonData.event === "useDigger") {
        UseDigger(jsonData);
    }

    if (jsonData.event === "IntoToBase") {
        location.href = "http://" + window.location.host + "/lobby";
    }

    if (jsonData.event === "changeSector") {
        location.href = "http://" + window.location.host + "/global";
    }

    if (jsonData.event === "MoveCloud") {
        CreateCloud(jsonData)
    }

    if (jsonData.event === "handlerClose") {
        CloseTunnel(jsonData)
    }

    if (jsonData.event === "handlerOpen") {
        OpenTunnel(jsonData)
    }

    if (jsonData.event === "AnomalyCatch") {
        AnomalyCatch(jsonData)
    }

    if (jsonData.event === "DamageUnit") {
        FillSquadBlock(jsonData.squad)
    }

    if (jsonData.event === "FillSquadBlock") {
        FillSquadBlock(jsonData.squad)
    }

    if (jsonData.event === "DeadSquad") {
        SquadDead(jsonData.other_user);
    }

    if (jsonData.event === "GetMissions") {
        FillMissionsSelect(jsonData.missions, jsonData.mission_uuid)
    }

    if (jsonData.event === "GetPortalPointToGlobalPath") {
        let {x, y} = GetXYCenterHex(jsonData.q, jsonData.r);
        if (game && game.squad && jsonData.name === "mission") game.squad.missionMove = {x: x, y: y, radius: 0}
    }

    if (jsonData.event === "InitMiningOre") {
        InitMiningOre(jsonData.short_unit.id, jsonData.slot, jsonData.type_slot, jsonData.equip)
    }
}