function GameInfo() {
    var step = document.getElementById('step');
    step.innerHTML = game.Step;

    var phaseGame = document.getElementById('phase');
    phaseGame.innerHTML = game.Phase;
}