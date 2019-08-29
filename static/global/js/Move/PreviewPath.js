let oldPoint = [];

function PreviewPath(jsonData) {

    while (0 < oldPoint.length) {
        let label = oldPoint.shift();
        if (label) label.destroy();
    }

    // это нужно для отрисовки пути на мине карте
    game.units[jsonData.short_unit.id].moveTo = jsonData.path[jsonData.path.length - 1];

    CreateMiniMap();

    for (let i = 0; jsonData.path && i < jsonData.path.length; i++) {
        if (i % 3 === 0 || i + 1 === jsonData.path.length) {
            let label = game.floorObjectLayer.create(jsonData.path[i].x, jsonData.path[i].y, 'pathCell');
            label.anchor.setTo(0.5);
            label.scale.set(0.25);

            let tween = game.add.tween(label).to({
                alpha: 1
            }, 100 * (i + 1), Phaser.Easing.Linear.None, true);

            tween.onComplete.add(function () {
                label.destroy();
            });

            oldPoint.push(label)
        }
    }
}