function EquipsRepair() {
    inventorySocket.send(JSON.stringify({
        event: "EquipsRepair"
    }));
}

function overEquipButton() {
    let equippingCells = document.getElementsByClassName("inventoryEquipping");
    // todo юниты которы сломаны или в них что то сломано надо тоже светить
    for (let i = 0; i < equippingCells.length; i++) {
        if (equippingCells[i].className === "inventoryEquipping active weapon") {
            //todo оружие
        } else if (equippingCells[i].className === "inventoryEquipping active") {
            if (JSON.parse(equippingCells[i].slotData).equip) {
                let percentHP = CreateHealBar(equippingCells[i], "equip", false);
                if (percentHP !== 100) {
                    equippingCells[i].style.boxShadow = "0 0 5px 3px rgb(255, 243, 38)";
                }
            }
        }
    }
}

function outEquipButton() {
    let equippingCells = document.getElementsByClassName("inventoryEquipping");

    for (let i = 0; i < equippingCells.length; i++) {
        if (equippingCells[i].className === "inventoryEquipping active") {
            equippingCells[i].style.boxShadow = "0 0 10px rgba(0, 0, 0, 1)";
        }
    }
}