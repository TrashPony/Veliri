function ChangePhase(jsonMessage) {

    ChangePhaseNotification(jsonMessage);

    game.Step = JSON.parse(jsonMessage).game_step;
    game.Phase = JSON.parse(jsonMessage).game_phase;
    GameInfo();

    if (game.Phase === "move") {
        LoadQueueUnits();
    } else {
        document.getElementById("moveUnit").style.visibility = "hidden";
        document.getElementById("queueMove").style.visibility = "hidden";
    }

    if (game.Phase === "targeting" && JSON.parse(jsonMessage).flee_battle) {
        // TODO кнопка ливнуть из боя
    }

    game.user.ready = JSON.parse(jsonMessage).ready;
    InitPlayer();

    let units = JSON.parse(jsonMessage).units;
    for (let x in units) {
        if (units.hasOwnProperty(x)) {
            for (let y in units[x]) {
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