function FillPowerPanel(bodyShip, panel) {
    let powerPanel = document.getElementById(panel);
    let usePower = 0;

    equipSumPower(bodyShip.equippingI); // считаем всю энергию модулей
    equipSumPower(bodyShip.equippingII);
    equipSumPower(bodyShip.equippingIII);
    equipSumPower(bodyShip.equippingIV);
    equipSumPower(bodyShip.equippingV);

    for (let i in bodyShip.weapons) { // и докидываем энергию оружия
        if (bodyShip.weapons.hasOwnProperty(i) && bodyShip.weapons[i].weapon) {
            usePower = usePower + bodyShip.weapons[i].weapon.power;
        }
    }

    function equipSumPower(equip) {
        for (let i in equip) {
            if (equip.hasOwnProperty(i) && equip[i].equip) {
                usePower = usePower + equip[i].equip.power
            }
        }
    }

    if (panel === "powerPanel") {
        powerPanel.innerHTML = "<span class='Value'> Энергия: <br>" + usePower + "/" + bodyShip.max_power + "</span>";
    } else {
        powerPanel.innerHTML = "<span class='Value'>" + usePower + "/" + bodyShip.max_power + "</span>";
    }
}

function FillCubePanel(bodyShip, panel) {
    let cubePanel = document.getElementById(panel);
    let useSize = 0;

    equipSumPower(bodyShip.equippingI);
    equipSumPower(bodyShip.equippingII);
    equipSumPower(bodyShip.equippingIII);
    equipSumPower(bodyShip.equippingIV);
    equipSumPower(bodyShip.equippingV);

    for (let i in bodyShip.weapons) {
        if (bodyShip.weapons.hasOwnProperty(i) && bodyShip.weapons[i].weapon) {
            useSize = useSize + bodyShip.weapons[i].weapon.size;
            if (bodyShip.weapons[i].ammo) {
                useSize = useSize + bodyShip.weapons[i].ammo.size * bodyShip.weapons[i].ammo_quantity
            }
        }
    }

    function equipSumPower(equip) {
        for (let i in equip) {
            if (equip.hasOwnProperty(i) && equip[i].equip) {
                useSize = useSize + equip[i].equip.size
            }
        }
    }

    cubePanel.innerHTML = "<span class='Value'>" + useSize + "/" + bodyShip.capacity_size + "</span>";
}