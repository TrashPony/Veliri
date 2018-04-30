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