function PreviewPath(jsonData) {
    console.log(jsonData)
    if (!game || !game.units) return;

    let unit = game.units[jsonData.short_unit.id];

    if (unit && !unit.oldPoint) {
        unit.oldPoint = []
    }

    while (unit && unit.oldPoint && 0 < unit.oldPoint.length) {
        let label = unit.oldPoint.shift();
        if (label) label.destroy();
    }

    if (!jsonData.path) {
        return
    }
    // это нужно для отрисовки пути на мине карте
    unit.moveTo = jsonData.path[jsonData.path.length - 1];

    CreateMiniMap();

    for (let i = 0; jsonData.path && i < jsonData.path.length; i++) {
        let label = game.floorObjectLayer.create(jsonData.path[i].x, jsonData.path[i].y, 'pathCell');
        label.anchor.setTo(0.5);
        label.scale.set(0.25);

        let tween = game.add.tween(label).to({
            alpha: 1
        }, jsonData.path[i].millisecond * (i + 1), Phaser.Easing.Linear.None, true);

        tween.onComplete.add(function () {
            label.destroy();
        });

        unit.oldPoint.push(label)
    }
}