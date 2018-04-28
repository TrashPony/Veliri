function MatherShipsParse(jsonMessage) {
    var matherShips = JSON.parse(jsonMessage).mather_ships;
    if (matherShips.length > 0) {
        var sliderContent = document.getElementById("sliderContent");
        sliderContent.matherShips = matherShips;
    }
}

function ConfigurationMatherShip(matherShip) {

    var type = document.getElementById("MatherShipType");
    var slotSize = document.getElementById("MatherShipSlotSize");
    var unitsTD = document.getElementById("unitsTD");
    var paramsTD = document.getElementById("paramsTD");
    var sliderContent = document.getElementById("sliderContent");

    sliderContent.style.backgroundImage = "url(/lobby/img/" + sliderContent.matherShips[0].type + ".png)";
    type.innerHTML = "<spen class='Value'>" + sliderContent.matherShips[0].type + "</spen>";
    slotSize.innerHTML = "Размер доков:" + "<spen class='Value'>" + sliderContent.matherShips[0].unit_slot_size + "</spen>";

    var hp = document.createElement("span");
    hp.innerHTML = "Hp: " + matherShip.hp + "<br>";
    var armor = document.createElement("span");
    armor.innerHTML = "Armor: " + matherShip.armor + "<br>";
    var rangeView = document.createElement("span");
    rangeView.innerHTML = "RangeView: " + matherShip.range_view + "<br>";

    var equipmentSpan = document.createElement("span");
    equipmentSpan.innerHTML = "Оборудование: ";
    paramsTD.appendChild(equipmentSpan);

    for (var i = 0; i < sliderContent.matherShips[0].unit_slots; i++) {
        var boxUnit = document.createElement("div");
        boxUnit.className = "boxUnit";
        boxUnit.innerHTML = "+";
        boxUnit.onclick = InitCreateUnit; // TODO создать метода добавления и создания юнитов
        unitsTD.appendChild(boxUnit);
        // TODO передовать в параметры выбранный слот в шипе что бы знать куда положить/заменить модуль
    }

    for (var j = 0; j < matherShip.equipment_slots; j++) {
        var boxEqiup = document.createElement("div");
        boxEqiup.style.textAlign = "center";
        boxEqiup.className = "boxEquip";
        boxEqiup.innerHTML = "+";
        boxEqiup.onclick = InitEquippingMenu;
        paramsTD.appendChild(boxEqiup);
        // TODO передовать в параметры выбранный слот в шипе что бы знать куда положить/заменить модуль
    }

    var button = document.createElement("input");
    button.type = "button";
    button.value = "Применить";
    button.className = "lobbyButton";
    button.style.position = "absolute";
    button.style.right = "15px";

    paramsTD.appendChild(document.createElement("br"));
    paramsTD.appendChild(hp);
    paramsTD.appendChild(armor);
    paramsTD.appendChild(rangeView);
    paramsTD.appendChild(button);
}