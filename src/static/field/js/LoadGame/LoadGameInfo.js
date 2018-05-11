function GameInfo(gameInfo) {

    var step = document.getElementById('step');
    step.innerHTML = gameInfo.Step;

    var phaseGame = document.getElementById('phase');
    phaseGame.innerHTML = gameInfo.Phase;

    phase = gameInfo.Phase;
}