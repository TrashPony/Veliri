function MarkTarget(target) {

    let q = target.q;
    let r = target.r;

    let mark = game.add.sprite(0, 0, 'MarkTarget'); // создаем метку
    mark.scale.set(.32);
    mark.alpha = 0.8;
    mark.z = 1;

    if (game.map.OneLayerMap[q][r].sprite) {
        game.map.OneLayerMap[q][r].sprite.addChild(mark);
    }
}

function DeleteMarkTarget(target) {
    if (target) {
        let q = target.q;
        let r = target.r;

        for (let i in game.map.OneLayerMap[q][r].sprite.children) {
            if (game.map.OneLayerMap[q][r].sprite.children[i].key === "MarkTarget") {
                game.map.OneLayerMap[q][r].sprite.children[i].destroy()
            }
        }
    }
}