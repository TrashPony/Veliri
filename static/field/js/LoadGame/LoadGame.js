let game;

function LoadGame(jsonMessage) {

    game = CreateGame(
        JSON.parse(jsonMessage).map,
        function () {
            game.user = {};
            game.user.name = JSON.parse(jsonMessage).user_name;
            game.user.ready = JSON.parse(jsonMessage).ready;
            game.user.equip = JSON.parse(jsonMessage).equip;

            game.Step = JSON.parse(jsonMessage).game_step;
            game.Phase = JSON.parse(jsonMessage).game_phase;

            game.unitStorage = JSON.parse(jsonMessage).unit_storage;

            // Creates objects
            game.units = JSON.parse(jsonMessage).units;
            game.hostileUnits = JSON.parse(jsonMessage).hostile_units;
            game.memoryHostileUnit = JSON.parse(jsonMessage).memory_hostile_unit;
            game.user.watch = JSON.parse(jsonMessage).watch;

            game.map.selectSprites = [];

            game.camera.scale.x = 1.5;
            game.camera.scale.y = 1.5;

            GameInfo();
            InitPlayer();
            LoadHoldUnits();
            LoadQueueUnits();
            CreateMyGameUnits();
            CreateHostileGameUnits();
            LoadOpenCoordinate();
            CreateMiniMap();
        },
        "battle");
}