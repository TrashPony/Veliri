let game;
let Data;
let debug = true;

function Game(jsonData) {
    Data = jsonData;
    game = CreateGame(jsonData.map, LoadGame, "global");
}

function LoadGame() {

    game.input.onUp.add(StopSelectableUnits, this);
    game.input.onUp.add(UnSelectUnit, this);

    // запоминаем последние тайминги нажатий
    game.input.onUp.add(function (pointer) {
        if (game.input.activePointer.leftButton.isDown) {
            game.input.activePointer.leftButton.lastDuration = pointer.duration;
        } else if (game.input.activePointer.rightButton.isDown) {
            game.input.activePointer.rightButton.lastDuration = pointer.duration;
        }
    });

    game.bmdTerrain.sprite.inputEnabled = true;
    game.bmdTerrain.sprite.input.priorityID = 1;
    game.bmdTerrain.sprite.events.onInputUp.add(initMove);

    // TODO
    // game.camera.scale.x = 1.5;
    // game.camera.scale.y = 1.5;

    game.user_name = Data.user.login;
    game.user_id = Data.user.id;

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

    CreateDynamicObjects(Data.dynamic_objects)

    setTimeout(function () {
        CreateMiniMap();
        global.send(JSON.stringify({
            event: "RefreshRadar"
        }));
        if (debug) {
            CreateAnomalies(Data.map.anomalies)
        }
    }, 1000);
}