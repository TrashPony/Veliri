var alpha;
var alphaTarget;

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

    if (alphaTarget === undefined) {
        alpha = false;
    }

    if (game.SelectTargetLineLayer.alpha >= 0.3 && !alpha) {
        game.SelectTargetLineLayer.alpha = game.SelectTargetLineLayer.alpha - 0.005;
        if (game.SelectTargetLineLayer.alpha < 0.3) {
            alphaTarget = true;
        }
    }

    if (alpha) {
        game.SelectTargetLineLayer.alpha = game.SelectTargetLineLayer.alpha + 0.005;
        if (game.SelectTargetLineLayer.alpha >= 0.8) {
            alphaTarget = false
        }
    }
}