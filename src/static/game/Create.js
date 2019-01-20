let LoadFunc;
let Map;

function CreateGame(map, loadFunc) {
    LoadFunc = loadFunc;
    Map = map;

    return new Phaser.Game('100', '100', Phaser.WEBGL, 'main', {
        preload: preload,
        create: create,
        update: update,
        render: render
    });
}

function create(game) {


    // размеры гексов карты по умолчанию
    game.hexagonWidth = 100;
    game.hexagonHeight = 111;
    // параметры смещения тени игры
    game.shadowXOffset = 8;
    game.shadowYOffset = 10;
    // игровая карта
    game.map = Map;

    game.physics.startSystem(Phaser.Physics.ARCADE);

    game.time.advancedTiming = true; // настройка fts
    game.time.desiredFps = 60;       // макс фпс 60
    game.time.slowMotion = 0;        // плавный переход в мин фпс

    game.stage.disableVisibilityChange = true; // не дает уснуть игры при сворачивание браузера

    game.world.setBounds(0, 0, (game.hexagonWidth + 5) * game.map.QSize, 185 * game.map.RSize / 2); //размеры карты
    game.stage.backgroundColor = "#242424"; //цвет фона

    game.bmdTerrain = game.make.bitmapData((game.hexagonWidth + 5) * game.map.QSize, 185 * game.map.RSize / 2);
    game.add.image(0, 0, game.bmdTerrain); //bitmapData для отрисовки статичного нижнего слоя

    game.floorLayer = game.add.group();
    game.floorSelectLineLayer = game.add.group();
    game.floorObjectSelectLineLayer = game.add.group();

    // уровень обьектов которые под юнитом
    game.floorObjectLayer = game.add.group();

    // UNITS
    game.unitLayer = game.add.group();

    // уровень обьектов которые над юнитом
    game.floorOverObjectLayer = game.add.group();

    game.SelectLayer = game.add.group();
    game.SelectLayer.alpha = 0.4;

    game.SelectRangeLayer = game.add.group();
    game.SelectRangeLayer.alpha = 0.6;

    game.SelectLineLayer = game.add.group();
    game.SelectLineLayer.alpha = 0.9;
    game.add.tween(game.SelectLineLayer).to({alpha: 0.4}, 1500, "Linear").loop(true).yoyo(true).start();

    game.SelectTargetLineLayer = game.add.group();
    game.SelectTargetLineLayer.alpha = 0.9;
    game.add.tween(game.SelectTargetLineLayer).to({alpha: 0.4}, 1500, "Linear").loop(true).yoyo(true).start();

    game.effectsLayer = game.add.group();

    game.artilleryBulletLayer = game.add.group();
    game.weaponEffectsLayer = game.add.group();

    game.flyObjectsLayer = game.add.group();

    game.cloudsLayer = game.add.group();

    game.fogOfWar = game.add.group();
    game.fogOfWar.alpha = 0.5;

    game.redactorButton = game.add.group();
    game.redactorMetaText = game.add.group();

    game.icon = game.add.group();

    CreateMap().then(function () {
        if (LoadFunc) {
            LoadFunc();
        }

        if (game.map.reservoir) {
            CreateReservoirs()
        }

        if (game.typeService === "battle") {
            CreateMyGameUnits();
            CreateHostileGameUnits();
            LoadOpenCoordinate();
        }
    });
}