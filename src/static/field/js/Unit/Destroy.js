function UnitDestroy() {
    for (var x in game.units) {
        if (game.units.hasOwnProperty(x)) {
            for (var y in game.units[x]) {
                if (game.units[x].hasOwnProperty(y)) {
                    var unit = game.units[x][y];
                    if (unit.destroy && unit.sprite.alpha < 0.1) {
                        console.log("Destroy unit");
                        unit.sprite.destroy();
                        unit.shadow.destroy();
                        delete game.units[x][y];
                    } else {
                        if (unit.destroy) {
                            game.add.tween(unit.sprite).to({alpha: 0}, 100, Phaser.Easing.Linear.None, true);
                            game.add.tween(unit.shadow).to({alpha: 0}, 100, Phaser.Easing.Linear.None, true);
                        }
                    }
                }
            }
        }
    }
}