var alpha;

function AlphaSelect() {
    if (alpha === undefined) {
        alpha = false;
    }

    if (game.SelectLineLayer.alpha >= 0.3 && !alpha) {
        game.SelectLineLayer.alpha = game.SelectLineLayer.alpha - 0.01;
        if (game.SelectLineLayer.alpha < 0.3) {
            alpha = true;
        }
    }

    if (alpha) {
        game.SelectLineLayer.alpha = game.SelectLineLayer.alpha + 0.01;
        if (game.SelectLineLayer.alpha >= 0.8) {
            alpha = false
        }
    }
}