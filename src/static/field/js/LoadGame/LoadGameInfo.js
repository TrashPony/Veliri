function GameInfo(jsonMessage) {

    var step = document.getElementById('step');
    step.innerHTML = JSON.parse(jsonMessage).game_step;

    var phaseGame = document.getElementById('phase');
    phaseGame.innerHTML = JSON.parse(jsonMessage).game_phase;

    phase = JSON.parse(jsonMessage).game_phase;

}