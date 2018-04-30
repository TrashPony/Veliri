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

    if (squad.units !== null && squad.units.length > 0) {
        //todo
    }

    for (var slot in squad.equip) {
        if(squad.equip.hasOwnProperty(slot)) {
            var boxEquip= document.getElementById(slot + ":equipSlot");
            if (squad.equip[slot] !== null) {
                if (boxEquip) {
                    boxEquip.equip = squad.equip[slot];
                    boxEquip.innerHTML = " ";
                    boxEquip.style.backgroundImage = "url(/lobby/img/" + squad.equip[slot].type + ".png)";
                } else {
                    var equippingPanel = document.getElementById("equippingPanel");

                    var boxErrorEquip = document.createElement("div");

                    boxErrorEquip.equip = squad.equip[slot];
                    boxErrorEquip.className = "boxEquip Error";
                    boxErrorEquip.innerHTML = "+";
                    boxErrorEquip.id = slot + ":equipSlot";
                    boxErrorEquip.innerHTML = " ";
                    boxErrorEquip.style.backgroundImage = "url(/lobby/img/" + squad.equip[slot].type + ".png)";
                    boxErrorEquip.onclick = function () {
                        DeleteEquip(this, slot)
                    };
                    equippingPanel.appendChild(boxErrorEquip)
                }
            }
        }
    }
}