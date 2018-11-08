function FillingEquipPanel(equipPanel, unit) {

    let countSlots;

    if (unit.body.mother_ship) { // по диз доку и мс максимум 5 слотов у юнита 3
        countSlots = 6;
    } else {
        countSlots = 4;
    }

    let weapon = document.createElement("div");
    weapon.id = "weaponSlotSubMenu";

    for (let weaponSlot in unit.body.weapons) { // оружие может быть только 1 под диз доке, масив это обман
        if (unit.body.weapons.hasOwnProperty(weaponSlot) && unit.body.weapons[weaponSlot].weapon) {
            weapon.className = "weaponSlotSubMenu Active";
            weapon.style.backgroundImage = "url(/assets/units/weapon/" + unit.body.weapons[weaponSlot].weapon.name + ".png)";

            weapon.onclick = function () {
                RemoveSelect();
                field.send(JSON.stringify({
                    event: "SelectWeapon",
                    q: Number(unit.q),
                    r: Number(unit.r)
                }));
            };

            let ammoBox = document.createElement("div");
            ammoBox.id = "ammoBox";
            if (unit.body.weapons[weaponSlot].ammo) {
                ammoBox.style.backgroundImage = "url(/assets/units/ammo/" + unit.body.weapons[weaponSlot].ammo.name + ".png)";
            } else {
                ammoBox.className = "blink"
            }

            equipPanel.appendChild(ammoBox);
        } else {
            weapon.className = "weaponSlotSubMenu noActive";
        }
    }

    equipPanel.appendChild(weapon);

    for (let i = 1; i < countSlots; i++) {
        let equipping = document.createElement("div");
        equipping.id = "equipSlotIII" + i;

        if (unit.body.equippingIII.hasOwnProperty(i) && unit.body.equippingIII[i].equip) {

            equipping.style.backgroundImage = "url(/assets/units/equip/" + unit.body.equippingIII[i].equip.name + ".png)";

            if (!unit.body.equippingIII[i].used && !(unit.power < unit.body.equippingIII[i].equip.use_power) && unit.on_map) {
                equipping.className = "equipSlotIII Active";
                equipping.onclick = function () {
                    RemoveSelect();
                    field.send(JSON.stringify({
                        event: "SelectEquip",
                        q: Number(unit.q),
                        r: Number(unit.r),
                        equip_type: 3,
                        number_slot: unit.body.equippingIII[i].number_slot
                    }));
                };

                equipping.onmouseover = function () {
                    TipEquipOn(unit.body.equippingIII[i].equip);
                };
            } else {
                if (unit.body.equippingIII[i].used) {
                    equipping.onmouseover = function () { TipNotAllowEquip("Снаряжение уже использвано"); };
                } else if (unit.power < unit.body.equippingIII[i].equip.use_power) {
                    equipping.onmouseover = function () { TipNotAllowEquip("Не достаточно энергии"); };
                }
                equipping.className = "equipSlotIII notAllow";
            }

            equipping.onmouseout = function () {
                TipEquipOff();
            };
        } else {
            equipping.className = "equipSlotIII noActive";
        }

        equipPanel.appendChild(equipping);
    }

    for (let i = 1; i < countSlots; i++) {
        let equipping = document.createElement("div");
        equipping.id = "equipSlotII" + i;

        if (unit.body.equippingII.hasOwnProperty(i) && unit.body.equippingII[i].equip) {

            equipping.style.backgroundImage = "url(/assets/units/equip/" + unit.body.equippingII[i].equip.name + ".png)";

            if (!unit.body.equippingII[i].used && !(unit.power < unit.body.equippingII[i].equip.use_power) && unit.on_map) {
                equipping.className = "equipSlotII Active";
                equipping.onclick = function () {
                    RemoveSelect();
                    field.send(JSON.stringify({
                        event: "SelectEquip",
                        q: Number(unit.q),
                        r: Number(unit.r),
                        equip_type: 2,
                        number_slot: unit.body.equippingII[i].number_slot
                    }));
                };
                equipping.onmouseover = function () {
                    TipEquipOn(unit.body.equippingII[i].equip);
                };
            } else {
                if (unit.body.equippingII[i].used) {
                    equipping.onmouseover = function () { TipNotAllowEquip("Снаряжение уже использвано"); };
                } else if (unit.power < unit.body.equippingII[i].equip.use_power) {
                    equipping.onmouseover = function () { TipNotAllowEquip("Не достаточно энергии"); };
                }
                equipping.className = "equipSlotII notAllow";
            }

            equipping.onmouseout = function () {
                TipEquipOff();
            };
        } else {
            equipping.className = "equipSlotII noActive";
        }
        equipPanel.appendChild(equipping);
    }

    let power = document.createElement("div");
    power.id = "powerPanel";
    power.innerHTML = "<span>POWER " + unit.power + "/" + unit.max_power + "</span>";
    equipPanel.appendChild(power);

    let actionPoints = document.createElement("div");
    actionPoints.id = "actionPoints";
    actionPoints.innerHTML = "<span class='Value'>AP " + unit.action_points + "/" + +unit.speed + "</span>";
    equipPanel.appendChild(actionPoints);
}