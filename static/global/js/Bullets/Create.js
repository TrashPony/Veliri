function CreateBullet(bulletState) {

    let group = game.bulletLayer;
    if (bulletState.artillery) {
        group = game.artilleryLayer
    }

    let shadowBullet = group.create(
        bulletState.x + GetOffsetBulletShadow(bulletState.z).x,
        bulletState.y + GetOffsetBulletShadow(bulletState.z).y,
        bulletState.ammo.name);

    shadowBullet.anchor.setTo(0.5, 0.5);
    shadowBullet.scale.setTo(GetBulletSize(bulletState.z).x);
    shadowBullet.angle = bulletState.rotate;
    shadowBullet.tint = 0x000000;
    shadowBullet.alpha = 0.3;

    let bullet = group.create(bulletState.x, bulletState.y, bulletState.ammo.name);
    bullet.anchor.setTo(0.5, 0.5);
    bullet.scale.setTo(GetBulletSize(bulletState.z).x);
    bullet.angle = bulletState.rotate;

    bullet.shadow = shadowBullet;
    game.bullets[bulletState.uuid] = bullet;

    return bullet
}