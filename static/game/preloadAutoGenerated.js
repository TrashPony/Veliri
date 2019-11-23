function autoPreload() { // файл который заполняет автоматом при добавление новый координат
    game.load.spritesheet('baseCore', 'http://' + window.location.host + '/assets/map/animate/baseCore.png', 256, 256, 50);
    game.load.spritesheet('recycler', 'http://' + window.location.host + '/assets/map/animate/recycler.png', 256, 256, 50);
    game.load.image('destroySpaceShip', 'http://' + window.location.host + '/assets/map/objects/destroySpaceShip.png');
    game.load.spritesheet('danger_becon', 'http://' + window.location.host + '/assets/map/animate/danger_becon.png', 512, 512, 2);
}