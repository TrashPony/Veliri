function ReadResponse(jsonMessage) {
    var event = JSON.parse(jsonMessage).event;

    if (event === "LoadGame") {
        LoadGame(jsonMessage)
    }

    if (event === "SelectStorageUnit") {
        SelectCoordinateUnitCreate(jsonMessage)
    }

    if (event === "SelectMoveUnit") {
        SelectMoveCoordinateCreate(jsonMessage)
    }

    if (event === "GetTargets") {
        SelectTargetCoordinateCreate(jsonMessage);
    }

    if (event === "UpdateWatchMap") {
        var watch = JSON.parse(jsonMessage).update;
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
        InitMoveUnit(jsonMessage);
    }


    /*
       if (event === "InitUnit") {   // +
           InitUnit(jsonMessage);
       }

       if (event === "InitStructure") { // +
           InitStructure(jsonMessage);
       }

       if (event === "PlaceUnit") {
           PlaceUnit(jsonMessage);
       }

       if (event === "SelectUnit") {  // + -
           SelectUnit(jsonMessage);
       }

       if (event === "Attack") {
           AttackUnit(jsonMessage);
       }

       if (event === "SelectCoordinateCreate") {
           SelectCoordinateUnitCreate(jsonMessage)
       }

       if (event === "OpenCoordinate") { // +
           OpenCoordinate(jsonMessage)
       }

       if (event === "DellCoordinate") { // +
           CloseCoordinate(jsonMessage)
       }

       if (event === "MouseOver") { // +
           ReadInfoMouseOver(jsonMessage);
       }

       if (event === "TargetUnit") {
           TargetUnit();
       }*/
}