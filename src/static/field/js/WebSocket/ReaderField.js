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
        SelectTargetCoordinateCreate(jsonMessage);
    }

    if (event === "UpdateWatchMap") {
        console.log(jsonMessage);
        let watch = JSON.parse(jsonMessage).update;
        UpdateWatchZone(watch);
    }

    if (event === "PlaceUnit") {
        PlaceUnit(jsonMessage);
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
        UpdateUnit(jsonMessage);
    }

    if (event === "UseUnitEquip") {
        AnimateUseUnitEquip(jsonMessage);
    }

    if (event === "UseMapEquip") {
        AnimateUseMapEquip(jsonMessage);
    }

    if (event === "Error") {
        console.log(jsonMessage);
    }
}