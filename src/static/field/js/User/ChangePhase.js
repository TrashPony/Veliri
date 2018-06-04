function ChangePhase(jsonMessage) {
    console.log(jsonMessage);

    game.Step = JSON.parse(jsonMessage).game_step;
    game.Phase = JSON.parse(jsonMessage).game_phase;
    GameInfo();

    game.user.ready = JSON.parse(jsonMessage).ready;
    InitPlayer();

    var units = JSON.parse(jsonMessage).units;
    for (var x in units) {
        if (units.hasOwnProperty(x)) {
            for (var y in units[x]) {
                if (units[x].hasOwnProperty(y)) {
                    if (game.units.hasOwnProperty(x)) {
                        if (game.units[x].hasOwnProperty(y)) {
                            game.units[x][y].action = units[x][y].action;
                            game.units[x][y].target = units[x][y].target;

                            ActivationUnit(game.units[x][y])
                        }
                    }
                }
            }
        }
    }
}