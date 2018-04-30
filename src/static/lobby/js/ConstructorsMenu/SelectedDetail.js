function SelectDetail(detail, unitElement, pic, picDetail, onMouse) {

    var oldPic = document.getElementById(pic);
    if (oldPic){
        oldPic.remove();
    }

    var picUnit = document.getElementById("picUnit");
    var detailUnitBox = document.getElementById(unitElement);
    detailUnitBox.detail = detail;
    detailUnitBox.style.backgroundImage = "url(/lobby/img/" + detail.name + ".png)";

    detailUnitBox.onmouseover = function () {
        onMouse(this.detail);
    };

    detailUnitBox.onmouseout = function () {
        TipOff();
    };

    // тут происходит магия ¯\_(ツ)_/¯
    var picDet = document.getElementById(pic);

    if (!picDet){
        picDet = detailUnitBox.cloneNode(false);
        picDet.id = pic;
        picDet.className = picDetail;
        picDet.style.backgroundImage = "url(/lobby/img/" + detail.name + ".png)";
    }

    picUnit.appendChild(picDet);

    SendEventAddOrDelDetail()
}