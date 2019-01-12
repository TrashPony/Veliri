function alertError(jsonData) {
    // TODO не рабочий код
    let event = JSON.parse(jsonData).event;
    let error = JSON.parse(jsonData).error;

    console.log(error);

    if (event === "ms error") {

        let powerPanel = document.getElementById("powerPanel");

        let start = Date.now();

        let timer = setInterval(function () {
            let timePassed = Date.now() - start;
            if (timePassed >= 600) {
                clearInterval(timer);
                powerPanel.style.border = "1px solid #25a0e1";
                powerPanel.style.boxShadow = "none";
                return;
            }
            powerPanel.style.boxShadow = "inset 1px 1px 25px 1px rgba(255,0,0,1)";
            powerPanel.style.border = "1px solid #e10006";
        }, 20);

    } else if (event === "unit error") {

        let panel;

        if (JSON.parse(jsonData).error === "lacking size") {
            panel = document.getElementById("unitCubePanel");
        } else if (JSON.parse(jsonData).error === "lacking power") {
            panel = document.getElementById("unitPowerPanel");
        } else if (JSON.parse(jsonData).error === "wrong standard size") {
            panel = document.getElementById("weaponTypePanel");
        }

        if (panel) {
            let start = Date.now();
            let timer = setInterval(function () {
                let timePassed = Date.now() - start;
                if (timePassed >= 600) {
                    clearInterval(timer);
                    panel.style.border = "1px solid #25a0e1";
                    panel.style.boxShadow = "none";
                    return;
                }
                panel.style.boxShadow = "inset 1px 1px 25px 1px rgba(255,0,0,1)";
                panel.style.border = "1px solid #e10006";
            }, 20);
        }
    }
}