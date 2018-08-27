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