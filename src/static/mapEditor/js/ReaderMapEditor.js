function ReaderMapEditor(jsonMessage) {
    if (JSON.parse(jsonMessage).event === "MapList") {
        CreateMapList(jsonMessage)
    }
    if (JSON.parse(jsonMessage).event === "MapSelect") {
        createGame(jsonMessage)
    }
}