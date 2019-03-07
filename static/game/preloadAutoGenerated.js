function autoPreload() { // файл который заполняет автоматом при добавление новый координат
    game.load.image('0032', 'http://' + window.location.host + '/assets/map/objects/0032.png');
    game.load.image('Mother', 'http://' + window.location.host + '/assets/map/objects/Mother.png');
    game.load.image('beach2_tl', 'http://' + window.location.host + '/assets/map/objects/beach2_tl.png');
    game.load.image('beach2_tm_01', 'http://' + window.location.host + '/assets/map/objects/beach2_tm_01.png');
    game.load.spritesheet('tunel', 'http://' + window.location.host + '/assets/map/animate/tunel.png', 540, 960, 2);
    game.load.spritesheet('baseCore', 'http://' + window.location.host + '/assets/map/animate/baseCore.png', 256, 256, 50);
    game.load.spritesheet('recycler', 'http://' + window.location.host + '/assets/map/animate/recycler.png', 256, 256, 50);
    game.load.image('destroySpaceShip', 'http://' + window.location.host + '/assets/map/objects/destroySpaceShip.png');
    game.load.spritesheet('danger_becon', 'http://' + window.location.host + '/assets/map/animate/danger_becon.png', 512, 512, 2);
    game.load.image('mountain_1', 'http://' + window.location.host + '/assets/map/objects/mountain_1.png');
    game.load.image('mountain_1_2', 'http://' + window.location.host + '/assets/map/objects/mountain_1_2.png');
    game.load.image('mountain_2', 'http://' + window.location.host + '/assets/map/objects/mountain_2.png');
    game.load.image('shaurma', 'http://' + window.location.host + '/assets/map/objects/shaurma.png');
}