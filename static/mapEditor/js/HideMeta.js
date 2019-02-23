let metaAlpha = 0;

function HideMeta() {
    if (game) {

        if (metaAlpha === 0) {
            metaAlpha = 1;
        } else {
            metaAlpha = 0;
        }

        game.redactorMetaText.forEach(function (e) {
            e.alpha = metaAlpha;
        })
    }
}