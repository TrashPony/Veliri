function PreviewPath(jsonData) {
    for (let i = 0; i < jsonData.path.length; i++) {
        let label = game.SelectRangeLayer.create(jsonData.path[i].x, jsonData.path[i].y, 'pathCell');
        label.anchor.setTo(0.5);
        label.scale.set(0.5);

        let tween = game.add.tween(label).to({
            alpha: 1
        }, 100 * (i+1), Phaser.Easing.Linear.None, true);

        tween.onComplete.add(function () {
            label.destroy();
        })
    }
}