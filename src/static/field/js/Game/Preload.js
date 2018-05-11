function preload() {
    game.stage.disableVisibilityChange = true; // не дает оставиться выполнения скрипта если окно скрыто
    //MatherShips MotherTrucker
    game.load.image('MotherTrucker', 'http://' + window.location.host + '/assets/MotherTrucker.png');
    // Units
    game.load.image('scout', 'http://' + window.location.host + '/assets/tank.png');
    // Structures
    game.load.image('respawn', 'http://' + window.location.host + '/assets/respawn.png');
    // Map Objects
    game.load.image('obstacle', 'http://' + window.location.host + '/assets/obstacle.png');
    game.load.image('floor', 'http://' + window.location.host + '/assets/openCell.jpg');
    // Interface marks
    game.load.image('MarkMoveLastCell', 'http://' + window.location.host + '/assets/toMove.png');
    // Fonts
    game.load.bitmapFont('carrier_command', 'https://examples.phaser.io/assets/fonts/bitmapFonts/carrier_command.png', 'https://examples.phaser.io/assets/fonts/bitmapFonts/carrier_command.xml');
}