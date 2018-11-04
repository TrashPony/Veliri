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

function FillMSWeaponTypePanel(bodyShip, panel) {

    let weaponPanel = document.getElementById(panel);
    weaponPanel.innerHTML = "";

    let weaponTypeIcon  = document.createElement("div");
    weaponTypeIcon.id = "weaponTypeIcon";
    weaponPanel.appendChild(weaponTypeIcon);

    weaponTypePanel(bodyShip, weaponPanel, "weaponTypeAllow", "weaponTypeNotAllow")
}


function FillUnitWeaponTypePanel(bodyShip, panel) {
    let weaponPanel = document.getElementById(panel);
    weaponPanel.innerHTML = "";

    weaponTypePanel(bodyShip, weaponPanel, "weaponUnitTypeAllow", "weaponUnitTypeNotAllow")
}

function weaponTypePanel(bodyShip, parent, classAllow, classNotAllow) {

    let smallWeapon = document.createElement("div");
    smallWeapon.className = classNotAllow;
    parent.appendChild(smallWeapon);
    let mediumWeapon = document.createElement("div");
    mediumWeapon.className = classNotAllow;
    parent.appendChild(mediumWeapon);
    let heavyWeapon = document.createElement("div");
    heavyWeapon.className = classNotAllow;
    parent.appendChild(heavyWeapon);

    if (bodyShip.standard_size_small) {
        smallWeapon.className = classAllow;
    }

    if (bodyShip.standard_size_medium) {
        mediumWeapon.className = classAllow;
    }

    if (bodyShip.standard_size_big) {
        heavyWeapon.className = classAllow;
    }
}