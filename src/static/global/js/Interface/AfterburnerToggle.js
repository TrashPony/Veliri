let aburner = false;

function AfterburnerToggle() {
    global.send(JSON.stringify({
        event: "AfterburnerToggle"
    }))
}

function Afterburner(afterburner) {
    let burner = document.getElementById("Afterburner");
    let statusBar = document.getElementById("statusBar");

    if (afterburner) {
        aburner = true;
        statusBar.style.animation = "Afterburner 1000ms infinite";
        statusBar.style.backgroundImage = "linear-gradient(1deg, rgba(255, 33, 33, 0.6), rgba(225, 37, 37, 0.6) 6px)";
        statusBar.style.border = "1px solid #e12525";
        burner.style.filter = "drop-shadow(0px 0px 4px #dd0019)"
    } else {
        aburner = false;
        statusBar.style.animation = "none";
        statusBar.style.backgroundImage = "linear-gradient(1deg, rgba(33, 176, 255, 0.6), rgba(37, 160, 225, 0.6) 6px)";
        statusBar.style.border = "1px solid #25a0e1";
        burner.style.filter = "none"
    }
}