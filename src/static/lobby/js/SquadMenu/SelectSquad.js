function SelectSquad(select) {

    DeleteInfoSquad();

    var squad = select.options[select.selectedIndex];
    var idParse = squad.id.split(':'); // "id:squad"

    lobby.send(JSON.stringify({
        event: "SelectSquad",
        squad_id: Number(idParse[0])
    }));

    if (squad.matherShip !== null && squad.matherShip.id !== 0) {
        var sliderContent = document.getElementById("sliderContent");
        if (sliderContent.matherShips) {
            for (var i = 0; i < sliderContent.matherShips.length; i++) { // прокручиваем слайдер до нужного ))
                if (sliderContent.matherShips[i].id === squad.matherShip.id) {

                    var tmpMatherShips = sliderContent.matherShips[0];
                    sliderContent.matherShips[0] = sliderContent.matherShips[i];
                    sliderContent.matherShips[i] = tmpMatherShips;

                    ConfigurationMatherShip(sliderContent.matherShips[0]);
                }
            }
        }
    }

    for (var unitSlot in squad.units) {
        if(squad.units.hasOwnProperty(unitSlot)) {
            var boxUnit= document.getElementById(unitSlot + ":unitSlot");
            if (squad.units[unitSlot] !== null) {
                if (boxUnit) {
                    boxUnit.unit = squad.units[unitSlot];
                    boxUnit.innerHTML = " ";
                    boxUnit.style.backgroundImage = "url(/assets/test1.png)"; // todo как то генерить картинку юнита
                } else {

                    var unitsTD = document.getElementById("unitsTD");

                    var boxErrorUnit= document.createElement("div");

                    boxErrorUnit.equip = squad.equip[unitSlot];
                    boxErrorUnit.className = "boxUnit Error";
                    boxErrorUnit.id = unitSlot + ":unitSlot";
                    boxErrorUnit.innerHTML = " ";
                    boxErrorUnit.style.backgroundImage = "url(/assets/test1.png)";
                    boxErrorUnit.onclick = function () {
                        DeleteUnit(this, unitSlot)
                    };
                    unitsTD.appendChild(boxErrorUnit);
                }
            }
        }
    }

    for (var equipSlot in squad.equip) {
        if(squad.equip.hasOwnProperty(equipSlot)) {
            var boxEquip= document.getElementById(equipSlot + ":equipSlot");
            if (squad.equip[equipSlot] !== null) {
                if (boxEquip) {
                    boxEquip.equip = squad.equip[equipSlot];
                    boxEquip.innerHTML = " ";
                    boxEquip.style.backgroundImage = "url(/assets/" + squad.equip[equipSlot].type + ".png)";
                } else {
                    var equippingPanel = document.getElementById("equippingPanel");

                    var boxErrorEquip = document.createElement("div");

                    boxErrorEquip.equip = squad.equip[equipSlot];
                    boxErrorEquip.className = "boxEquip Error";
                    boxErrorEquip.id = equipSlot + ":equipSlot";
                    boxErrorEquip.innerHTML = " ";
                    boxErrorEquip.style.backgroundImage = "url(/assets/" + squad.equip[equipSlot].type + ".png)";
                    boxErrorEquip.onclick = function () {
                        DeleteEquip(this, equipSlot)
                    };
                    equippingPanel.appendChild(boxErrorEquip)
                }
            }
        }
    }
}