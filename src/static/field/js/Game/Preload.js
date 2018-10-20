function preload() {
    game.stage.disableVisibilityChange = true; // не дает оставиться выполнения скрипта если окно скрыто

    // Select
        // line
    game.load.spritesheet('linePlace', 'http://' + window.location.host + '/assets/select/place.png', 80, 100, 6);
    game.load.spritesheet('lineMove', 'http://' + window.location.host + '/assets/select/move.png', 80, 100, 6);
    game.load.spritesheet('lineTarget', 'http://' + window.location.host + '/assets/select/target.png', 80, 100, 6);
        // unit
    game.load.spritesheet('MySelectUnit', 'http://' + window.location.host + '/assets/select/mySelectUnit.png', 100, 100, 3);
    game.load.spritesheet('HostileSelectUnit', 'http://' + window.location.host + '/assets/select/hostileUnitSelect.png', 100, 100, 3);
        // zone
    game.load.spritesheet('selectEmpty', 'http://' + window.location.host + '/assets/select/empty.png', 80, 100, 6);
    game.load.spritesheet('selectTarget', 'http://' + window.location.host + '/assets/select/TargetSet.png', 80, 100, 3);
        // path
    game.load.spritesheet('pathCell', 'http://' + window.location.host + '/assets/select/pathSelect.png', 80, 100, 1);

    //Equip_Animate
    game.load.spritesheet('EnergyShield', 'http://' + window.location.host + '/assets/equipAnimate/energy_shield_animate.png', 100, 100, 20);
    game.load.spritesheet('RepairKit', 'http://' + window.location.host + '/assets/equipAnimate/repair_kit_animate.png', 100, 100, 9);
    game.load.image('labelEffects', 'http://' + window.location.host + '/assets/effects/label_effects.png');

    //MatherShips MotherTrucker
    game.load.spritesheet('Mother', 'http://' + window.location.host + '/assets/Mother360.png', 200, 200, 360);

    // Units
    game.load.spritesheet('tank', 'http://' + window.location.host + '/assets/Tank360.png', 100, 100, 360);

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

    //muzzle fire Effects
    game.load.spritesheet('fireMuzzle_1', 'http://' + window.location.host + '/assets/fire_effects/fireMuzzle_1.png', 50, 50, 3);
    game.load.spritesheet('fireMuzzle_2', 'http://' + window.location.host + '/assets/fire_effects/fireMuzzle_2.png', 50, 50, 3);

    //explosions
    game.load.spritesheet('explosion_1', 'http://' + window.location.host + '/assets/fire_effects/explosion_1.png', 100, 100, 30);

    //  Load the Google WebFont Loader script
    game.load.script('webfont', '//ajax.googleapis.com/ajax/libs/webfont/1.4.7/webfont.js');
}