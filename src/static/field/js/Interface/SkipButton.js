function ActiveSkipButton(unitX, unitY) {
    var skipButton = document.getElementById("SkipButton");
    skipButton.className = "button";

    skipButton.onclick = function () {
        field.send(JSON.stringify({
            event: "SkipMoveUnit",
            x: Number(unitX),
            y: Number(unitY)
        }));
    }
}

function DeactiveSkipButton() {
    var skipButton = document.getElementById("SkipButton");
    skipButton.className = "button noActive";
    skipButton.onclick = null;
}
