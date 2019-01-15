function UseDigger(jsonData) {
    console.log(jsonData)

    game.floorObjectSelectLineLayer.forEach(function (sprite) {
        sprite.visible = false;
    });

    if (jsonData.other_user.squad_id === game.squad.id) {

        if (!game.squad.equipDrons) {
            game.squad.equipDrons = [];
        }

        let equipSlot = GetEquip(jsonData.type_slot, jsonData.slot);
        let equipDrone = LaunchDrone(equipSlot.equip.name, game.squad);

        let xy = GetXYCenterHex(jsonData.q, jsonData.r);
        game.squad.equipDrons.push({
            drone: equipDrone,
            xy: xy,
            id: "miningEquip" + jsonData.type_slot + "" + jsonData.slot,
            toSquad: false,
            move: false,
        });

        let progressBar = document.getElementById("miningEquip" + jsonData.type_slot + jsonData.slot);
        progressBar.style.animation = "none";
        setTimeout(function () {
            progressBar.style.animation = "reload " + equipSlot.equip.reload + "s linear 1";
        }, 10);
    } else {
        for (let i = 0; game.otherUsers && i < game.otherUsers.length; i++) {
            if (game.otherUsers[i].user_name === jsonData.other_user.user_name) {
                if (!game.otherUsers[i].equipDrons) {
                    game.otherUsers[i].equipDrons = [];
                }

                let equipDrone = LaunchDrone(jsonData.name, game.otherUsers[i]);
                let xy = GetXYCenterHex(jsonData.q, jsonData.r);
                game.otherUsers[i].equipDrons.push({
                    drone: equipDrone,
                    xy: xy,
                    id: "miningEquip" + jsonData.type_slot + "" + jsonData.slot,
                    toSquad: false,
                    move: false,
                });
            }
        }
    }

    if (jsonData.dynamic_object && jsonData.dynamic_object.texture_object !== '') {
        setTimeout(function () {
            CreateDynamicObject(jsonData)
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
    shadowDrone.scale.set(0.10);
    shadowDrone.anchor.setTo(0.5);
    shadowDrone.alpha = 0;
    shadowDrone.tint = 0x000000;
    game.physics.enable(shadowDrone, Phaser.Physics.ARCADE);

    let equipDrone = game.flyObjectsLayer.create(squad.sprite.x, squad.sprite.y, name);
    equipDrone.scale.set(0.10);
    equipDrone.anchor.setTo(0.5);
    equipDrone.alpha = 0;
    game.physics.enable(equipDrone, Phaser.Physics.ARCADE);


    equipDrone.shadow = shadowDrone;

    game.add.tween(equipDrone.shadow).to({angle: 360}, 3000, Phaser.Easing.Linear.None, true, 0, 0, false).loop(true);
    game.add.tween(equipDrone).to({angle: 360}, 3000, Phaser.Easing.Linear.None, true, 0, 0, false).loop(true);

    game.add.tween(equipDrone.shadow).to({alpha: 0.3}, 700, Phaser.Easing.Linear.None, true, 0);
    game.add.tween(equipDrone).to({alpha: 1}, 700, Phaser.Easing.Linear.None, true, 0);

    game.add.tween(equipDrone.shadow.scale).to({x: 0.2, y: 0.2}, 700, Phaser.Easing.Linear.None, true, 0);
    game.add.tween(equipDrone.scale).to({x: 0.2, y: 0.2}, 700, Phaser.Easing.Linear.None, true, 0);

    game.add.tween(equipDrone.shadow).to({
        x: squad.sprite.x + game.shadowXOffset * 5,
        y: squad.sprite.y + game.shadowYOffset * 5
    }, 700, Phaser.Easing.Linear.None, true, 0);

    return equipDrone
}