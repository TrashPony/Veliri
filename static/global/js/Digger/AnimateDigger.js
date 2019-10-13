function AnimateDigger(unit) {

    if (unit.selectDiggerLine) {
        unit.selectDiggerLine.graphics.clear();
        unit.selectDiggerLine.graphics.lineStyle(3, 0x00ff00, 0.2);
        unit.selectDiggerLine.graphics.drawCircle(unit.sprite.x, unit.sprite.y, unit.selectDiggerLine.radius);
        unit.selectDiggerLine.graphics.lineStyle(1, 0x00ff00, 1);
        unit.selectDiggerLine.graphics.drawCircle(unit.sprite.x, unit.sprite.y, unit.selectDiggerLine.radius);
    }

    for (let i in unit.equipDrons) {
        if (unit.equipDrons[i] && unit.equipDrons[i].drone.alive) {
            FlyDrone(unit.equipDrons[i]);
            if (unit.equipDrons[i].toSquad) {
                BackDrone(unit.equipDrons[i], unit)
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
    if (dist > 50) {
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
        game.add.tween(droneSetting.drone.shadow.scale).to({x: 0.05, y: 0.05}, 700, Phaser.Easing.Linear.None, true, 0);
        game.add.tween(droneSetting.drone.scale).to({x: 0.05, y: 0.05}, 700, Phaser.Easing.Linear.None, true, 0);
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
            dust.minParticleScale = 0.25;
            dust.maxParticleScale = 0.5;
            dust.setAlpha(0.2, 0.6);
            dust.setXSpeed(-50, 50);
            dust.setYSpeed(-50, 50);
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
                    x = droneSetting.drone.x - Math.random() * 15;
                    y = droneSetting.drone.y - Math.random() * 15;
                } else if (direction === 2) {
                    x = droneSetting.drone.x - Math.random() * 15;
                    y = droneSetting.drone.y + Math.random() * 15;
                } else if (direction === 3) {
                    x = droneSetting.drone.x + Math.random() * 15;
                    y = droneSetting.drone.y + Math.random() * 15;
                } else if (direction === 4) {
                    x = droneSetting.drone.x + Math.random() * 15;
                    y = droneSetting.drone.y - Math.random() * 15;
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
    laser.out.lineStyle(3, 0x10EDFF, 1);
    laser.out.moveTo(droneSetting.drone.x, droneSetting.drone.y);
    laser.out.lineTo(droneSetting.xy.x, droneSetting.xy.y);

    laser.in.clear();
    laser.in.lineStyle(1, 0xFFFFFF, 1);
    laser.in.moveTo(droneSetting.drone.x, droneSetting.drone.y);
    laser.in.lineTo(droneSetting.xy.x, droneSetting.xy.y);
}

function createLaser() {
    let laserOut = game.add.graphics(0, 0);
    laserOut.lineStyle(3, 0xFFEDFF, 1);

    let laserIn = game.add.graphics(0, 0);
    laserIn.lineStyle(1, 0xFFFFFF, 1);

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