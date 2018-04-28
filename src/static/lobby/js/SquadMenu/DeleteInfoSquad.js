function DeleteInfoSquad() {

    var paramsTD = document.getElementById("paramsTD");

    while (paramsTD.childNodes.length > 0) {
        paramsTD.removeChild(paramsTD.childNodes[0]);
    }

    var unitBox = document.getElementsByClassName("boxUnit");

    while (unitBox.length > 0) {
        unitBox[0].remove();
    }

    var type = document.getElementById("MatherShipType");
    var slotSize = document.getElementById("MatherShipSlotSize");
    var sliderContent = document.getElementById("sliderContent");

    sliderContent.style.backgroundImage = "";
    type.innerHTML = "";
    slotSize.innerHTML = "";
}