let game;
let Data;

function Game(jsonData) {
    Data = jsonData;
    game = CreateGame(jsonData.map, LoadGame);
    game.typeService = "global";
    game.evacuations = [];
}

function LoadGame() {
    game.input.onDown.add(initMove, game);

    CreateUser(Data.squad);
    CreateBase(Data.bases);
    CreateOtherUsers(Data.other_users);
    CreateMiniMap(Data.map);
    CreateBoxes(Data.boxes)
}