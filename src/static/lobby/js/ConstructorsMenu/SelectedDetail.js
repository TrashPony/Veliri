function SelectDetail(detail, unitElement, pic, picDetail) {
    var picUnit = document.getElementById("picUnit");
    var detailUnitBox = document.getElementById(unitElement);
    detailUnitBox.detail = detail;
    detailUnitBox.style.backgroundImage = detail.style.backgroundImage;
    // тут происходит магия ¯\_(ツ)_/¯
    var picDet = document.getElementById(pic);

    if (!picDet){
        picDet = detailUnitBox.cloneNode(false);
        picDet.id = pic;
        picDet.className = picDetail;
        picDet.style.backgroundImage = detail.style.backgroundImage;
    }

    picUnit.appendChild(picDet);
}