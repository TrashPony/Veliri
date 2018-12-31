function UpdateReservoir(jsonData) {
    for (let q in game.map.reservoir) {
        for (let r in game.map.reservoir[q]) {
            if (game.map.reservoir[q][r] && jsonData.q === Number(q) && jsonData.r === Number(r)) {
                game.map.reservoir[q][r].count = jsonData.count;

                let tip = document.getElementById("reservoirTip" + q + "" + r);
                if (tip) {
                    document.getElementById("countOre" + q + "" + r).innerHTML = jsonData.count;
                }
            }
        }
    }
}

function DestroyReservoir(jsonData) {
    for (let q in game.map.reservoir) {
        for (let r in game.map.reservoir[q]) {
            if (game.map.reservoir[q][r] && jsonData.q === Number(q) && jsonData.r === Number(r)) {

                let tween = game.add.tween(game.map.reservoir[q][r].sprite).to({alpha: 0}, 1000, Phaser.Easing.Linear.None, true, 0);
                game.add.tween(game.map.reservoir[q][r].sprite.shadow).to({alpha: 0}, 1000, Phaser.Easing.Linear.None, true, 0);
                tween.onComplete.add(function () {
                    game.map.reservoir[q][r].sprite.destroy();
                    game.map.reservoir[q][r].sprite.shadow.destroy();

                    game.map.OneLayerMap[q][r].move = true;
                    CreateMiniMap()
                })
            }
        }
    }
}