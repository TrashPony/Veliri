function create() {

    game.physics.startSystem(Phaser.Physics.ARCADE);

    game.time.advancedTiming = true; // настройка fts
    game.time.desiredFps = 60;       // макс фпс 60
    game.time.slowMotion = 0;        // плавный переход в мин фпс

    game.stage.disableVisibilityChange = true; // не дает уснуть игры при сворачивание браузера
    // todo  0.87 подгон под тестовую карту, костыль
    game.world.setBounds(0, 0, (game.hexagonHeight * game.map.RSize) * 0.87, game.hexagonWidth * game.map.QSize); //размеры карты
    game.stage.backgroundColor = "#242424"; //цвет фона

    game.floorLayer = game.add.group();
    game.floorObjectLayer = game.add.group();

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

    game.fogOfWar = game.add.group();
    game.fogOfWar.alpha = 0.5;

    CreateMap();
    CreateMyGameUnits();
    CreateHostileGameUnits();
    LoadOpenCoordinate();
}