var game;

function LoadGame(jsonMessage) {
    var gameMap = JSON.parse(jsonMessage).map;

    var tileSize = 100; // ширина и высота спрайта в сетке грида
    game = new Phaser.Game(GetWidth(tileSize, gameMap), GetHeight(tileSize, gameMap), Phaser.CANVAS, 'main', {
        preload: preload,
        create: create,
        update: update,
        render: render
    }); //создаем игровое поле с высотой и шир

    game.user = {};
    game.user.name = JSON.parse(jsonMessage).user_name;
    game.user.ready = JSON.parse(jsonMessage).ready;
    game.user.equip = JSON.parse(jsonMessage).equip;

    game.Step = JSON.parse(jsonMessage).game_step;
    game.Phase = JSON.parse(jsonMessage).game_phase;

    game.unitStorage = JSON.parse(jsonMessage).unit_storage;
    game.tileSize = tileSize;

    // Creates objects
    game.map = gameMap;
    game.units = JSON.parse(jsonMessage).units;
    game.hostileUnits = JSON.parse(jsonMessage).hostile_units;
    game.matherShip = JSON.parse(jsonMessage).mather_ship;
    game.hostileMatherShips = JSON.parse(jsonMessage).hostile_mather_ships;
    game.user.watch = JSON.parse(jsonMessage).watch;

    game.map.selectSprites = [];

    GameInfo();
    InitPlayer();
    LoadHoldUnits();
}