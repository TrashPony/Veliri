function GameInfo() {
    let step = document.getElementById('step');
    step.innerHTML = game.Step;

    let phaseGame = document.getElementById('phase');
    phaseGame.innerHTML = game.Phase;

    if (game.Phase !== "move") {
        document.getElementById("queue").style.visibility = "hidden";
    }
}