function MarkTarget(target) {

    var x = target.x;
    var y = target.y;

    var mark = game.add.sprite(0, 0, 'MarkTarget'); // создаем метку
    mark.scale.set(.32);
    mark.alpha = 0.8;
    mark.z = 1;

    if (game.map.OneLayerMap[x][y].sprite) {
        game.map.OneLayerMap[x][y].sprite.addChild(mark);
    }
}

function DeleteMarkTarget(target) {
    if (target) {
        var x = target.x;
        var y = target.y;

        if (game.map.OneLayerMap[x][y].sprite.children.length > 0) {
            var mark = game.map.OneLayerMap[x][y].sprite.getChildAt(0);
            mark.destroy();
        }
    }
}