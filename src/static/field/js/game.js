var game;
var tileWidth = 100; // ширина и высота спрайта в сетке грида
var MY_ID;
var countWidthGrid;
var countHeightGrid;

var UNIT_SPEED = 100;
var TARGET_MOVE_RANGE = 10;

var cells = {};    // карта со всеми ячейками брать их так var cell = cells[1+":"+1];
var units = {};    // карта со всеми юнитами брать их так var unit = units[1+":"+1];

function Game(x, y) {

    countWidthGrid = x;
    countHeightGrid = y;

    var width;//получаем ширину монитора
    var height; //получаем высоту монитора

    if (window.innerWidth < tileWidth * countWidthGrid) {
        width = window.innerWidth;
    } else {
        width = tileWidth * countWidthGrid
    }

    if (window.innerHeight < tileWidth * countHeightGrid) {
        height = window.innerHeight;
    } else {
        height = tileWidth * countHeightGrid;
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
    game.load.image('floor', 'http://642e0559eb9c.sn.mynetname.net:8080/field/img/openCell.jpg');
    game.load.image('scout', 'http://642e0559eb9c.sn.mynetname.net:8080/field/img/tank.png');
    game.load.image('obstacle', 'http://642e0559eb9c.sn.mynetname.net:8080/field/img/obstacle.png');

    game.load.bitmapFont('carrier_command', 'https://examples.phaser.io/assets/fonts/bitmapFonts/carrier_command.png', 'https://examples.phaser.io/assets/fonts/bitmapFonts/carrier_command.xml');
}

function create() {

    game.physics.startSystem(Phaser.Physics.ARCADE);

    game.time.advancedTiming = true; // настройка fts
    game.time.desiredFps = 60;       // макс фпс 60
    game.time.slowMotion = 0;        // плавный переход в мин фпс

    game.world.setBounds(0, 0, tileWidth * countWidthGrid, tileWidth * countHeightGrid); //размеры карты

    game.stage.backgroundColor = "#242424"; //цвет фона


    for (var x = 0; x < countWidthGrid; x++) {
        for (var y = 0; y < countHeightGrid; y++) {
            var floorSprite = game.add.tileSprite(x * tileWidth, y * tileWidth, tileWidth, tileWidth, 'floor');

            game.add.bitmapText(x * tileWidth + tileWidth / 2, y * tileWidth + tileWidth / 2, 'carrier_command', x + ":" + y, 12);
            floorSprite.id = x + ":" + y;
            floorSprite.tint = 0x757575;
            floorSprite.inputEnabled = true; // включаем ивенты на спрайт
            floorSprite.events.onInputDown.add(SelectTarget, this);
            cells[floorSprite.id] = floorSprite;
        }
    }
}

function update() {
    MoveUnit();
    GrabCamera(); // функцуия для перетаскивания карты мышкой /* Магия */
}

function render() {

}


