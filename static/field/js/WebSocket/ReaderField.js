function ReadResponse(jsonMessage) {
    let event = JSON.parse(jsonMessage).event;

    if (event === "LoadGame") {
        LoadGame(jsonMessage)
    }

    if (event === "SelectStorageUnit") {
        SelectCoordinateUnitCreate(jsonMessage)
    }

    if (event === "SelectMoveUnit") {
        SelectMoveCoordinateCreate(jsonMessage)
    }

    if (event === "GetTargets" || event === "GetFirstTargets") {
        SelectTargetCoordinateCreate(jsonMessage, SelectWeaponTarget);
    }

    if (event === "GetEquipMapTargets") {
        MarkEquipSelect(jsonMessage);
    }

    if (event === "GetEquipMyUnitTargets" || event === "GetEquipMyUnitTargets" ||
        event === "GetEquipMySelfTarget" || event === "GetEquipAllUnitTarget") {
        SelectTargetUnit(jsonMessage);
    }

    if (event === "UpdateWatchMap") {
        let watch = JSON.parse(jsonMessage).update;
        UpdateWatchZone(watch);
    }

    if (event === "Ready") {
        ReadyUser(jsonMessage);
    }

    if (event === "ChangePhase") {
        ChangePhase(jsonMessage);
    }

    if (event === "MoveUnit") {
        CreatePathToUnit(jsonMessage);
    }

    if (event === "HostileUnitMove") {
        MoveHostileUnit(jsonMessage)
    }

    if (event === "UpdateUnit") {
        let unitStat = JSON.parse(jsonMessage).unit;
        UpdateUnit(unitStat);
    }

    if (event === "PreviewPath") {
        CreatePreviewPath(jsonMessage)
    }

    if (event === "AttackPhase") {
        AttackPhase(jsonMessage)
    }

    if (event === "QueueMove") {
        MoveNotification(jsonMessage)
    }

    if (event === "UpdateMemoryUnit") {
        UpdateMemoryUnit(jsonMessage)
    }

    if (event === "UpdateGameZone") {
        MarkGameZone(JSON.parse(jsonMessage).game_zone);
    }

    if (event === 'leave') {
        LeaveBattle(false);
    }

    if (event === 'softLeave') {
        LeaveBattle(true);
    }

    if (event === 'initReload') {
        initReload(JSON.parse(jsonMessage))
    }

    if (event === 'Reload') {
        ReloadMark(JSON.parse(jsonMessage));
    }

    if (event === 'timeToLeave') {
        LeaveTimer(JSON.parse(jsonMessage).seconds)
    }

    if (event === 'timeToChangePhase') {
        document.getElementById("stepTime").innerHTML = JSON.parse(jsonMessage).seconds;
    }

    if (event === 'OpenDiplomacy') {
        CreateDiplomacyMenu(JSON.parse(jsonMessage))
    }

    if (event === 'UpdateUnitStorage') {
        game.unitStorage = JSON.parse(jsonMessage).unit_storage;
        LoadHoldUnits();
    }

    if (event === 'toGlobal') {
        location.href = "../../../global";
    }

    if (event === "Error") {
        console.log(jsonMessage);
    }

    jsonMessage = null;
}