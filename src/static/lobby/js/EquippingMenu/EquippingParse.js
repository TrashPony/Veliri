function EquippingParse(jsonMessage) {

    var equipping = JSON.parse(jsonMessage).equipping;
    var tableEquip = document.getElementById("tableEquip");

    console.log(equipping);

    for (var i = 0; i < equipping.length; i++) {
        var tr = CreateEquipRow(equipping[i].type, equipping[i].id, equipping[i].specification);
        tableEquip.appendChild(tr);
    }
}
// TODO передовать в параметры выбранный слот в шипе что бы знать куда положить/заменить модуль
function CreateEquipRow(type, id, value) {
    var tr = document.createElement("tr");
    tr.className = "EquipRow";
    var tdIcon = document.createElement("td");
    tdIcon.className = "tdIconEquip";
    tdIcon.style.backgroundImage = "url(/lobby/img/" + type + ".png)";

    var tdSpecification = document.createElement("td");
    tdSpecification.innerHTML = "<spen class='Value'>" + type + "</spen><br>" + value + "<br>";

    var acceptButton = document.createElement("input");
    acceptButton.type = "button";
    acceptButton.value = "Выбрать";
    acceptButton.className = "EquipButton";
    acceptButton.id = id+":equip";
    acceptButton.onclick = function () {
        SelectEquip(this.id);
    };

    tdSpecification.appendChild(acceptButton);
    tr.appendChild(tdIcon);
    tr.appendChild(tdSpecification);

    return tr;
}

function SelectEquip(id) {
    var equippingMenu = document.getElementById("equippingMenu");
    var equipIdParse = id.split(':'); // "id:equip"
    var equipSlot = equippingMenu.equipSlot.split(':'); //"0:equipSlot"

    if (equippingMenu.equip === undefined || equippingMenu.equip === null) {
        lobby.send(JSON.stringify({
            event: "AddEquipment",
            equip_id: Number(equipIdParse[0]),
            equip_slot: Number(equipSlot[0])
        }));
    } else {
        lobby.send(JSON.stringify({
            event: "ReplaceEquipment",
            equip_id: Number(equipIdParse[0]),
            equip_slot: Number(equipSlot[0])
        }));
    }

    EquipBackToLobby();
}