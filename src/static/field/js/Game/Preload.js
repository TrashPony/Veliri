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
    game.load.image('selectTarget_4', 'http://' + window.location.host + '/assets/select/target4.png');
    game.load.image('selectTarget_5', 'http://' + window.location.host + '/assets/select/target5.png');

    game.load.spritesheet('MySelectUnit', 'http://' + window.location.host + '/assets/select/mySelectUnit.png', 100, 100, 3);
    game.load.spritesheet('HostileSelectUnit', 'http://' + window.location.host + '/assets/select/hostileUnitSelect.png', 100, 100, 3);

    game.load.spritesheet('selectEmpty', 'http://' + window.location.host + '/assets/select/empty.png', 100, 100, 6);
    game.load.spritesheet('selectTarget', 'http://' + window.location.host + '/assets/select/TargetSet.png', 100, 100, 6);

    //Equip_Animate
    game.load.spritesheet('EnergyShield', 'http://' + window.location.host + '/assets/equipAnimate/energy_shield_animate.png', 100, 100, 20);
    game.load.spritesheet('RepairKit', 'http://' + window.location.host + '/assets/equipAnimate/repair_kit_animate.png', 100, 100, 9);

    //MatherShips MotherTrucker
    game.load.image('MotherTrucker', 'http://' + window.location.host + '/assets/MotherTrucker.png');
    game.load.image('FuryRoad', 'http://' + window.location.host + '/assets/FuryRoad.png');

    // Units
    game.load.image('tank', 'http://' + window.location.host + '/assets/tank.png');
    game.load.spritesheet('tank360', 'http://' + window.location.host + '/assets/tank360.png', 100, 100, 360);

    // Map Objects
    game.load.image('terrain_1', 'http://' + window.location.host + '/assets/map/tree1.png');
    game.load.image('terrain_2', 'http://' + window.location.host + '/assets/map/tree2.png');
    game.load.image('terrain_3', 'http://' + window.location.host + '/assets/map/tree3.png');
    game.load.image('crater', 'http://' + window.location.host + '/assets/map/crater.png');

    game.load.image('desert1', 'http://' + window.location.host + '/assets/map/desert/desert_1.jpg');
    game.load.image('desert2', 'http://' + window.location.host + '/assets/map/desert/desert_2.jpg');
    game.load.image('desert3', 'http://' + window.location.host + '/assets/map/desert/desert_3.jpg');
    game.load.image('desert4', 'http://' + window.location.host + '/assets/map/desert/desert_4.jpg');
    game.load.image('desert5', 'http://' + window.location.host + '/assets/map/desert/desert_5.jpg');
    game.load.image('desert6', 'http://' + window.location.host + '/assets/map/desert/desert_6.jpg');
    game.load.image('desert7', 'http://' + window.location.host + '/assets/map/desert/desert_7.jpg');
    game.load.image('desert8', 'http://' + window.location.host + '/assets/map/desert/desert_8.jpg');
    game.load.image('desert9', 'http://' + window.location.host + '/assets/map/desert/desert_9.jpg');

    game.load.image('sand_stone_04', 'http://' + window.location.host + '/assets/map/sand_stone_04.png');
    game.load.image('sand_stone_05', 'http://' + window.location.host + '/assets/map/sand_stone_05.png');
    game.load.image('sand_stone_06', 'http://' + window.location.host + '/assets/map/sand_stone_06.png');
    game.load.image('sand_stone_07', 'http://' + window.location.host + '/assets/map/sand_stone_07.png');
    game.load.image('sand_stone_08', 'http://' + window.location.host + '/assets/map/sand_stone_08.png');

    // fog
    game.load.image('FogOfWar', 'http://' + window.location.host + '/assets/map/fogOfWar.png');

    // Interface marks
    game.load.image('MarkMoveLastCell', 'http://' + window.location.host + '/assets/select/toMove.png');
    game.load.image('MarkTarget', 'http://' + window.location.host + '/assets/select/aim.png');

    // Fonts
    game.load.bitmapFont('carrier_command', 'https://examples.phaser.io/assets/fonts/bitmapFonts/carrier_command.png', 'https://examples.phaser.io/assets/fonts/bitmapFonts/carrier_command.xml');

    //Bar
    game.load.image('healBar', 'http://' + window.location.host + '/assets/bar/healBar.png');
    game.load.image('heal', 'http://' + window.location.host + '/assets/bar/heal.png');

}