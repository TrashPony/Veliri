function ReaderMapEditor(jsonMessage) {
    if (JSON.parse(jsonMessage).event === "MapList") {
        CreateMapList(jsonMessage)
    }
    if (JSON.parse(jsonMessage).event === "MapSelect") {
        createGame(jsonMessage)
    }

    if (JSON.parse(jsonMessage).event === "getAllTypeCoordinate") {
        ViewAllTypeCoordinate(JSON.parse(jsonMessage).type_coordinates)
    }
}