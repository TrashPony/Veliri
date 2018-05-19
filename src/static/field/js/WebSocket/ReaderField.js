function ReadResponse(jsonMessage) {
    var event = JSON.parse(jsonMessage).event;

    //console.log(jsonMessage);

    if (event === "LoadGame") {
        LoadGame(jsonMessage)
    }

    if (event === "SelectStorageUnit") {
        SelectCoordinateUnitCreate(jsonMessage)
    }

    /*if (event === "InitPlayer") { // +
        console.log(jsonMessage);
        InitPlayer(jsonMessage);
    }

    if (event === "InitMap") {    // +
        FieldCreate(jsonMessage);
    }

    if (event === "InitUnit") {   // +
        InitUnit(jsonMessage);
    }

    if (event === "InitStructure") { // +
        InitStructure(jsonMessage);
    }

    if (event === "PlaceUnit") {
        PlaceUnit(jsonMessage);
    }

    if (event === "Ready") {      // +
        ReadyReader(jsonMessage);
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

    if (event === "MoveUnit") { // +
        InitMoveUnit(jsonMessage);
    }

    if (event === "TargetUnit") {
        TargetUnit();
    }*/
}