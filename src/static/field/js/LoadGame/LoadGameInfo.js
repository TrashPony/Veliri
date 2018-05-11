function GameInfo() {
    var step = document.getElementById('step');
    step.innerHTML = game.gameInfo.Step;

    var phaseGame = document.getElementById('phase');
    phaseGame.innerHTML = game.gameInfo.Phase;
}