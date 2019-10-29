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