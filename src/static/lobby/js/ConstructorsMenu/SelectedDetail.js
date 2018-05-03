function SelectDetail(detail, unitElement, onMouse) {

    var detailUnitBox = document.getElementById(unitElement);
    detailUnitBox.detail = detail;
    detailUnitBox.style.backgroundImage = "url(/assets/" + detail.name + ".png)";

    detailUnitBox.onmouseover = function () {
        onMouse(this.detail);
    };

    detailUnitBox.onmouseout = function () {
        TipOff();
    };
}