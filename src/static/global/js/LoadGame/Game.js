let game;

function Game(jsonData) {
    game = CreateGame(jsonData.map);
    game.typeService = "global";

    setTimeout(function () { // todo костыль связаной с прогрузкой карты )
        CreateUser(jsonData.squad);
        game.input.onDown.add(initMove, game);
        CreateBase(jsonData.bases);
        CreateOtherUsers(jsonData.other_users);
        CreateMiniMap(jsonData.map);
        CreateBoxes(jsonData.boxes)
    }, 1500);
}