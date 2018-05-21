function preload() {
    game.stage.disableVisibilityChange = true; // не дает оставиться выполнения скрипта если окно скрыто

    // Select
    game.load.image('selectPlace_1', 'http://' + window.location.host + '/assets/select/place1.png');
    game.load.image('selectPlace_2', 'http://' + window.location.host + '/assets/select/place2.png');
    game.load.image('selectPlace_3', 'http://' + window.location.host + '/assets/select/place3.png');
    game.load.image('selectPlace_4', 'http://' + window.location.host + '/assets/select/place4.png');
    game.load.image('selectPlace_5', 'http://' + window.location.host + '/assets/select/place5.png');

    game.load.image('selectMove_1', 'http://' + window.location.host + '/assets/select/move1.png');
    game.load.image('selectMove_2', 'http://' + window.location.host + '/assets/select/move2.png');
    game.load.image('selectMove_3', 'http://' + window.location.host + '/assets/select/move3.png');
    game.load.image('selectMove_4', 'http://' + window.location.host + '/assets/select/move4.png');
    game.load.image('selectMove_5', 'http://' + window.location.host + '/assets/select/move5.png');

    game.load.image('selectTarget_1', 'http://' + window.location.host + '/assets/select/target1.png');
    game.load.image('selectTarget_2', 'http://' + window.location.host + '/assets/select/target2.png');
    game.load.image('selectTarget_3', 'http://' + window.location.host + '/assets/select/target3.png');

    game.load.spritesheet('selectEmpty', 'http://' + window.location.host + '/assets/select/empty.png', 100, 100, 6);

    //MatherShips MotherTrucker
    game.load.image('MotherTrucker', 'http://' + window.location.host + '/assets/MotherTrucker.png');
    game.load.image('FuryRoad', 'http://' + window.location.host + '/assets/FuryRoad.png');

    // Units
    game.load.image('tank', 'http://' + window.location.host + '/assets/tank.png');
    game.load.spritesheet('tank360', 'http://' + window.location.host + '/assets/tank360.png', 100, 100, 360);


    // Structures
    game.load.image('respawn', 'http://' + window.location.host + '/assets/respawn.png');

    // Map Objects
    game.load.image('wall', 'http://' + window.location.host + '/assets/obstacle.png');
    game.load.image('floor', 'http://' + window.location.host + '/assets/openCell.jpg');
    game.load.image('terrain_1', 'http://' + window.location.host + '/assets/tree1.png');
    game.load.image('terrain_2', 'http://' + window.location.host + '/assets/tree2.png');

    // fog
    game.load.image('FogOfWar', 'http://' + window.location.host + '/assets/fogOfWar.png');

    // Interface marks
    game.load.image('MarkMoveLastCell', 'http://' + window.location.host + '/assets/toMove.png');

    // Fonts
    game.load.bitmapFont('carrier_command', 'https://examples.phaser.io/assets/fonts/bitmapFonts/carrier_command.png', 'https://examples.phaser.io/assets/fonts/bitmapFonts/carrier_command.xml');

}