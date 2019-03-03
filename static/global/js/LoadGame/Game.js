let game;
let Data;
let debug = false;

function Game(jsonData) {
    Data = jsonData;
    game = CreateGame(jsonData.map, LoadGame);
    game.typeService = "global";
    game.evacuations = [];
}

function LoadGame() {
    game.input.onDown.add(initMove, game);

    game.camera.scale.x = 1.5;
    game.camera.scale.y = 1.5;

    CreateUser(Data.squad);
    CreateOtherUsers(Data.other_users);
    CreateBase(Data.bases);
    CreateBoxes(Data.boxes);
    CreateMiniMap(Data.map);
    ThoriumBar(Data.squad.mather_ship.body.thorium_slots);
    FillSquadBlock(Data.squad);
    FillUserMeta(Data.credits, Data.experience, Data.squad);
    Anomaly(Data.squad);

    if (debug) {
        CreateGeoData(Data.map.geo_data);
    }

    FocusMS();
}