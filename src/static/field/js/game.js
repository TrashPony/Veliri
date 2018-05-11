var game;
var tileWidth = 100; // ширина и высота спрайта в сетке грида
var User;
var GameMap;

var UNIT_SPEED = 100;
var TARGET_MOVE_RANGE = 10;

var cells = {};    // карта со всеми ячейками брать их так var cell = cells[1+":"+1];
var units = {};    // карта со всеми юнитами брать их так var unit = units[1+":"+1];

function Game(gameMap) {
    GameMap = gameMap;

    var width;//получаем ширину монитора
    var height; //получаем высоту монитора

    if (window.innerWidth < tileWidth * gameMap.XSize) {
        width = window.innerWidth;
    } else {
        width = tileWidth * gameMap.XSize
    }

    if (window.innerHeight < tileWidth * gameMap.YSize) {
        height = window.innerHeight;
    } else {
        height = tileWidth * gameMap.YSize;
    }

    game = new Phaser.Game(width, height, Phaser.CANVAS, 'main', {
        preload: preload,
        create: create,
        update: update,
        render: render
    }); //создаем игровое поле с высотой и шир
}

function preload() {
    game.stage.disableVisibilityChange = true; // не дает оставиться выполнения скрипта если окно скрыто
    // Units
    game.load.image('scout', 'http://' + window.location.host + '/field/img/tank.png');
    // Structures
    game.load.image('respawn', 'http://' + window.location.host + '/field/img/respawn.png');
    // Map Objects
    game.load.image('obstacle', 'http://' + window.location.host + '/field/img/obstacle.png');
    game.load.image('floor', 'http://' + window.location.host + '/field/img/openCell.jpg');
    // Interface marks
    game.load.image('MarkMoveLastCell', 'http://' + window.location.host + '/field/img/toMove.png');
    // Fonts
    game.load.bitmapFont('carrier_command', 'https://examples.phaser.io/assets/fonts/bitmapFonts/carrier_command.png', 'https://examples.phaser.io/assets/fonts/bitmapFonts/carrier_command.xml');
}

function create() {

    game.physics.startSystem(Phaser.Physics.ARCADE);

    game.time.advancedTiming = true; // настройка fts
    game.time.desiredFps = 60;       // макс фпс 60
    game.time.slowMotion = 0;        // плавный переход в мин фпс

    game.world.setBounds(0, 0, tileWidth * GameMap.XSize, tileWidth * GameMap.YSize); //размеры карты

    game.stage.backgroundColor = "#242424"; //цвет фона


    CreateMap(GameMap)
}

function update() {
    MoveUnit();
    GrabCamera(); // функцуия для перетаскивания карты мышкой /* Магия */
}

function render() {

}


