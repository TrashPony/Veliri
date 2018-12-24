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
    CreateOtherUsers(Data.other_users);
    CreateBase(Data.bases);
    CreateBoxes(Data.boxes);
    CreateMiniMap(Data.map);
}