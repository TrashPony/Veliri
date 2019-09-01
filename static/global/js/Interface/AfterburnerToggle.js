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
        burner.style.filter = "drop-shadow(0px 0px 4px #dd0019)"
    } else {
        aburner = false;
        statusBar.style.animation = "none";
        burner.style.filter = "none"
    }
}