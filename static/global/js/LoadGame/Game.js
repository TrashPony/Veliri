let game;
let Data;
let debug = false;

function Game(jsonData) {
    Data = jsonData;
    game = CreateGame(jsonData.map, LoadGame, "global");
    game.evacuations = [];
}

function LoadGame() {

    game.input.onUp.add(initMove, this);
    game.input.onUp.add(StopSelectableUnits, this);
    game.input.onUp.add(UnSelectUnit, this);

    game.camera.scale.x = 1.5;
    game.camera.scale.y = 1.5;

    game.user_name = Data.user.login;
    game.user_id = Data.user.id;
    game.my_squad_sprite = {};

    game.units = Data.short_units;

    CreateUnits(game.units);
    CreateBase(Data.bases);
    CreateBoxes(Data.boxes);
    ThoriumBar(Data.squad.mather_ship.body.thorium_slots);
    FillSquadBlock(Data.squad);
    Anomaly(Data.squad);

    //todo FillUserMeta(Data.credits, Data.experience, Data.squad);


    setTimeout(function () {
        CreateMiniMap();
        if (debug) {
            CreateGeoData(Data.map.geo_data);
            CreateAnomalies(Data.map.anomalies)
        }
    }, 1000)
}