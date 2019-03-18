let visible = true;

function AnomalyCatch(jsonData) {
    document.getElementById("catchAnomaly").style.visibility = 'visible';
    visible = true;
    setTimeout(function () {
        visible = false;
    }, 900);

    setTimeout(function () {
        if (!visible) {
            document.getElementById("catchAnomaly").style.visibility = 'hidden';
        }
    }, 1300);
    console.log(jsonData)
}