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

    for (var key in squad.equip) {
        if(squad.equip.hasOwnProperty(key)) {
            var boxEquip= document.getElementById(key + ":equipSlot");
            if (boxEquip) {
                boxEquip.equip = squad.equip[key];
                boxEquip.innerHTML = " ";
                boxEquip.style.backgroundImage = "url(/lobby/img/" + squad.equip[key].type + ".png)";
            } else {
                var equippingPanel = document.getElementById("equippingPanel");

                var boxErrorEquip = document.createElement("div");

                boxErrorEquip.equip = squad.equip[key];
                boxErrorEquip.className = "boxEquip Error";
                boxErrorEquip.innerHTML = "+";
                boxErrorEquip.id = key+":equipSlot";
                boxErrorEquip.innerHTML = " ";
                boxErrorEquip.style.backgroundImage = "url(/lobby/img/" + squad.equip[key].type + ".png)";

                equippingPanel.appendChild(boxErrorEquip)
            }
        }
    }
}