function GameInfo() {
    let step = document.getElementById('step');
    step.innerHTML = game.Step;

    let phaseGame = document.getElementById('phase');
    phaseGame.innerHTML = game.Phase;

    if (game.Phase !== "move") {
        document.getElementById("moveUnit").style.visibility = "hidden";
        document.getElementById("queueMove").style.visibility = "hidden";
    }
}