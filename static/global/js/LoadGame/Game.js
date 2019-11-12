let game;
let Data;
let debug = true;

function Game(jsonData) {
    Data = jsonData;
    game = CreateGame(jsonData.map, LoadGame, "global");
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

    game.units = {};
    game.bullets = {};
    game.radar_marks = {};
    game.boxes = [];
    game.map.reservoir = [];

    CreateBase(Data.bases);
    ThoriumBar(Data.squad.mather_ship.body.thorium_slots);
    FillSquadBlock(Data.squad);
    Anomaly(Data.squad);

    FillUserMeta(Data.credits, Data.experience, Data.squad);
    ChangeGravity(Data.high_gravity);

    for (let x in Data.dynamic_objects) {
        for (let y in Data.dynamic_objects[x]) {
            let object = Data.dynamic_objects[x][y];
            if (object.texture !== '') {
                CreateObject(object, object.x, object.y);
            }
            if (object.animate_sprite_sheets !== '') {
                CreateAnimate(object, object.x, object.y);
            }
        }
    }

    setTimeout(function () {
        CreateMiniMap();
        global.send(JSON.stringify({
            event: "RefreshRadar"
        }));
        if (debug) {
            CreateGeoData(Data.map.geo_data);
            CreateAnomalies(Data.map.anomalies)
        }
    }, 1000);


}