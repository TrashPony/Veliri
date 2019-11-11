function ReadResponse(jsonData) {

    if (jsonData.event === "InitGame") {
        Game(jsonData);
        return;
    }

    let awaitReady = function (jsonData) {
        if (gameReady) {
            ReadResponse(jsonData)
        } else {
            setTimeout(() => awaitReady(jsonData), 50);
        }
    };

    if (!gameReady) {
        // игра еще не создалась, что бы не пропустить сообщения пусть они ждут пока игра не поднимется
        awaitReady(jsonData);
        return
    }

    if (jsonData.event === "RefreshRadar") {
        RemoveAllMark();
        removeAllObj();

        // TODO окончание загрузки
    }

    if (jsonData.event === "focusMS") {
        FocusUnit(jsonData.short_unit.id);
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

    if (jsonData.event === "RotateGun") {
        if (!game || !game.units) return;
        let unit = game.units[jsonData.short_unit.id];
        let path = jsonData.path_unit;

        RotateGun(unit, path.rotate_gun, path.millisecond);
    }

    if (jsonData.event === "FireWeapon") {
        // TODO отыгрываем анимаю выстрела weapon(много ифов ибо каждое оружие стрелять особенно ебано) в позиции х,у
    }

    if (jsonData.event === "FlyBullet") {
        FlyBullet(jsonData);
    }

    if (jsonData.event === "ExplosionBullet") {
        // TODO проигрываем взрыв в точке снаряда, (удаляем снаряд)
        //   появляется кратер который прилетает с бека, если не прилетает то нет
    }

    if (jsonData.event === "BoxTo") {
        BoxMove(jsonData.path_unit, jsonData.box_id)
    }

    if (jsonData.event === "MoveStop") {
        MoveStop(jsonData)
    }

    if (jsonData.event === "WorkOutThorium") {
        ThoriumBar(jsonData.unit, jsonData.thorium_slots)
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
        ChangeGravity(jsonData.high_gravity)
    }

    if (jsonData.event === "DestroyBox") {
        DestroyBox(jsonData.box_id, true)
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

    if (jsonData.event === "InitDigger") {
        SelectDigger(jsonData.short_unit.id, jsonData.slot, jsonData.type_slot, jsonData.equip)
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

    if (jsonData.event === "PlaceUnit") {
        CreateNewUnit(jsonData.short_unit)
    }

    if (jsonData.event === "RemoveUnit") {
        RemoveUnit(jsonData)
    }

    if (jsonData.event === "CreateRect") {
        FindPathDebug(jsonData)
    }

    if (jsonData.event === "CreateLine") {
        FindPathDebug(jsonData)
    }

    if (jsonData.event === "ClearPath") {
        findPath.clear();
    }

    if (jsonData.event === "CreatePolygon") {
        PolygonDraw(jsonData.polygon)
    }

    if (jsonData.event === "ClearPolygon") {
        if (polygonCanvas) polygonCanvas.clear();
    }

    if (jsonData.event === "NewFormationPos") {
        Data.squad = jsonData.squad;
        fillFormation(Data.squad, scaleFormation);
    }

    if (jsonData.event === "radarWork") {
        RadarWork(jsonData)
    }

    if (jsonData.event === "markMove") {
        MoveMark(jsonData)
    }
}

let moveDebug = null;
let findPath = null;
let polygonCanvas = null;

function PolygonDraw(polygon) {
    if (!game) return;

    if (!polygonCanvas) {
        setTimeout(function () {
            polygonCanvas = game.add.graphics(0, 0);
        }, 2500)
    }

    polygonCanvas.lineStyle(1, 0xFF0000, 0.8);
    for (let i in polygon.sides) {

        let side = polygon.sides[i];

        polygonCanvas.moveTo(side.x_1, side.y_1);
        polygonCanvas.lineTo(side.x_2, side.y_2);

        polygonCanvas.endFill();
    }
    /*
    squad.rectDebag.moveTo(bodyRec.sides[0].x1, bodyRec.sides[0].y1);
    squad.rectDebag.lineTo(bodyRec.sides[0].x2, bodyRec.sides[0].y2);

    squad.rectDebag.lineTo(bodyRec.sides[1].x1, bodyRec.sides[1].y1);
    squad.rectDebag.lineTo(bodyRec.sides[1].x2, bodyRec.sides[1].y2);

    squad.rectDebag.lineTo(bodyRec.sides[2].x1, bodyRec.sides[2].y1);
    squad.rectDebag.lineTo(bodyRec.sides[2].x2, bodyRec.sides[2].y2);

    squad.rectDebag.lineTo(bodyRec.sides[3].x1, bodyRec.sides[3].y1);
    squad.rectDebag.lineTo(bodyRec.sides[3].x2, bodyRec.sides[3].y2);
    squad.rectDebag.endFill();
    */
}

function FindPathDebug(jsonData) {
    if (!game) return;

    if (!moveDebug) {
        setTimeout(function () {
            moveDebug = game.add.graphics(0, 0);
            findPath = game.add.graphics(0, 0);
        }, 2500)
    }

    let color = 0xFFFFFF;
    if (jsonData.color === "green") color = 0x00FF00;
    if (jsonData.color === "red") color = 0xFF0000;
    if (jsonData.color === "blue") color = 0x0000FF;
    if (jsonData.color === "orange") color = 0xff9200;
    if (jsonData.color === "black") color = 0x000000;

    moveDebug.lineStyle(1, color, 1);
    findPath.lineStyle(1, color, 0.5);

    if (jsonData.event === "CreateRect") {
        if (jsonData.color !== "white") {
            moveDebug.drawRect(jsonData.x, jsonData.y, jsonData.rect_size, jsonData.rect_size);
        } else {
            findPath.drawRect(jsonData.x, jsonData.y, jsonData.rect_size, jsonData.rect_size);
        }
    }

    if (jsonData.event === "CreateLine") {
        moveDebug.moveTo(jsonData.x, jsonData.y);
        moveDebug.lineTo(jsonData.to_x, jsonData.to_y);
    }
}