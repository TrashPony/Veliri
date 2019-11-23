function FlyBullet(jsonData) {
    let bulletState = jsonData.bullet;
    let bullet = game.bullets[jsonData.bullet.uuid];
    let path = jsonData.path_unit;

    if (!bullet) {
        bullet = CreateBullet(bulletState)
    }

    game.add.tween(bullet).to(
        {x: path.x, y: path.y},
        path.millisecond,
        Phaser.Easing.Linear.None, true, 0
    );

    game.add.tween(bullet.shadow).to(
        {
            x: path.x + GetOffsetBulletShadow(bulletState.z).x,
            y: path.y + GetOffsetBulletShadow(bulletState.z).y
        },
        path.millisecond,
        Phaser.Easing.Linear.None, true, 0
    );

    ShortDirectionRotateTween(bullet, Phaser.Math.degToRad(path.rotate), path.millisecond);
    ShortDirectionRotateTween(bullet.shadow, Phaser.Math.degToRad(path.rotate), path.millisecond);

    game.add.tween(bullet.scale).to(GetBulletSize(bulletState.z), path.millisecond, Phaser.Easing.Linear.None, true, 0);
    game.add.tween(bullet.shadow.scale).to(GetBulletSize(bulletState.z), path.millisecond, Phaser.Easing.Linear.None, true, 0);
}

function GetBulletSize(z) {
    if (z > 1) {
        return {
            x: 0.2 + (z - 1) / 5,
            y: 0.2 + (z - 1) / 5,
        }
    } else {
        return {
            x: 0.2,
            y: 0.2,
        }
    }
}

function GetOffsetBulletShadow(z) {
    if (z > 1) {
        return {
            x: game.shadowXOffset * (z + ((z - 1) * 5)),
            y: game.shadowYOffset * (z + ((z - 1) * 5)),
        }
    } else {
        return {
            x: game.shadowXOffset * z,
            y: game.shadowYOffset * z,
        }
    }
}

function FlyLaser(jsonData) {
    // лазер это тупо лучь который существуе и пропадает через время
    let fakeBullet = game.bulletLayer.create(jsonData.path_unit.x, jsonData.path_unit.y, "piu-piu", 0);
    fakeBullet.anchor.setTo(0.5);
    fakeBullet.alpha = 0;

    let fakeBulletEnd = game.bulletLayer.create(jsonData.path_unit.x, jsonData.path_unit.y, "piu-piu", 0);
    fakeBulletEnd.anchor.setTo(0.5);
    fakeBulletEnd.alpha = 0;

    fakeBullet.fakeBulletEnd = fakeBulletEnd;

    game.add.tween(fakeBullet).to({x: jsonData.bullet.target.x, y: jsonData.bullet.target.y},
        jsonData.path_unit.millisecond, Phaser.Easing.Linear.None, true, 0);

    setTimeout(function () {
        let end = game.add.tween(fakeBulletEnd).to({x: jsonData.bullet.target.x, y: jsonData.bullet.target.y},
            jsonData.path_unit.millisecond, Phaser.Easing.Linear.None, true, 0);

        end.onComplete.add(function () {
            fakeBullet.fakeBulletEnd.destroy();
            fakeBullet.destroy();
            fakeBullet = null;
        })
    }, 100);

    let laserOut = game.add.graphics(0, 0);
    laserOut.lineStyle(3, 0x10EDFF, 1);

    let laserIn = game.add.graphics(0, 0);
    laserIn.lineStyle(1, 0xFFFFFF, 1);

    let updateLaser = setInterval(function () {

        if (!fakeBullet) {
            clearInterval(updateLaser);
            laserOut.destroy();
            laserIn.destroy();
            return;
        }

        laserOut.clear();
        laserOut.lineStyle(3, 0x10EDFF, 1);
        laserOut.moveTo(fakeBullet.x, fakeBullet.y);
        laserOut.lineTo(fakeBullet.fakeBulletEnd.x, fakeBullet.fakeBulletEnd.y);

        laserIn.clear();
        laserIn.lineStyle(1, 0xFFFFFF, 1);
        laserIn.moveTo(fakeBullet.x, fakeBullet.y);
        laserIn.lineTo(fakeBullet.fakeBulletEnd.x, fakeBullet.fakeBulletEnd.y);
    }, 10)
}