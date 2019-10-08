function BoxMove(path, boxId) {
    let box = null;

    for (let i of game.boxes) {
        if (i.id === boxId) {
            box = i;
        }
    }

    if (box) {
        game.add.tween(box.sprite).to({
                x: path.x,
                y: path.y
            }, path.millisecond, Phaser.Easing.Linear.None, true, 0
        );

        game.add.tween(box.sprite.shadow).to({
                x: path.x + game.shadowXOffset,
                y: path.y + game.shadowYOffset
            }, path.millisecond, Phaser.Easing.Linear.None, true, 0
        );

        ShortDirectionRotateTween(box.sprite, Phaser.Math.degToRad(path.rotate), path.millisecond);
        ShortDirectionRotateTween(box.sprite.shadow, Phaser.Math.degToRad(path.rotate), path.millisecond);
    }
}