let game;

function LoadGame(jsonData) {
    game = CreateGame(jsonData.map);
    game.typeService = "global";

    setTimeout(function () { // todo костыль связаной с прогрузкой карты )
        CreateSquad(jsonData.squad);
        game.input.onDown.add(initMove, game);
    }, 1500);
}