function ReadResponse(jsonMessage) {
    var event = JSON.parse(jsonMessage).event;

    if (event === "InitPlayer") { // +
        InitPlayer(jsonMessage);
    }

    if (event === "InitMap") {    // +
        FieldCreate(jsonMessage);
    }

    if (event === "InitUnit") {   // +
        InitUnit(jsonMessage);
    }

    if (event === "InitStructure") {
        InitStructure(jsonMessage);
    }

    if (event === "InitObstacle") { // +
        InitObstacle(jsonMessage);
    }

    if (event === "CreateUnit") {
        CreateUnit(jsonMessage);
    }

    if (event === "Ready") {      // +
        ReadyReader(jsonMessage);
    }

    if (event === "SelectUnit") {  // +-
        setUnitAction(jsonMessage);
    }

    if (event === "Attack") {
        AttackUnit(jsonMessage);
    }

    if (event === "emptyCoordinate") {
        OpenCoordinate(jsonMessage)
    }

    if (event === "SelectCoordinateCreate") {
        SelectCoordinateCreate(jsonMessage)
    }

    if (event === "OpenCoordinate") { +
        OpenCoordinate(jsonMessage)
    }

    if (event === "DellCoordinate") { +
        DelUnit(jsonMessage)
    }

    if (event === "MouseOver") { +
        ReadInfoMouseOver(jsonMessage);
    }

    if (event === "MoveUnit") {
        console.log(jsonMessage);
        InitMoveUnit(jsonMessage);
    }

    if (event === "TargetUnit") {
        TargetUnit();
    }
}