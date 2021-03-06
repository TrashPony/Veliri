function UseDigger(jsonData) {

    game.floorObjectSelectLineLayer.forEach(function (sprite) {
        sprite.visible = false;
    });

    let unit = game.units[jsonData.short_unit.id];

    if (unit) {
        if (!unit.equipDrons) {
            unit.equipDrons = [];
        }

        let equipSlot = GetEquip(jsonData.short_unit.id, jsonData.type_slot, jsonData.slot);
        let equipDrone = LaunchDrone(equipSlot.equip.name, unit);

        unit.equipDrons.push({
            drone: equipDrone,
            xy: {x: jsonData.x, y: jsonData.y},
            id: "reloadEquip" + unit.id + "" + jsonData.type_slot + "" + jsonData.slot,
            toSquad: false,
            move: false,
            spriteCrater: jsonData.dynamic_object.texture_background,
            scaleCrater: jsonData.dynamic_object.background_scale,
            angleCrater: jsonData.dynamic_object.background_rotate
        });

        let progressBar = document.getElementById("reloadEquip" + unit.id + "" + jsonData.type_slot + "" + jsonData.slot);
        if (progressBar) {
            progressBar.style.animation = "none";
            setTimeout(function () {
                progressBar.style.animation = "reload " + equipSlot.equip.reload + "s linear 1";
            }, 10);
        }
    }

    if (jsonData.dynamic_object && jsonData.dynamic_object.texture_object !== '' && jsonData.texture_background !== '') {
        setTimeout(function () {
            CreateDynamicObjects(jsonData.dynamic_object, jsonData.q, jsonData.r, false)
        }, 5000);
    }

    if (jsonData.box) {
        setTimeout(function () {
            CreateBox(jsonData.box);
        }, 5000);
    }

    if (jsonData.reservoir) {
        setTimeout(function () {
            CreateReservoir(jsonData.reservoir, jsonData.reservoir.q, jsonData.reservoir.r);
        }, 5000);
    }
}

function LaunchDrone(name, squad) {
    let shadowDrone = game.flyObjectsLayer.create(squad.sprite.x + game.shadowXOffset, squad.sprite.y + game.shadowYOffset, name);
    shadowDrone.scale.set(0.05);
    shadowDrone.anchor.setTo(0.5);
    shadowDrone.alpha = 0;
    shadowDrone.tint = 0x000000;
    game.physics.enable(shadowDrone, Phaser.Physics.ARCADE);

    let equipDrone = game.flyObjectsLayer.create(squad.sprite.x, squad.sprite.y, name);
    equipDrone.scale.set(0.05);
    equipDrone.anchor.setTo(0.5);
    equipDrone.alpha = 0;
    game.physics.enable(equipDrone, Phaser.Physics.ARCADE);


    equipDrone.shadow = shadowDrone;

    game.add.tween(equipDrone.shadow).to({angle: 360}, 3000, Phaser.Easing.Linear.None, true, 0, 0, false).loop(true);
    game.add.tween(equipDrone).to({angle: 360}, 3000, Phaser.Easing.Linear.None, true, 0, 0, false).loop(true);

    game.add.tween(equipDrone.shadow).to({alpha: 0.3}, 700, Phaser.Easing.Linear.None, true, 0);
    game.add.tween(equipDrone).to({alpha: 1}, 700, Phaser.Easing.Linear.None, true, 0);

    game.add.tween(equipDrone.shadow.scale).to({x: 0.3, y: 0.3}, 700, Phaser.Easing.Linear.None, true, 0);
    game.add.tween(equipDrone.scale).to({x: 0.3, y: 0.3}, 700, Phaser.Easing.Linear.None, true, 0);

    game.add.tween(equipDrone.shadow).to({
        x: squad.sprite.x + game.shadowXOffset * 5,
        y: squad.sprite.y + game.shadowYOffset * 5
    }, 700, Phaser.Easing.Linear.None, true, 0);

    return equipDrone
}