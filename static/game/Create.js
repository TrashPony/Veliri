let LoadFunc;
let Map;
let TypeService;

function CreateGame(map, loadFunc, typeService) {
    LoadFunc = loadFunc;
    Map = map;
    TypeService = typeService;

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

    game.bulletLayer = game.add.group();
    // уровень обьектов которые над юнитом
    game.floorOverObjectLayer = game.add.group();

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
    };

    let UnitStatusBMD = game.make.bitmapData(game.camera.width, game.camera.height);
    let UnitStatusSprite = UnitStatusBMD.addToWorld();
    UnitStatusSprite.fixedToCamera = true;
    game.UnitStatusLayer = {
        bmd: UnitStatusBMD,
        sprite: UnitStatusSprite,
    };

    game.redactorButton = game.add.group();
    game.SelectLayer = game.add.group();
    game.redactorMetaText = game.add.group();
    game.geoDataLayer = game.add.group();

    game.icon = game.add.group();

    game.typeService = TypeService;
    game.canvas.id = 'GameCanvas';

    CreateMap();

    if (LoadFunc) {
        LoadFunc();
    }
}