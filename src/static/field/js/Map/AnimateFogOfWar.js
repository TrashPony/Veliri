function AnimateFog() {
    for (var fog in game.fogOfWar.children) {
        if (game.fogOfWar.children[fog].hide) {
            game.fogOfWar.children[fog].alpha = game.fogOfWar.children[fog].alpha - 0.01;
            if (game.fogOfWar.children[fog].alpha <= 0.01) {
                game.fogOfWar.children[fog].visible = false;
            }
        }
    }
}