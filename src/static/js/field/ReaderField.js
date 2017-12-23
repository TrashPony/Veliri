function ReadResponse(jsonMessage) {
    var event = JSON.parse(jsonMessage).event;

    var x;
    var y;

    if (event === "InitPlayer") {
        InitPlayer(jsonMessage);
    }

    if (event === "InitMap") {
        FieldCreate(jsonMessage)
    }

    if (event === "InitUnit") {
        InitUnit(jsonMessage);
    }

    if (event === "InitStructure") {
        InitStructure(jsonMessage);
    }

    if (event === "InitObstacle") {
        InitObstacle(jsonMessage);
    }

    if (event === "CreateUnit") {
        CreateUnit(jsonMessage);
    }

    if (event === "Ready") {
        ReadyReader(jsonMessage);
    }

    if (event === "SelectUnit") {
        setUnitAction(jsonMessage);
    }

    if (event === "Attack") {
        AttackUnit(jsonMessage);
    }

    if (event === "emptyCoordinate") {
        EmptyCoordinate(jsonMessage)
    }

    if (event === "SelectCoordinateCreate") {
        SelectCoordinateCreate(jsonMessage)
    }

    if (event === "OpenCoordinate") {
        x = JSON.parse(jsonMessage).x;
        y = JSON.parse(jsonMessage).y;
        var idCell = x + ":" + y;
        OpenCoordinate(idCell)
    }

    if (event === "DellCoordinate") {
        x = JSON.parse(jsonMessage).x;
        y = JSON.parse(jsonMessage).y;
        var idDell = x + ":" + y;
        DelUnit(idDell)
    }

    if (event === "MouseOver") {
        ReadInfoMouseOver(jsonMessage);
    }

    if (event === "MoveUnit") {
        MoveUnit(jsonMessage);
    }

    if (event === "TargetUnit") {
        TargetUnit();
    }
}