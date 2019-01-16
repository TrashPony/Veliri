function AnimateDigger() {
    if (game.squad.equipDrons && game.squad.equipDrons.length > 0) {
        for (let i in game.squad.equipDrons) {
            if (game.squad.equipDrons[i] && game.squad.equipDrons[i].drone.alive) {
                FlyDrone(game.squad.equipDrons[i]);
                if (game.squad.equipDrons[i].toSquad) {
                    BackDrone(game.squad.equipDrons[i], game.squad)
                }
            }
        }
    }

    for (let i = 0; game.otherUsers && i < game.otherUsers.length; i++) {
        if (game.otherUsers[i].equipDrons && game.otherUsers[i].equipDrons.length > 0) {
            for (let j in game.otherUsers[i].equipDrons) {
                if (game.otherUsers[i].equipDrons[j]) {
                    FlyDrone(game.otherUsers[i].equipDrons[j]);
                    if (game.otherUsers[i].equipDrons[j].toSquad) {
                        BackDrone(game.otherUsers[i].equipDrons[j], game.otherUsers[i]);
                    }
                }
            }
        }
    }
}

function BackDrone(droneSetting, squad) {
    let dist = game.physics.arcade.distanceToXY(droneSetting.drone, squad.sprite.x, squad.sprite.y);

    if (dist < 10) {
        droneSetting.drone.destroy();
        droneSetting.drone.shadow.destroy();
    }
    if (dist > 100) {
        game.physics.arcade.moveToXY(
            droneSetting.drone,
            squad.sprite.x,
            squad.sprite.y,
            100);
        game.physics.arcade.moveToXY(
            droneSetting.drone.shadow,
            squad.sprite.x + game.shadowXOffset * 5,
            squad.sprite.y + game.shadowYOffset * 5,
            100);
    } else {
        game.add.tween(droneSetting.drone.shadow).to({alpha: 0}, 300, Phaser.Easing.Linear.None, true, 0);
        game.add.tween(droneSetting.drone).to({alpha: 0}, 700, Phaser.Easing.Linear.None, true, 0);
        game.add.tween(droneSetting.drone.shadow.scale).to({x: 0.1, y: 0.1}, 700, Phaser.Easing.Linear.None, true, 0);
        game.add.tween(droneSetting.drone.scale).to({x: 0.1, y: 0.1}, 700, Phaser.Easing.Linear.None, true, 0);
        game.add.tween(droneSetting.drone.shadow).to({
            x: squad.sprite.x,
            y: squad.sprite.y
        }, 700, Phaser.Easing.Linear.None, true, 0);
    }
}

function FlyDrone(droneSetting) {
    if (!droneSetting.move) {
        droneSetting.move = true;
        let tween = game.add.tween(droneSetting.drone).to({
                x: droneSetting.xy.x,
                y: droneSetting.xy.y
            }, 3000, Phaser.Easing.Linear.None, true, 0
        );

        game.add.tween(droneSetting.drone.shadow).to({
                x: droneSetting.xy.x + game.shadowXOffset * 5,
                y: droneSetting.xy.y + +game.shadowYOffset * 5,
            }, 3000, Phaser.Easing.Linear.None, true, 0
        );

        tween.onComplete.add(function () {

            let laser = createLaser();
            createLineLaser(laser, droneSetting);

            let tween = game.add.tween(droneSetting.xy).to({
                    x: droneSetting.x - Math.random() * 100,
                    y: droneSetting.y + Math.random() * 100
                }, 100, Phaser.Easing.Linear.None, true, 0
            );

            let i = 0;

            let dust = game.add.emitter(droneSetting.drone.x, droneSetting.drone.y, 100);
            dust.makeParticles('smoke_puff');
            dust.minParticleScale = 0.5;
            dust.maxParticleScale = 1;
            dust.setAlpha(0.2, 0.6);
            dust.setXSpeed(-100, 100);
            dust.setYSpeed(-100, 100);
            dust.gravity = -20;
            dust.setAlpha(0, 0.5, 700, null, true);
            dust.start(false, 500, 100);
            setTimeout(function () {
                game.effectsLayer.add(dust);
            }, 100);

            let animateLaser = function () {
                i++;
                let x, y;
                let direction = Math.round(Math.random() * 4);
                if (direction <= 1) {
                    x = droneSetting.drone.x - Math.random() * 30;
                    y = droneSetting.drone.y - Math.random() * 30;
                } else if (direction === 2) {
                    x = droneSetting.drone.x - Math.random() * 30;
                    y = droneSetting.drone.y + Math.random() * 30;
                } else if (direction === 3) {
                    x = droneSetting.drone.x + Math.random() * 30;
                    y = droneSetting.drone.y + Math.random() * 30;
                } else if (direction === 4) {
                    x = droneSetting.drone.x + Math.random() * 30;
                    y = droneSetting.drone.y - Math.random() * 30;
                }

                let tween = game.add.tween(droneSetting.xy).to({
                        x: x,
                        y: y,
                    }, 20, Phaser.Easing.Linear.None, true, 0
                );

                createLineLaser(laser, droneSetting);

                if (i === 50) {
                    let object = gameObjectCreate(
                        droneSetting.drone.x,
                        droneSetting.drone.y,
                        droneSetting.spriteCrater,
                        droneSetting.scaleCrater,
                        false,
                        droneSetting.angleCrater,
                        0, 0, game.floorLayer);
                    object.alpha = 0;
                    game.add.tween(object).to({alpha: 1}, 50 * 20, Phaser.Easing.Linear.None, true, 0);
                }

                if (i < 100) {
                    tween.onComplete.add(animateLaser)
                } else {
                    laser.in.destroy();
                    laser.out.destroy();
                    dust.on = false;
                    droneSetting.toSquad = true;
                }
            };
            tween.onComplete.add(animateLaser)
        })
    }
}

function createLineLaser(laser, droneSetting) {
    laser.out.clear();
    laser.out.lineStyle(6, 0x10EDFF, 1);
    laser.out.moveTo(droneSetting.drone.x, droneSetting.drone.y);
    laser.out.lineTo(droneSetting.xy.x, droneSetting.xy.y);

    laser.in.clear();
    laser.in.lineStyle(2, 0xFFFFFF, 1);
    laser.in.moveTo(droneSetting.drone.x, droneSetting.drone.y);
    laser.in.lineTo(droneSetting.xy.x, droneSetting.xy.y);
}

function createLaser() {
    let laserOut = game.add.graphics(0, 0);
    laserOut.lineStyle(6, 0xFFEDFF, 1);

    let laserIn = game.add.graphics(0, 0);
    laserIn.lineStyle(2, 0xFFFFFF, 1);

    let blurX = game.add.filter('BlurX');
    let blurY = game.add.filter('BlurY');
    blurX.blur = 2;
    blurY.blur = 2;
    laserOut.filters = [blurX, blurY];
    blurX.blur = 1;
    blurY.blur = 1;
    laserIn.filters = [blurX, blurY];

    game.floorObjectLayer.add(laserOut);
    game.floorObjectLayer.add(laserIn);

    return {out: laserOut, in: laserIn};

}