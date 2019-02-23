function VisibleAnomalies(anomalies) {
    RemoveOldAnomaly();
    // вычислять точки куда присрать новый сигнал и вставлять
    let display = document.getElementById("anomalyDisplay");
    display.offsetHeight;


    for (let i = 0; i < anomalies.length; i++) {

        anomalies[i].rotate -= 180;
        if (anomalies[i].rotate < 0) {
            anomalies[i].rotate += 360
        }

        if (anomalies[i].rotate > 360) {
            anomalies[i].rotate -= 360
        }

        let radRotate = (anomalies[i].rotate) * Math.PI / 180;

        let yAttach = display.offsetHeight / 2 * Math.cos(radRotate);
        let xAttach = display.offsetHeight / 2 * Math.sin(radRotate);

        xAttach += display.offsetHeight / 2 - 8; // 8 это радиус дива пинга
        yAttach += display.offsetHeight / 2 - 8;

        let anomalyPing = document.createElement("div");
        anomalyPing.className = "anomalyPing";
        anomalyPing.style.left = xAttach + "px";
        anomalyPing.style.top = yAttach + "px";

        // мин 0.3 макс 3
        anomalyPing.style.transform = "scale("+ Number(3 - (0.3 * anomalies[i].signal/10)) +")";

        if (anomalies[i].type_anomaly === 999){ // неопознаный тип
            anomalyPing.style.background = "#15fff5";
        }

        if (anomalies[i].type_anomaly === 0 || anomalies[i].type_anomaly === 1){ // ящик
            anomalyPing.style.background = "#fff604";
        }

        if (anomalies[i].type_anomaly === 2){ // руда
            anomalyPing.style.background = "#13ff12";
        }

        if (anomalies[i].type_anomaly === 3){ // текс, квест
            anomalyPing.style.background = "#0b0fff";
        }

        display.appendChild(anomalyPing);
    }
}

function RemoveOldAnomaly() {
    let oldAnomalies = document.getElementsByClassName('anomalyPing');
    while (oldAnomalies.length > 0) {
        oldAnomalies[0].remove();
    }
}