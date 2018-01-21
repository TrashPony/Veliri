function MatherShipsParse(jsonMessage) {
    var matherShips = JSON.parse(jsonMessage).mather_ships;
    if (matherShips.length > 0) {
        var sliderContent = document.getElementById("sliderContent");
        sliderContent.matherShips = matherShips;

        NextSlider(sliderContent);
    }
}

function NextSlider(sliderContent) {
    var type = document.getElementById("MatherShipType");
    var slotSize = document.getElementById("MatherShipSlotSize");
    var unitsTD = document.getElementById("unitsTD");

    sliderContent.style.backgroundImage = "url(/lobby/img/" + sliderContent.matherShips[0].type + ".png)";

    type.innerHTML = "<spen class='Value'>" + sliderContent.matherShips[0].type + "</spen>";
    slotSize.innerHTML = "Размер доков:" + "<spen class='Value'>" + sliderContent.matherShips[0].unit_slot_size + "</spen>";

    for (var i = 0; i < sliderContent.matherShips[0].unit_slots; i++) {
        var boxUnit = document.createElement("div");
        boxUnit.className = "boxUnit";
        boxUnit.innerHTML = "+";

        unitsTD.appendChild(boxUnit);
    }
}