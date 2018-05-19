function create() {

    game.physics.startSystem(Phaser.Physics.ARCADE);

    game.time.advancedTiming = true; // настройка fts
    game.time.desiredFps = 60;       // макс фпс 60
    game.time.slowMotion = 0;        // плавный переход в мин фпс

    game.world.setBounds(0, 0, game.tileSize * game.map.XSize, game.tileSize * game.map.YSize); //размеры карты

    game.stage.backgroundColor = "#242424"; //цвет фона

    game.floorLayer = game.add.group();

    game.SelectLayer = game.add.group();
    game.SelectLayer.alpha = 0.3;

    game.SelectLineLayer = game.add.group();
    game.SelectLineLayer.alpha = 0.7;

    game.floorObjectLayer = game.add.group();



    CreateMap();
    CreateMyGameUnits();
    CreateHostileGameUnits();
    CreateMyMatherShip();
    CreateHostileMatherShips();
    LoadOpenCoordinate();
}