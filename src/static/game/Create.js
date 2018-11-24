function CreateGame(map) {
    let game = new Phaser.Game('100', '100', Phaser.AUTO, 'main', {
        preload: preload,
        create: create,
        update: update,
        render: render
    });

    // игровая карта, без карты нельзя построить игру
    game.map = map;

    // размеры гексов карты по умолчанию
    game.hexagonWidth = 100;
    game.hexagonHeight = 111;

    // параметры смещения тени игры
    game.shadowXOffset = 8;
    game.shadowYOffset = 10;

    return game
}

function create(game) {

    game.physics.startSystem(Phaser.Physics.ARCADE);

    game.time.advancedTiming = true; // настройка fts
    game.time.desiredFps = 60;       // макс фпс 60
    game.time.slowMotion = 0;        // плавный переход в мин фпс

    game.stage.disableVisibilityChange = true; // не дает уснуть игры при сворачивание браузера
    game.world.setBounds(0, 0, (game.hexagonWidth + 5) * game.map.QSize, 185 * game.map.RSize/2); //размеры карты
    game.stage.backgroundColor = "#242424"; //цвет фона

    game.floorLayer = game.add.group();
    game.floorObjectLayer = game.add.group();

    game.unitLayer = game.add.group();

    game.SelectLayer = game.add.group();
    game.SelectLayer.alpha = 0.4;

    game.SelectRangeLayer = game.add.group();
    game.SelectRangeLayer.alpha = 0.6;

    game.SelectLineLayer = game.add.group();
    game.SelectLineLayer.alpha = 0.9;
    game.add.tween(game.SelectLineLayer).to( { alpha: 0.4 }, 1500, "Linear").loop(true).yoyo(true).start();

    game.SelectTargetLineLayer = game.add.group();
    game.SelectTargetLineLayer.alpha = 0.9;
    game.add.tween(game.SelectTargetLineLayer).to( { alpha: 0.4 }, 1500, "Linear").loop(true).yoyo(true).start();

    game.effectsLayer = game.add.group();

    game.artilleryBulletLayer = game.add.group();
    game.weaponEffectsLayer = game.add.group();

    game.fogOfWar = game.add.group();
    game.fogOfWar.alpha = 0.5;

    CreateMap();
    if (game && game.typeService === "battle") {
        CreateMyGameUnits();
        CreateHostileGameUnits();
        LoadOpenCoordinate();
    }
}