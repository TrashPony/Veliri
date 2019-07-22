function ShortDirectionRotateTween(sprite, desiredRotation, time) {
    // эта функция ищет оптимальный угол для поворота
    let shortestAngle = getShortestAngle(Phaser.Math.radToDeg(desiredRotation), sprite.angle);
    let newAngle = sprite.angle + shortestAngle;
    return game.add.tween(sprite).to({
        'angle': newAngle
    }, time, Phaser.Easing.Linear.None, true);
}

function getShortestAngle(angle1, angle2) {
    // библиотечная функция не подходит
    let difference = angle2 - angle1;
    let times = Math.floor((difference - (-180)) / 360);
    return (difference - (times * 360)) * -1;
}