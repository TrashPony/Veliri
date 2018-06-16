function create() {

    game.physics.startSystem(Phaser.Physics.ARCADE);

    game.time.advancedTiming = true; // настройка fts
    game.time.desiredFps = 60;       // макс фпс 60
    game.time.slowMotion = 0;        // плавный переход в мин фпс

    game.stage.disableVisibilityChange = true; // не дает уснуть игры при сворачивание браузера

    game.world.setBounds(0, 0, game.tileSize * game.map.XSize, game.tileSize * game.map.YSize); //размеры карты

    game.stage.backgroundColor = "#242424"; //цвет фона

    game.floorLayer = game.add.group();

    game.SelectLayer = game.add.group();
    game.SelectLayer.alpha = 0.3;

    game.SelectLineLayer = game.add.group();
    game.SelectLineLayer.alpha = 0.7;
    game.add.tween(game.SelectLineLayer).to( { alpha: 0.2 }, 1500, "Linear").loop(true).yoyo(true).start();

    game.SelectTargetLineLayer = game.add.group();
    game.SelectTargetLineLayer.alpha = 0.7;
    game.add.tween(game.SelectTargetLineLayer).to( { alpha: 0.2 }, 1500, "Linear").loop(true).yoyo(true).start();

    game.floorObjectLayer = game.add.group();

    game.fogOfWar = game.add.group();
    game.fogOfWar.alpha = 0.5;

    CreateMap();
    CreateMyGameUnits();
    CreateHostileGameUnits();
    CreateMyMatherShip();
    CreateHostileMatherShips();
    LoadOpenCoordinate();
}