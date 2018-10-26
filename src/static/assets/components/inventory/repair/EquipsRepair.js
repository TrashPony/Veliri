function EquipsRepair() {
    inventorySocket.send(JSON.stringify({
        event: "EquipsRepair"
    }));
}

function overEquipButton() {

    if (document.getElementById("ConstructorUnit")) {
        let equippingCells = document.getElementsByClassName("UnitEquip");

        for (let i = 0; i < equippingCells.length; i++) {
            if (equippingCells[i].className === "UnitEquip active weapon") {
                if (JSON.parse(equippingCells[i].slotData).weapon) {
                    let percentHP = CreateHealBar(equippingCells[i], "weapon", false);
                    if (percentHP !== 100) {
                        equippingCells[i].style.boxShadow = "0 0 5px 3px rgb(255, 243, 38)";
                    }
                }
            } else if (equippingCells[i].className === "UnitEquip active") {
                if (JSON.parse(equippingCells[i].slotData).equip) {
                    let percentHP = CreateHealBar(equippingCells[i], "equip", false);
                    if (percentHP !== 100) {
                        equippingCells[i].style.boxShadow = "0 0 5px 3px rgb(255, 243, 38)";
                    }
                }
            }
        }
    }

    for (let slot = 1; slot <= 6; slot++) {
        let cell = document.getElementById("squad " + slot + 4);

        function colored(cell) {
            cell.className = "inventoryUnit select repair";
        }

        if (cell.slotData && JSON.parse(cell.slotData).unit) {

            let body = JSON.parse(cell.slotData).unit.body;

            if (body.hp < body.max_hp) {
                colored(cell)
            }

            if (checkBroken(body.equippingI)) {
                colored(cell)
            }

            if (checkBroken(body.equippingII)) {
                colored(cell)
            }

            if (checkBroken(body.equippingIII)) {
                colored(cell)
            }

            if (checkBroken(body.equippingIV)) {
                colored(cell)
            }

            if (checkBroken(body.equippingV)) {
                colored(cell)
            }

            function checkBroken(equip) {
                for (let i in equip) {
                    if (equip.hasOwnProperty(i) && equip[i].equip) {
                        if (equip[i].equip.max_hp > equip[i].hp) {
                            return true
                        }
                    }
                }

                return false
            }
        }
    }

    let equippingCells = document.getElementsByClassName("inventoryEquipping");

    for (let i = 0; i < equippingCells.length; i++) {
        if (equippingCells[i].className === "inventoryEquipping active weapon") {
            if (JSON.parse(equippingCells[i].slotData).weapon) {
                let percentHP = CreateHealBar(equippingCells[i], "weapon", false);
                if (percentHP !== 100) {
                    equippingCells[i].style.boxShadow = "0 0 5px 3px rgb(255, 243, 38)";
                }
            }
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

    let constructorUnit = document.getElementById("ConstructorUnit");
    if (constructorUnit) {
        let equippingCells = document.getElementsByClassName("UnitEquip");

        for (let i = 0; i < equippingCells.length; i++) {
            if (equippingCells[i].className === "UnitEquip active weapon") {
                equippingCells[i].style.boxShadow = "0 0 5px 3px rgb(255, 0, 0)";
            }
            if (equippingCells[i].className === "UnitEquip active") {
                equippingCells[i].style.boxShadow = "0 0 10px rgba(0, 0, 0, 1)";
            }
        }
    }

    for (let slot = 1; slot <= 6; slot++) {
        let cell = document.getElementById("squad " + slot + 4);
        if (cell.slotData && JSON.parse(cell.slotData).unit) {
            cell.className = "inventoryUnit active";
        }
    }

    let equippingCells = document.getElementsByClassName("inventoryEquipping");

    for (let i = 0; i < equippingCells.length; i++) {
        if (equippingCells[i].className === "inventoryEquipping active weapon") {
            equippingCells[i].style.boxShadow = "0 0 5px 3px rgb(255, 0, 0)";
        }
        if (equippingCells[i].className === "inventoryEquipping active") {
            equippingCells[i].style.boxShadow = "0 0 10px rgba(0, 0, 0, 1)";
        }
    }
}