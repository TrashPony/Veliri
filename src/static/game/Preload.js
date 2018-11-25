function preload() {
    game.stage.disableVisibilityChange = true; // не дает оставиться выполнения скрипта если окно скрыто

    game.load.script('BlurX', 'https://cdn.rawgit.com/photonstorm/phaser-ce/master/filters/BlurX.js');
    game.load.script('BlurY', 'https://cdn.rawgit.com/photonstorm/phaser-ce/master/filters/BlurY.js');
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
    game.load.image('Mother', 'http://' + window.location.host + '/assets/units/body/Mother.png');

    // Units
        //body
    game.load.image('heavy_tank', 'http://' + window.location.host + '/assets/units/body/heavy_tank.png');
    game.load.image('medium_tank', 'http://' + window.location.host + '/assets/units/body/medium_tank.png');
    game.load.image('light_tank', 'http://' + window.location.host + '/assets/units/body/light_tank.png');
        //weapon
    game.load.image('big_missile', 'http://' + window.location.host + '/assets/units/weapon/big_missile.png');
    game.load.image('artillery', 'http://' + window.location.host + '/assets/units/weapon/artillery.png');
    game.load.image('big_laser', 'http://' + window.location.host + '/assets/units/weapon/big_laser.png');
    game.load.image('small_laser', 'http://' + window.location.host + '/assets/units/weapon/small_laser.png');
    game.load.image('small_missile', 'http://' + window.location.host + '/assets/units/weapon/small_missile.png');
    game.load.image('tank_gun', 'http://' + window.location.host + '/assets/units/weapon/tank_gun.png');
        // bullets
    game.load.spritesheet('missile_bullet', 'http://' + window.location.host + '/assets/units/gameAmmo/rocket.png', 128, 128, 40);
    game.load.image('ballistics_small_bullet', 'http://' + window.location.host + '/assets/units/gameAmmo/ballistics_small_bullet.png');
    game.load.image('ballistics_artillery_bullet', 'http://' + window.location.host + '/assets/units/gameAmmo/ballistics_artillery_bullet.png');
        // fire effects
    game.load.image('smoke_puff', 'http://' + window.location.host + '/assets/fire_effects/smoke_puff.png');
    game.load.image('fire1', 'http://' + window.location.host + '/assets/fire_effects/fire1.png');
    game.load.image('fire2', 'http://' + window.location.host + '/assets/fire_effects/fire2.png');
    game.load.image('fire3', 'http://' + window.location.host + '/assets/fire_effects/fire3.png');

    // Map Terrains
    game.load.image('desert', 'http://' + window.location.host + '/assets/map/terrain/desert.png');

    // Map Objects
    game.load.image('terrain_1', 'http://' + window.location.host + '/assets/map/objects/terrain_1.png');
    game.load.image('terrain_2', 'http://' + window.location.host + '/assets/map/objects/terrain_2.png');
    game.load.image('terrain_3', 'http://' + window.location.host + '/assets/map/objects/terrain_3.png');
    game.load.image('crater', 'http://' + window.location.host + '/assets/map/objects/crater.png');

    game.load.image('sand_stone_04', 'http://' + window.location.host + '/assets/map/objects/sand_stone_04.png');
    game.load.image('sand_stone_05', 'http://' + window.location.host + '/assets/map/objects/sand_stone_05.png');
    game.load.image('sand_stone_06', 'http://' + window.location.host + '/assets/map/objects/sand_stone_06.png');
    game.load.image('sand_stone_07', 'http://' + window.location.host + '/assets/map/objects/sand_stone_07.png');
    game.load.image('sand_stone_08', 'http://' + window.location.host + '/assets/map/objects/sand_stone_08.png');

    game.load.image('fallen_01', 'http://' + window.location.host + '/assets/map/objects/fallen_01.png');
    game.load.image('fallen_02', 'http://' + window.location.host + '/assets/map/objects/fallen_02.png');

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
    game.load.spritesheet('explosion_2', 'http://' + window.location.host + '/assets/fire_effects/explosion_2.png', 128, 128, 4);

    //  Load the Google WebFont Loader script
    game.load.script('webfont', '//ajax.googleapis.com/ajax/libs/webfont/1.4.7/webfont.js');

    // Buttons
    game.load.image('buttonPlus', 'http://' + window.location.host + '/assets/buttons/buttonPlus.png');
    game.load.image('buttonMinus', 'http://' + window.location.host + '/assets/buttons/buttonMinus.png');

    autoPreload();
}