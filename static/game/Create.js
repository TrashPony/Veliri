let LoadFunc;
let Map;
let TypeService;

function CreateGame(map, loadFunc, typeService) {
    LoadFunc = loadFunc;
    Map = map;
    TypeService = typeService;
    //TODO что бы работал блюр на линиях и эмиторы надо делать WEBGL
    return new Phaser.Game('100', '100', Phaser.Canvas, 'main', {
        preload: preload,
        create: create,
        update: update,
        render: render
    });
}

function create(game) {
    // параметры смещения тени игры
    game.shadowXOffset = 4;
    game.shadowYOffset = 5;
    // игровая карта
    game.map = Map;

    game.physics.startSystem(Phaser.Physics.ARCADE);

    game.time.advancedTiming = true; // настройка fts
    game.time.desiredFps = 60;       // макс фпс 60
    game.time.slowMotion = 1;        // плавный переход в мин фпс

    //game.stage.disableVisibilityChange = true; // не дает уснуть игры при сворачивание браузера
    game.world.setBounds(0, 0, game.map.XSize, game.map.YSize); //размеры карты
    game.stage.backgroundColor = "#242424"; //цвет фона

    game.bmdTerrain = game.make.bitmapData(game.map.XSize, game.map.YSize);
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

    game.PreviewPath = game.add.group();
    game.PreviewPath.alpha = 0.6;

    game.SelectLineLayer = game.add.group();
    game.SelectLineLayer.alpha = 0.9;
    game.add.tween(game.SelectLineLayer).to({alpha: 0.4}, 1500, "Linear").loop(true).yoyo(true).start();

    game.GameZone = game.add.group();
    game.GameZone.alpha = 0.9;
    game.add.tween(game.GameZone).to({alpha: 0.4}, 1500, "Linear").loop(true).yoyo(true).start();

    game.SelectTargetLineLayer = game.add.group();
    game.SelectTargetLineLayer.alpha = 0.9;
    game.add.tween(game.SelectTargetLineLayer).to({alpha: 0.4}, 1500, "Linear").loop(true).yoyo(true).start();

    game.effectsLayer = game.add.group();

    game.artilleryBulletLayer = game.add.group();
    game.weaponEffectsLayer = game.add.group();

    game.flyObjectsLayer = game.add.group();

    game.cloudsLayer = game.add.group();

    let fogBmd = game.make.bitmapData(game.camera.width, game.camera.height);
    let fogSprite = fogBmd.addToWorld();
    fogSprite.fixedToCamera = true;
    game.FogOfWar = {
        bmd: fogBmd,
        sprite: fogSprite,
        overviewCircle: null,
        ms: null,
    };

    game.redactorButton = game.add.group();
    game.redactorMetaText = game.add.group();
    game.geoDataLayer = game.add.group();

    game.icon = game.add.group();

    game.typeService = TypeService;

    CreateMap();

    if (LoadFunc) {
        LoadFunc();
    }

    if (game.map.reservoir) {
        CreateReservoirs()
    }
}