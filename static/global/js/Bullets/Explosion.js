function ExplosionBullet(jsonData) {
    // TODO проигрываем взрыв в точке снаряда
    //   появляется кратер который прилетает с бека, если не прилетает то нет
    let bullet = game.bullets[jsonData.bullet.uuid];
    if (bullet) {
        deleteBullet(bullet);
    }
}

function deleteBullet(bullet) {
    if (bullet.shadow) {
        bullet.shadow.destroy();
    }
    bullet.destroy();
    delete game.bullets[bullet.uuid];
}