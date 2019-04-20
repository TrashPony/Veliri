function FillSquadBlock(squad) {
    fillSquadUnit("MS", squad.mather_ship);
    fillDamageEquip("damageMS", squad.mather_ship);
    fillMiningBlock(squad.mather_ship);

    for (let i in squad.mather_ship.units) {
        if (squad.mather_ship.units.hasOwnProperty(i) && squad.mather_ship.units[i].unit) {
            fillSquadUnit("unitSlot" + i, squad.mather_ship.units[i].unit);
            fillDamageEquip("damageUnit", squad.mather_ship.units[i].unit);
        }
    }
}

function fillSquadUnit(id, unit) {
    let msBlock = document.getElementById(id);
    msBlock.style.background = "none";

    for (let i in msBlock.childNodes) {
        if (msBlock.childNodes[i].nodeName === "SPAN") msBlock.childNodes[i].style.visibility = "hidden";

        if (msBlock.childNodes[i].className === "body") {
            msBlock.childNodes[i].style.background = "url(/assets/units/body/" + unit.body.name + ".png)" +
                " center center / contain no-repeat, rgba(76, 76, 76, 0.66)";
            msBlock.childNodes[i].style.visibility = "visible";
        }
        if (msBlock.childNodes[i].className === "weapon") {
            for (let j in unit.body.weapons) {
                if (unit.body.weapons.hasOwnProperty(j) && unit.body.weapons[j].weapon) {
                    msBlock.childNodes[i].style.background = "url(/assets/units/weapon/" + unit.body.weapons[j].weapon.name
                        + ".png) center center / contain no-repeat, rgba(76, 76, 76, 0.66)";
                }
            }
            msBlock.childNodes[i].style.visibility = "visible";
        }
        if (msBlock.childNodes[i].className === "ammo") {
            for (let j in unit.body.weapons) {
                if (unit.body.weapons.hasOwnProperty(j) && unit.body.weapons[j].weapon && unit.body.weapons[j].ammo) {
                    msBlock.childNodes[i].style.background = "url(/assets/units/ammo/" + unit.body.weapons[j].ammo.name
                        + ".png) center center / contain no-repeat, rgba(76, 76, 76, 0.66)";
                }
            }
            msBlock.childNodes[i].style.visibility = "visible";
        }
    }
}

function fillDamageEquip(id, unit) {
    let damageBlock = document.getElementById(id);
    damageBlock.innerHTML = "";

    if (unit.hp < unit.body.max_hp) {
        let wrapper = document.createElement("div");
        let damage = document.createElement("div");
        damage.className = "damageItem";
        damage.style.background = "url(/assets/units/body/" + unit.body.name + ".png)" +
            " center center / contain no-repeat";

        let healBar = createHealBat(unit.hp, unit.body.max_hp);

        if (unit.body.mother_ship && 100 / (unit.body.max_hp / unit.hp) < 25) {
            document.getElementById("criticalDamage").style.visibility = 'visible';
        }

        wrapper.appendChild(healBar);
        wrapper.appendChild(damage);

        damageBlock.appendChild(wrapper)
    }

    for (let j in unit.body.weapons) {
        if (unit.body.weapons.hasOwnProperty(j) && unit.body.weapons[j].weapon) {
            if (unit.body.weapons[j].weapon && unit.body.weapons[j].hp < unit.body.weapons[j].weapon.max_hp) {
                let wrapper = document.createElement("div");
                let damage = document.createElement("div");
                damage.className = "damageItem";
                damage.style.background = "url(/assets/units/weapon/" + unit.body.weapons[j].weapon.name + ".png)" +
                    " center center / contain no-repeat";

                let healBar = createHealBat(unit.body.weapons[j].hp, unit.body.weapons[j].weapon.max_hp);

                wrapper.appendChild(healBar);
                wrapper.appendChild(damage);

                damageBlock.appendChild(wrapper)
            }
        }
    }

    function checkEquip(equips) {
        for (let i in equips) {
            if (equips.hasOwnProperty(i)) {
                if (equips[i].equip && equips[i].hp < equips[i].equip.max_hp) {
                    let wrapper = document.createElement("div");
                    let damage = document.createElement("div");
                    damage.className = "damageItem";
                    damage.style.background = "url(/assets/units/equip/" + equips[i].equip.name + ".png)" +
                        " center center / contain no-repeat";

                    let healBar = createHealBat(equips[i].hp, equips[i].equip.max_hp);

                    wrapper.appendChild(healBar);
                    wrapper.appendChild(damage);

                    damageBlock.appendChild(wrapper)
                }
            }
        }
    }

    checkEquip(unit.body.equippingI);
    checkEquip(unit.body.equippingII);
    checkEquip(unit.body.equippingIII);
    checkEquip(unit.body.equippingIV);
    checkEquip(unit.body.equippingV);
}

function createHealBat(hp, maxHP) {
    let backHealBar = document.createElement("div");
    backHealBar.className = "damageBackHealBar";

    let percentHP = 100 / (maxHP / hp);

    let healBar = document.createElement("div");
    healBar.className = "damageHealBar";
    healBar.style.width = percentHP + "%";

    if (percentHP === 100) {
        backHealBar.style.opacity = "0"
    } else if (percentHP < 100 && percentHP >= 75) {
        healBar.style.backgroundColor = "#fff326"
    } else if (percentHP < 75 && percentHP >= 50) {
        healBar.style.backgroundColor = "#fac227"
    } else if (percentHP < 50 && percentHP >= 25) {
        healBar.style.backgroundColor = "#fa7b31"
    } else if (percentHP < 25 && hp > 1) {
        healBar.style.backgroundColor = "#ff2615"
    } else if (hp === 0) {
        healBar.style.opacity = "0";
    }

    backHealBar.appendChild(healBar);
    return backHealBar
}

function fillMiningBlock(unit) {
    let mining = document.getElementById("MiningPanel");
    mining.innerHTML = "";

    function checkEquip(equips, type) {
        for (let i in equips) {
            if (equips.hasOwnProperty(i)) {
                if (equips[i].equip && equips[i].hp > 0 && (equips[i].equip.applicable === "ore" ||
                    equips[i].equip.applicable === "geo_scan" || equips[i].equip.applicable === "digger")) {
                    mining.style.visibility = "visible";

                    let equipBlock = document.createElement("div");
                    equipBlock.style.background = "url(/assets/units/equip/icon/" + equips[i].equip.name + ".png)" +
                        " center center / contain no-repeat, rgba(76, 76, 76, 0.66)";

                    if (equips[i].equip.applicable === "ore") {
                        equipBlock.onclick = function () {
                            InitMiningOre(equips[i].equip, i, type);
                        };
                    }

                    if (equips[i].equip.applicable === "digger") {
                        equipBlock.onclick = function () {
                            InitDigger(equips[i].equip, i, type);
                        };
                    }

                    let progressBar = document.createElement("div");
                    progressBar.id = "miningEquip" + equips[i].type_slot + i;
                    equipBlock.appendChild(progressBar);

                    mining.appendChild(equipBlock);
                }
            }
        }
    }

    checkEquip(unit.body.equippingI, 1);
    checkEquip(unit.body.equippingII, 2);
    checkEquip(unit.body.equippingIII, 3);
    checkEquip(unit.body.equippingIV, 4);
    checkEquip(unit.body.equippingV, 5);
}