function FieldCreate(jsonMessage) {
    var gameMap = JSON.parse(jsonMessage).Map;

    Game(gameMap) // создаем окно игры размером х:у
}