function ReadResponse(jsonMessage) {
    let event = JSON.parse(jsonMessage).event;

    if (event === "LoadGame") {
        Game(jsonMessage)
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
        ChangePhaseNotification(jsonMessage);
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
        console.log(jsonMessage);
        AttackPhase(jsonMessage)
    }

    if (event === "QueueMove") {
        MoveNotification(jsonMessage)
    }

    if (event === "UpdateMemoryUnit") {
        UpdateMemoryUnit(jsonMessage)
    }

    if (event === "Error") {
        console.log(jsonMessage);
    }

    jsonMessage = null;
}