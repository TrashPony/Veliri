function FillingEquipPanel(equipPanel, unit) {

    let weapon = document.createElement("div");
    weapon.id = "weaponSlotSubMenu";

    for (let weaponSlot in unit.body.weapons) { // оружие может быть только 1 под диз доке, масив это обман
        if (unit.body.weapons.hasOwnProperty(weaponSlot) && unit.body.weapons[weaponSlot].weapon) {
            weapon.className = "weaponSlotSubMenu Active";
            weapon.style.backgroundImage = "url(/assets/" + unit.body.weapons[weaponSlot].weapon.name + ".png)";

            weapon.onclick = function () {
                RemoveSelect();
                field.send(JSON.stringify({
                    event: "SelectWeapon",
                    x: Number(unit.x),
                    y: Number(unit.y)
                }));
            };


            let ammoBox = document.createElement("div");
            ammoBox.id = "ammoBox";
            if (unit.body.weapons[weaponSlot].ammo) {
                ammoBox.style.backgroundImage = "url(/assets/" + unit.body.weapons[weaponSlot].ammo.name + ".png)";
            } else {
                ammoBox.className = "blink"
            }

            equipPanel.appendChild(ammoBox);
        } else {
            weapon.className = "weaponSlotSubMenu noActive";
        }
    }

    equipPanel.appendChild(weapon);

    for (let i = 1; i < 4; i++) {
        let equipping = document.createElement("div");
        equipping.id = "equipSlotIII" + i;

        if (unit.body.equippingIII.hasOwnProperty(i) && unit.body.equippingIII[i].equip) {
            equipping.className = "equipSlotIII Active";
            equipping.style.backgroundImage = "url(/assets/" + unit.body.equippingIII[i].equip.name + ".png)";

            equipping.onclick = function () {
                RemoveSelect();
                field.send(JSON.stringify({
                    event: "SelectEquip",
                    x: Number(unit.x),
                    y: Number(unit.y),
                    equip_type: 3,
                    number_slot: unit.body.equippingIII[i].number_slot
                }));
            };
        } else {
            equipping.className = "equipSlotIII noActive";
        }

        equipPanel.appendChild(equipping);
    }

    for (let i = 1; i < 4; i++) {
        let equipping = document.createElement("div");
        equipping.id = "equipSlotII" + i;

        if (unit.body.equippingII.hasOwnProperty(i) && unit.body.equippingII[i].equip) {
            equipping.className = "equipSlotII Active";
            equipping.style.backgroundImage = "url(/assets/" + unit.body.equippingII[i].equip.name + ".png)";

            equipping.onclick = function () {
                RemoveSelect();
                field.send(JSON.stringify({
                    event: "SelectEquip",
                    x: Number(unit.x),
                    y: Number(unit.y),
                    equip_type: 2,
                    number_slot: unit.body.equippingII[i].number_slot
                }));
            };
        } else {
            equipping.className = "equipSlotIII noActive";
        }

        equipPanel.appendChild(equipping);
    }

    let power = document.createElement("div");
    power.id = "powerPanel";
    power.innerHTML = "<span>POWER " + unit.power + "/" + unit.body.max_power + "</span>";
    equipPanel.appendChild(power);

    console.log(unit);
}