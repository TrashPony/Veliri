function autoPreload() { // файл который заполняет автоматом при добавление новый координат
    game.load.image('0032', 'http://' + window.location.host + '/assets/map/objects/0032.png');
    game.load.spritesheet('256_256', 'http://' + window.location.host + '/assets/map/animate/256_256.png', 256, 256, 60);
    game.load.spritesheet('256_256', 'http://' + window.location.host + '/assets/map/animate/256_256.png', 256, 256, 60);
    game.load.image('Mother', 'http://' + window.location.host + '/assets/map/objects/Mother.png');
    game.load.image('beach2_tl', 'http://' + window.location.host + '/assets/map/objects/beach2_tl.png');
    game.load.image('beach2_tm_01', 'http://' + window.location.host + '/assets/map/objects/beach2_tm_01.png');
    game.load.spritesheet('tunel', 'http://' + window.location.host + '/assets/map/animate/tunel.png', 540, 960, 2);
    game.load.spritesheet('baseCore', 'http://' + window.location.host + '/assets/map/animate/baseCore.png', 256, 256, 50);
    game.load.spritesheet('recycler', 'http://' + window.location.host + '/assets/map/animate/recycler.png', 256, 256, 50);
    game.load.image('destroySpaceShip', 'http://' + window.location.host + '/assets/map/objects/destroySpaceShip.png');
}