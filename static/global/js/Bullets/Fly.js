function FlyBullet(jsonData) {
    let bullet = game.bullets[jsonData.bullet.uuid];

    if (!bullet) {

        let shadowBullet = game.bulletLayer.create(
            jsonData.bullet.x + game.shadowXOffset * jsonData.bullet.z,
            jsonData.bullet.y + game.shadowYOffset * jsonData.bullet.z,
            jsonData.bullet.ammo.name);

        shadowBullet.anchor.setTo(0.5, 0.5);
        shadowBullet.scale.setTo(0.2);
        shadowBullet.angle = jsonData.bullet.rotate;
        shadowBullet.tint = 0x000000;
        shadowBullet.alpha = 0.3;

        bullet = game.bulletLayer.create(jsonData.bullet.x, jsonData.bullet.y, jsonData.bullet.ammo.name);
        bullet.anchor.setTo(0.5, 0.5);
        bullet.scale.setTo(0.2);
        bullet.angle = jsonData.bullet.rotate;

        bullet.shadow = shadowBullet;
        game.bullets[jsonData.bullet.uuid] = bullet;
    }

    let path = jsonData.path_unit;

    game.add.tween(bullet).to(
        {x: path.x, y: path.y},
        path.millisecond,
        Phaser.Easing.Linear.None, true, 0
    );

    game.add.tween(bullet.shadow).to(
        {
            x: path.x + game.shadowXOffset * jsonData.bullet.z,
            y: path.y + game.shadowYOffset * jsonData.bullet.z
        },
        path.millisecond,
        Phaser.Easing.Linear.None, true, 0
    );

    console.log(jsonData)
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
    }, jsonData.path_unit.millisecond);

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