function HideMeta() {

    if (game) {
        game.redactorMetaText.forEach(function (e) {
            if (e.alpha === 0) {
                e.alpha = 1
            } else {
                e.alpha = 0
            }
        })
    }
}