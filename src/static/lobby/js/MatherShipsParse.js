function MatherShipsParse(jsonMessage) {
    var matherShips = JSON.parse(jsonMessage).mather_ships;
    if (matherShips.length > 0) {
        var sliderContent = document.getElementById("sliderContent");
        sliderContent.matherShips = matherShips;

        NextSlide(sliderContent);
    }
}

function ConfigurationMatherShip() {
    var paramsTD = document.getElementById("paramsTD");
    var matherShip = document.getElementById("sliderContent").matherShips[0];

    while (paramsTD.childNodes.length > 0) {
        paramsTD.removeChild(paramsTD.childNodes[0]);
    }

    var hp = document.createElement("span");
    hp.innerHTML = "Hp: " + matherShip.hp + "<br>";
    var armor = document.createElement("span");
    armor.innerHTML = "Armor: " + matherShip.armor + "<br>";
    var rangeView = document.createElement("span");
    rangeView.innerHTML = "RangeView: " + matherShip.range_view + "<br>";


    var equipmentSpan = document.createElement("span");
    equipmentSpan.innerHTML = "Оборудование: ";
    paramsTD.appendChild(equipmentSpan);

    for (var i = 0; i < matherShip.equipment_slots; i++) {
        var boxUnit = document.createElement("div");
        boxUnit.style.textAlign = "center";
        boxUnit.style.verticalAlign = "middle";
        boxUnit.className = "boxEquip";
        boxUnit.innerHTML = "+";
        boxUnit.onclick = AddEquipment; // TODO создать метода добавления экипировки
        paramsTD.appendChild(boxUnit);
    }

    paramsTD.appendChild(document.createElement("br"));
    paramsTD.appendChild(hp);
    paramsTD.appendChild(armor);
    paramsTD.appendChild(rangeView);

    var button = document.createElement("input");
    button.type = "button";
    button.value = "Применить";
    button.className = "lobbyButton";
    button.style.position = "absolute";
    button.style.right = "15px";
    //button.onclick = ; TODO отправить данные на сервер
    paramsTD.appendChild(button);

}