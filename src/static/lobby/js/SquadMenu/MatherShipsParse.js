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

    sliderContent.style.backgroundImage = "url(/assets/" + sliderContent.matherShips[0].type + ".png)";
    type.innerHTML = "<spen class='Value'>" + sliderContent.matherShips[0].type + "</spen>";
    slotSize.innerHTML = "Размер доков:" + "<spen class='Value'>" + sliderContent.matherShips[0].unit_slot_size + "</spen>";

    var hp = document.createElement("span");
    hp.innerHTML = "Hp: " + matherShip.hp + "<br>";
    var armor = document.createElement("span");
    armor.innerHTML = "Armor: " + matherShip.armor + "<br>";
    var rangeView = document.createElement("span");
    rangeView.innerHTML = "RangeView: " + matherShip.range_view + "<br>";


    var equipmentSpan = document.createElement("span");
    equipmentSpan.innerHTML = "Модули: ";
    paramsTD.appendChild(equipmentSpan);

    var equippingPanel = CreateEquippingPanel(matherShip);
    paramsTD.appendChild(equippingPanel);

    for (var i = 0; i < sliderContent.matherShips[0].unit_slots; i++) {
        var boxUnit = document.createElement("div");

        boxUnit.className = "boxUnit";
        boxUnit.innerHTML = "+";
        boxUnit.id = i + ":unitSlot";

        boxUnit.onclick = function () {
            InitCreateUnit(this);
        };

        unitsTD.appendChild(boxUnit);
    }

    paramsTD.appendChild(document.createElement("br"));
    paramsTD.appendChild(hp);
    paramsTD.appendChild(armor);
    paramsTD.appendChild(rangeView);

    var button = document.createElement("input");
    button.type = "button";
    button.value = "Удалить";
    button.id = "ButtonDelSquad";
    button.onclick = function () {
        DeleteSquad();
    };

    paramsTD.appendChild(button);
}

function CreateEquippingPanel(matherShip) {
    var equippingPanel = document.createElement("div");
    equippingPanel.id = "equippingPanel";

    for (var j = 0; j < matherShip.equipment_slots; j++) {
        var boxEquip = document.createElement("div");

        boxEquip.className = "boxEquip";
        boxEquip.innerHTML = "+";
        boxEquip.id = j + ":equipSlot";

        boxEquip.onclick = function () {
            InitEquippingMenu(this);
        };

        equippingPanel.appendChild(boxEquip);
    }

    return equippingPanel
}