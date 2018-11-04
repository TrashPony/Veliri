let game;

function createGame(jsonMessage) {
    game = new Phaser.Game('100', '100', Phaser.CANVAS, 'map', {
        preload: preload,
        create: create,
        update: update,
        render: render
    });

    game.hexagonWidth = 80;
    game.hexagonHeight = 100;
    game.map = JSON.parse(jsonMessage).map;
}

function create() {
    game.physics.startSystem(Phaser.Physics.ARCADE);
    game.time.advancedTiming = true;
    game.time.desiredFps = 60;
    game.time.slowMotion = 0;
    game.stage.disableVisibilityChange = true; // не дает уснуть игры при сворачивание браузера
    game.stage.backgroundColor = "#242424"; //цвет фона

    game.floorLayer = game.add.group();
    game.floorObjectLayer = game.add.group();
    game.effectsLayer = game.add.group();

    game.world.setBounds(0, 0, game.hexagonHeight * game.map.QSize, game.hexagonHeight * game.map.RSize); //размеры карты
    CreateMap();
}

function preload() {
    game.stage.disableVisibilityChange = true; // не дает оставиться выполнения скрипта если окно скрыто

    game.load.image('labelEffects', 'http://' + window.location.host + '/assets/effects/label_effects.png');

    // Map Objects
    game.load.image('hexagon', 'http://' + window.location.host + '/assets/map/hexagon.png');

    game.load.image('terrain_1', 'http://' + window.location.host + '/assets/map/tree1.png');
    game.load.image('terrain_2', 'http://' + window.location.host + '/assets/map/tree2.png');
    game.load.image('terrain_3', 'http://' + window.location.host + '/assets/map/tree3.png');
    game.load.image('crater', 'http://' + window.location.host + '/assets/map/crater.png');

    game.load.image('sand_stone_04', 'http://' + window.location.host + '/assets/map/sand_stone_04.png');
    game.load.image('sand_stone_05', 'http://' + window.location.host + '/assets/map/sand_stone_05.png');
    game.load.image('sand_stone_06', 'http://' + window.location.host + '/assets/map/sand_stone_06.png');
    game.load.image('sand_stone_07', 'http://' + window.location.host + '/assets/map/sand_stone_07.png');
    game.load.image('sand_stone_08', 'http://' + window.location.host + '/assets/map/sand_stone_08.png');

    game.load.image('fallen_01', 'http://' + window.location.host + '/assets/map/fallen_01.png');
    game.load.image('fallen_02', 'http://' + window.location.host + '/assets/map/fallen_02.png');
    // Fonts
    game.load.bitmapFont('carrier_command', 'https://examples.phaser.io/assets/fonts/bitmapFonts/carrier_command.png', 'https://examples.phaser.io/assets/fonts/bitmapFonts/carrier_command.xml');
    //  Load the Google WebFont Loader script
    game.load.script('webfont', '//ajax.googleapis.com/ajax/libs/webfont/1.4.7/webfont.js');
}

function render() {

}

function GrabCamera() {
    if (game.input.activePointer.rightButton.isDown) { // ловит нажатие правой кнопки маши в игре
        if (game.origDragPoint) {
            game.camera.x += game.origDragPoint.x - game.input.activePointer.position.x; // перемещать камеру по сумме, перемещенную мышью с момента последнего обновления
            game.camera.y += game.origDragPoint.y - game.input.activePointer.position.y;
        }
        game.origDragPoint = game.input.activePointer.position.clone(); // установите новое начало перетаскивания в текущую позицию
    } else {
        game.origDragPoint = null;
    }
}

function update() {
    GrabCamera(); // функцуия для перетаскивания карты мышкой /* Магия */
    game.floorObjectLayer.sort('y', Phaser.Group.SORT_ASCENDING);
}