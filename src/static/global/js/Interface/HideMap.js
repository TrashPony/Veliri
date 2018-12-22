let hide = false;

function HideMap() {
    let mapBox = document.getElementById("miniMap");
    let canvas = document.getElementById("canvasMap");
    let buttons = document.getElementsByClassName("zoomButton");

    if (!hide) {

        for (let i = 0; i < buttons.length; i++) {
            buttons[i].style.opacity = "0"
        }
        mapBox.style.transition = "1s";
        mapBox.style.height = "20px";
        canvas.style.height = "0";
        canvas.style.opacity = "0";

        setTimeout(function () {
            mapBox.style.transition = "0s";
        }, 1000);
        hide = true;
    } else {

        for (let i = 0; i < buttons.length; i++) {
            buttons[i].style.opacity = "1"
        }

        mapBox.style.transition = "1s";
        mapBox.style.height = "230px";
        canvas.style.height = "200px";
        canvas.style.opacity = "1";

        setTimeout(function () {
            mapBox.style.transition = "0s";
            CreateMiniMap(game.map)
        }, 1000);
        hide = false;
    }
}