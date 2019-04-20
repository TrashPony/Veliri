function preload() {
    game.load.script('BlurX', 'https://cdn.rawgit.com/photonstorm/phaser-ce/master/filters/BlurX.js');
    game.load.script('BlurY', 'https://cdn.rawgit.com/photonstorm/phaser-ce/master/filters/BlurY.js');
    // Select
    // line
    game.load.spritesheet('linePlace', 'http://' + window.location.host + '/assets/select/place.png', 80, 100, 6);
    game.load.spritesheet('lineMove', 'http://' + window.location.host + '/assets/select/move.png', 80, 100, 6);
    game.load.spritesheet('lineTarget', 'http://' + window.location.host + '/assets/select/target.png', 80, 100, 6);
    game.load.spritesheet('lineGameZone', 'http://' + window.location.host + '/assets/select/lineGameZone.png', 80, 100, 6);
    // unit
    game.load.spritesheet('MySelectUnit', 'http://' + window.location.host + '/assets/select/mySelectUnit.png', 100, 100, 3);
    game.load.spritesheet('HostileSelectUnit', 'http://' + window.location.host + '/assets/select/hostileUnitSelect.png', 100, 100, 3);
    // zone
    game.load.spritesheet('selectEmpty', 'http://' + window.location.host + '/assets/select/empty.png', 80, 100, 6);
    game.load.spritesheet('mapEditor', 'http://' + window.location.host + '/assets/select/mapEditor.png', 80, 100, 6);

    game.load.spritesheet('selectTarget', 'http://' + window.location.host + '/assets/select/TargetSet.png', 80, 100, 3);
    // path
    game.load.spritesheet('pathCell', 'http://' + window.location.host + '/assets/select/pathSelect.png', 80, 100, 1);

    //Equip_Animate
    game.load.spritesheet('ReloadEquip', 'http://' + window.location.host + '/assets/reload.png', 256, 256, 30);
    game.load.spritesheet('EnergyShield', 'http://' + window.location.host + '/assets/equipAnimate/energy_shield_animate.png', 100, 100, 20);
    game.load.spritesheet('RepairKit', 'http://' + window.location.host + '/assets/equipAnimate/repair_kit_animate.png', 100, 100, 9);
    game.load.image('labelEffects', 'http://' + window.location.host + '/assets/effects/label_effects.png');

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

    // equips
    game.load.image('digger', 'http://' + window.location.host + '/assets/units/equip/digger.png');
    game.load.image('extracted', 'http://' + window.location.host + '/assets/units/equip/extracted.png');

    // Effects
    game.load.spritesheet('Smoke45Frames', 'http://' + window.location.host + '/assets/fire_effects/Smoke45Frames.png', 256, 256, 45);

    // Map Terrains
    game.load.image('desert', 'http://' + window.location.host + '/assets/map/terrain/desert.png');
    game.load.image('desert_2', 'http://' + window.location.host + '/assets/map/terrain/desert_2.png');

    // Map Objects
    game.load.image('crater', 'http://' + window.location.host + '/assets/map/objects/crater.png');
    game.load.image('crater_2', 'http://' + window.location.host + '/assets/map/objects/crater_2.png');
    game.load.image('fallen_01', 'http://' + window.location.host + '/assets/map/objects/fallen_01.png');
    game.load.image('fallen_02', 'http://' + window.location.host + '/assets/map/objects/fallen_02.png');

    game.load.image('evacuation', 'http://' + window.location.host + '/assets/map/objects/evacuation.png');

    // Boxes
    game.load.image('box', 'http://' + window.location.host + '/assets/map/boxes/box.png');
    game.load.image('secure_underground_box', 'http://' + window.location.host + '/assets/map/boxes/secure_underground_box.png');
    game.load.image('underground_box', 'http://' + window.location.host + '/assets/map/boxes/underground_box.png');

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
    game.load.image('baseIcon', 'http://' + window.location.host + '/assets/buttons/baseIcon.png');
    game.load.image('baseResp', 'http://' + window.location.host + '/assets/buttons/baseResp.png');

    // Buttons
    game.load.image('buttonRotate', 'http://' + window.location.host + '/assets/buttons/rotate.png');

    // Reservoir
    game.load.image('copper_ore_low', 'http://' + window.location.host + '/assets/resource/map_sprites/copper_ore/low.png');
    game.load.image('copper_ore_middle', 'http://' + window.location.host + '/assets/resource/map_sprites/copper_ore/middle.png');
    game.load.image('copper_ore_full', 'http://' + window.location.host + '/assets/resource/map_sprites/copper_ore/full.png');

    game.load.image('oil', 'http://' + window.location.host + '/assets/resource/map_sprites/oil.png');

    game.load.image('silicon_ore_low', 'http://' + window.location.host + '/assets/resource/map_sprites/silicon_ore/low.png');
    game.load.image('silicon_ore_middle', 'http://' + window.location.host + '/assets/resource/map_sprites/silicon_ore/middle.png');
    game.load.image('silicon_ore_full', 'http://' + window.location.host + '/assets/resource/map_sprites/silicon_ore/full.png');

    game.load.image('iron_ore_low', 'http://' + window.location.host + '/assets/resource/map_sprites/iron_ore/low.png');
    game.load.image('iron_ore_middle', 'http://' + window.location.host + '/assets/resource/map_sprites/iron_ore/middle.png');
    game.load.image('iron_ore_full', 'http://' + window.location.host + '/assets/resource/map_sprites/iron_ore/full.png');

    game.load.image('thorium_ore_low', 'http://' + window.location.host + '/assets/resource/map_sprites/thorium_ore/low.png');
    game.load.image('thorium_ore_middle', 'http://' + window.location.host + '/assets/resource/map_sprites/thorium_ore/middle.png');
    game.load.image('thorium_ore_full', 'http://' + window.location.host + '/assets/resource/map_sprites/thorium_ore/full.png');

    game.load.image('titanium_ore_low', 'http://' + window.location.host + '/assets/resource/map_sprites/titanium_ore/low.png');
    game.load.image('titanium_ore_middle', 'http://' + window.location.host + '/assets/resource/map_sprites/titanium_ore/middle.png');
    game.load.image('titanium_ore_full', 'http://' + window.location.host + '/assets/resource/map_sprites/titanium_ore/full.png');

    game.load.image('infoAnomaly', 'http://' + window.location.host + '/assets/info.png');

    //Brush
    game.load.image('brush', 'http://' + window.location.host + '/assets/terrainTextures/brush.png');
    game.load.image('desertDunes', 'http://' + window.location.host + '/assets/terrainTextures/desertDunes.jpg');
    game.load.image('desertDunes_2', 'http://' + window.location.host + '/assets/terrainTextures/desertDunes_2.jpg');
    game.load.image('xenos', 'http://' + window.location.host + '/assets/terrainTextures/xenos.jpg');
    game.load.image('xenos_2', 'http://' + window.location.host + '/assets/terrainTextures/xenos_2.jpg');
    game.load.image('arctic', 'http://' + window.location.host + '/assets/terrainTextures/arctic.jpg');
    game.load.image('arctic_2', 'http://' + window.location.host + '/assets/terrainTextures/arctic_2.jpg');
    game.load.image('tundra', 'http://' + window.location.host + '/assets/terrainTextures/tundra.jpg');
    game.load.image('tundra_2', 'http://' + window.location.host + '/assets/terrainTextures/tundra_2.jpg');
    game.load.image('grass_1', 'http://' + window.location.host + '/assets/terrainTextures/grass_1.jpg');
    game.load.image('grass_2', 'http://' + window.location.host + '/assets/terrainTextures/grass_2.jpg');
    game.load.image('grass_3', 'http://' + window.location.host + '/assets/terrainTextures/grass_3.jpg');
    game.load.image('soil', 'http://' + window.location.host + '/assets/terrainTextures/soil.jpg');
    game.load.image('soil_2', 'http://' + window.location.host + '/assets/terrainTextures/soil_2.jpg');
    game.load.image('ravine_1_9', 'http://' + window.location.host + '/assets/terrainTextures/ravine_1_9.png');
    game.load.image('ravine_2_9', 'http://' + window.location.host + '/assets/terrainTextures/ravine_2_9.png');

    // Mountains
    game.load.image('mountain_1', 'http://' + window.location.host + '/assets/map/objects/mountains/mountain_1.png');
    game.load.image('mountain_1_2', 'http://' + window.location.host + '/assets/map/objects/mountains/mountain_1_2.png');
    game.load.image('mountain_2', 'http://' + window.location.host + '/assets/map/objects/mountains/mountain_2.png');
    game.load.image('mountain_2_2', 'http://' + window.location.host + '/assets/map/objects/mountains/mountain_2_2.png');
    game.load.image('mountain_3', 'http://' + window.location.host + '/assets/map/objects/mountains/mountain_3.png');
    game.load.image('mountain_3_2', 'http://' + window.location.host + '/assets/map/objects/mountains/mountain_3_2.png');
    game.load.image('mountain_4', 'http://' + window.location.host + '/assets/map/objects/mountains/mountain_4.png');
    game.load.image('mountain_4_2', 'http://' + window.location.host + '/assets/map/objects/mountains/mountain_4_2.png');
    game.load.image('mountain_5', 'http://' + window.location.host + '/assets/map/objects/mountains/mountain_5.png');
    game.load.image('mountain_5_2', 'http://' + window.location.host + '/assets/map/objects/mountains/mountain_5_2.png');
    game.load.image('mountain_6', 'http://' + window.location.host + '/assets/map/objects/mountains/mountain_6.png');
    game.load.image('mountain_6_2', 'http://' + window.location.host + '/assets/map/objects/mountains/mountain_6_2.png');
    game.load.image('mountain_7', 'http://' + window.location.host + '/assets/map/objects/mountains/mountain_7.png');
    game.load.image('mountain_7_2', 'http://' + window.location.host + '/assets/map/objects/mountains/mountain_7_2.png');
    game.load.image('mountain_8', 'http://' + window.location.host + '/assets/map/objects/mountains/mountain_8.png');

    // Ravine
    game.load.image('ravine_1_1', 'http://' + window.location.host + '/assets/map/objects/ravines/ravine_1_1.png');
    game.load.image('ravine_1_2', 'http://' + window.location.host + '/assets/map/objects/ravines/ravine_1_2.png');
    game.load.image('ravine_1_3', 'http://' + window.location.host + '/assets/map/objects/ravines/ravine_1_3.png');
    game.load.image('ravine_1_4', 'http://' + window.location.host + '/assets/map/objects/ravines/ravine_1_4.png');
    game.load.image('ravine_1_5', 'http://' + window.location.host + '/assets/map/objects/ravines/ravine_1_5.png');
    game.load.image('ravine_1_6', 'http://' + window.location.host + '/assets/map/objects/ravines/ravine_1_6.png');
    game.load.image('ravine_1_7', 'http://' + window.location.host + '/assets/map/objects/ravines/ravine_1_7.png');
    game.load.image('ravine_1_8', 'http://' + window.location.host + '/assets/map/objects/ravines/ravine_1_8.png');
    game.load.image('ravine_1_9', 'http://' + window.location.host + '/assets/map/objects/ravines/ravine_1_9.png');
    game.load.image('ravine_1_10', 'http://' + window.location.host + '/assets/map/objects/ravines/ravine_1_10.png');
    game.load.image('ravine_1_Full', 'http://' + window.location.host + '/assets/map/objects/ravines/ravine_1_Full.png');

    game.load.image('ravine_2_1', 'http://' + window.location.host + '/assets/map/objects/ravines/ravine_2_1.png');
    game.load.image('ravine_2_2', 'http://' + window.location.host + '/assets/map/objects/ravines/ravine_2_2.png');
    game.load.image('ravine_2_3', 'http://' + window.location.host + '/assets/map/objects/ravines/ravine_2_3.png');
    game.load.image('ravine_2_4', 'http://' + window.location.host + '/assets/map/objects/ravines/ravine_2_4.png');
    game.load.image('ravine_2_5', 'http://' + window.location.host + '/assets/map/objects/ravines/ravine_2_5.png');
    game.load.image('ravine_2_6', 'http://' + window.location.host + '/assets/map/objects/ravines/ravine_2_6.png');
    game.load.image('ravine_2_7', 'http://' + window.location.host + '/assets/map/objects/ravines/ravine_2_7.png');
    game.load.image('ravine_2_8', 'http://' + window.location.host + '/assets/map/objects/ravines/ravine_2_8.png');
    game.load.image('ravine_2_9', 'http://' + window.location.host + '/assets/terrainTextures/ravine_2_9.png');

    // Plants
    game.load.image('plant_1', 'http://' + window.location.host + '/assets/map/objects/plants/plant_1.png');
    game.load.image('plant_2', 'http://' + window.location.host + '/assets/map/objects/plants/plant_2.png');
    game.load.image('plant_2_2', 'http://' + window.location.host + '/assets/map/objects/plants/plant_2_2.png');
    game.load.image('plant_2_3', 'http://' + window.location.host + '/assets/map/objects/plants/plant_2_3.png');
    game.load.image('plant_2_4', 'http://' + window.location.host + '/assets/map/objects/plants/plant_2_4.png');
    game.load.image('plant_2_5', 'http://' + window.location.host + '/assets/map/objects/plants/plant_2_5.png');
    game.load.image('plant_2_6', 'http://' + window.location.host + '/assets/map/objects/plants/plant_2_6.png');
    game.load.image('plant_3', 'http://' + window.location.host + '/assets/map/objects/plants/plant_3.png');
    game.load.image('plant_3_2', 'http://' + window.location.host + '/assets/map/objects/plants/plant_3_2.png');
    game.load.image('plant_3_3', 'http://' + window.location.host + '/assets/map/objects/plants/plant_3_3.png');
    game.load.image('plant_4', 'http://' + window.location.host + '/assets/map/objects/plants/plant_4.png');
    game.load.image('plant_5', 'http://' + window.location.host + '/assets/map/objects/plants/plant_5.png');
    game.load.image('plant_5_2', 'http://' + window.location.host + '/assets/map/objects/plants/plant_5_2.png');

    // roads
    game.load.image('road_1_1', 'http://' + window.location.host + '/assets/map/objects/roads/road_1_1.png');
    game.load.image('road_1_2', 'http://' + window.location.host + '/assets/map/objects/roads/road_1_2.png');
    game.load.image('road_1_3', 'http://' + window.location.host + '/assets/map/objects/roads/road_1_3.png');
    game.load.image('road_1_4', 'http://' + window.location.host + '/assets/map/objects/roads/road_1_4.png');
    game.load.image('road_1_5', 'http://' + window.location.host + '/assets/map/objects/roads/road_1_5.png');
    game.load.image('road_1_6', 'http://' + window.location.host + '/assets/map/objects/roads/road_1_6.png');

    game.load.image('road_2_1', 'http://' + window.location.host + '/assets/map/objects/roads/road_2_1.png');
    game.load.image('road_2_2', 'http://' + window.location.host + '/assets/map/objects/roads/road_2_2.png');
    game.load.image('road_2_3', 'http://' + window.location.host + '/assets/map/objects/roads/road_2_3.png');
    game.load.image('road_2_4', 'http://' + window.location.host + '/assets/map/objects/roads/road_2_4.png');
    game.load.image('road_2_5', 'http://' + window.location.host + '/assets/map/objects/roads/road_2_5.png');
    game.load.image('road_2_6', 'http://' + window.location.host + '/assets/map/objects/roads/road_2_6.png');
    game.load.image('road_2_7', 'http://' + window.location.host + '/assets/map/objects/roads/road_2_7.png');

    //Clouds
    game.load.image('cloud0', 'http://' + window.location.host + '/assets/map/clouds/cloud13.png');
    game.load.image('cloud1', 'http://' + window.location.host + '/assets/map/clouds/cloud1.png');
    game.load.image('cloud2', 'http://' + window.location.host + '/assets/map/clouds/cloud2.png');
    game.load.image('cloud3', 'http://' + window.location.host + '/assets/map/clouds/cloud3.png');
    game.load.image('cloud4', 'http://' + window.location.host + '/assets/map/clouds/cloud4.png');
    game.load.image('cloud5', 'http://' + window.location.host + '/assets/map/clouds/cloud5.png');
    game.load.image('cloud6', 'http://' + window.location.host + '/assets/map/clouds/cloud6.png');
    game.load.image('cloud7', 'http://' + window.location.host + '/assets/map/clouds/cloud7.png');
    game.load.image('cloud8', 'http://' + window.location.host + '/assets/map/clouds/cloud8.png');
    game.load.image('cloud9', 'http://' + window.location.host + '/assets/map/clouds/cloud9.png');
    game.load.image('cloud10', 'http://' + window.location.host + '/assets/map/clouds/cloud10.png');
    game.load.image('cloud11', 'http://' + window.location.host + '/assets/map/clouds/cloud11.png');
    game.load.image('cloud12', 'http://' + window.location.host + '/assets/map/clouds/cloud12.png');

    // Icons
    game.load.image('transportIcon', 'http://' + window.location.host + '/assets/icons/transport.png');
    game.load.image('baseInIcon', 'http://' + window.location.host + '/assets/icons/baseIcon.png');
    game.load.image('sectorOutIcon', 'http://' + window.location.host + '/assets/icons/sectorOutIcon.png');
    game.load.image('cancelIcon', 'http://' + window.location.host + '/assets/icons/cancelIcon.png');
    game.load.image('ammoCoordinate', 'http://' + window.location.host + '/assets/icons/ammoCoordinate.png');

    //Bodies
    game.load.image('MasherShip_1', 'http://' + window.location.host + '/assets/units/body/MasherShip_1.png');
    game.load.image('MasherShip_1_bottom', 'http://' + window.location.host + '/assets/units/body/MasherShip_1_bottom.png');

    game.load.image('MasherShip_2', 'http://' + window.location.host + '/assets/units/body/MasherShip_2.png');
    game.load.image('MasherShip_2_bottom', 'http://' + window.location.host + '/assets/units/body/MasherShip_2_bottom.png');

    game.load.image('MasherShip_3', 'http://' + window.location.host + '/assets/units/body/MasherShip_3.png');
    game.load.image('MasherShip_3_bottom', 'http://' + window.location.host + '/assets/units/body/MasherShip_3_bottom.png');

    // Units
    game.load.image('heavy_tank', 'http://' + window.location.host + '/assets/units/body/heavy_tank.png');
    game.load.image('medium_tank', 'http://' + window.location.host + '/assets/units/body/medium_tank.png');
    game.load.image('light_tank', 'http://' + window.location.host + '/assets/units/body/light_tank.png');

    autoPreload();
}