function SelectDetail(boxDetail, unitElement, pic, picDetail) {
    var picUnit = document.getElementById("picUnit");
    var detailUnitBox = document.getElementById(unitElement);
    detailUnitBox.detail = boxDetail.detail;
    detailUnitBox.style.backgroundImage = boxDetail.style.backgroundImage;
    // тут происходит магия ¯\_(ツ)_/¯
    var picDet = document.getElementById(pic);

    if (!picDet){
        picDet = detailUnitBox.cloneNode(false);
        picDet.id = pic;
        picDet.className = picDetail;
        picDet.style.backgroundImage = boxDetail.style.backgroundImage;
    }

    picUnit.appendChild(picDet);

    SendEventAddOrDelDetail()
}