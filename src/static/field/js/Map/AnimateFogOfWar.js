function AnimateFog() {
    for (var fog in game.fogOfWar.children) {
        if (game.fogOfWar.children[fog].hide) {
            if (game.fogOfWar.children[fog].alpha <= 0.01) {
                game.fogOfWar.children[fog].visible = false;
                continue
            }
            game.fogOfWar.children[fog].alpha = game.fogOfWar.children[fog].alpha - 0.01;

        } else {
            game.fogOfWar.children[fog].visible = true;
            if (game.fogOfWar.children[fog].alpha < 1) {
                game.fogOfWar.children[fog].alpha = game.fogOfWar.children[fog].alpha + 0.01;
            }
        }
    }
}