function CreateUnitSubMenu(unit) {
    if (game.Phase === "move" || game.Phase === "targeting") {

        let unitSubMenu = document.getElementById("UnitSubMenu");

        if (unitSubMenu) {
            unitSubMenu.remove();
        }

        unitSubMenu = document.createElement("div");
        unitSubMenu.id = "UnitSubMenu";

        unitSubMenu.style.left = stylePositionParams.left - 95 + 'px';
        unitSubMenu.style.top = stylePositionParams.top - 85 + 'px';
        unitSubMenu.style.display = "block";

        let equipPanel = document.createElement("div");
        equipPanel.id = "EquipPanel";
        unitSubMenu.appendChild(equipPanel);

        if (!unit.action && game.user.name === unit.owner) {

            fillingEquipPanel(equipPanel, unit);

            if (game.Phase === "move") {
                MoveSubMenu(equipPanel, unit);
            }

            if (game.Phase === "targeting") {
                TargetingSubMenu(equipPanel, unit);
            }
        } else {
            unitSubMenu.style.animation = "none";
            unitSubMenu.style.border = "0px";
        }

        if (unit.effects !== null && unit.effects.length > 0) {
            EffectsPanel(unitSubMenu, unit);
        }

        document.body.appendChild(unitSubMenu);
    }
}

function fillingEquipPanel(equipPanel, unit) {

    let weapon = document.createElement("div");
    weapon.id = "weaponSlotSubMenu";

    for (let weaponSlot in unit.body.weapons) { // оружие может быть только 1 под диз доке, масив это обман
        if (unit.body.weapons.hasOwnProperty(weaponSlot) && unit.body.weapons[weaponSlot].weapon) {
            weapon.className = "weaponSlotSubMenu Active";
            weapon.style.backgroundImage = "url(/assets/" + unit.body.weapons[weaponSlot].weapon.name + ".png)";

            let ammoBox = document.createElement("div");
            ammoBox.id = "ammoBox";
            if (unit.body.weapons[weaponSlot].ammo) {
                ammoBox.style.backgroundImage = "url(/assets/" + unit.body.weapons[weaponSlot].ammo.name + ".png)";
            } else {
                ammoBox.className = "blink"
            }
            // todo ammo box
            // todo onclick
            // todo mouse over/out

            equipPanel.appendChild(ammoBox);
        } else {
            weapon.className = "weaponSlotSubMenu noActive";
        }
    }

    equipPanel.appendChild(weapon);

    for (let i = 1; i < 3; i++) {
        let equipping = document.createElement("div");
        equipping.id = "equipSlotIII" + i;

        if (unit.body.equippingIII.hasOwnProperty(i) && unit.body.equippingIII[i].equip) {
            equipping.className = "equipSlotIII Active";
            equipping.style.backgroundImage = "url(/assets/" + unit.body.equippingIII[i].equip.name + ".png)";
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

function MoveSubMenu(equipPanel, unit) {
    let table = document.createElement("table");
    table.style.width = "95px";

    let tr = document.createElement("tr");
    let th = document.createElement("th");
    th.style.alignContent = "center";
    th.innerHTML = "Действия:";
    th.className = "h";

    tr.appendChild(th);

    let trSkip = document.createElement("tr");
    let tdSkip = document.createElement("td");
    tdSkip.style.alignContent = "center";
    let skipButton = document.createElement("input");
    skipButton.type = "button";
    skipButton.value = "Пропустить ход";
    skipButton.className = "button subMenu";

    skipButton.onclick = function () {
        field.send(JSON.stringify({
            event: "SkipMoveUnit",
            x: Number(unit.x),
            y: Number(unit.y)
        }));
    };

    tdSkip.appendChild(skipButton);
    trSkip.appendChild(tdSkip);

    table.appendChild(tr);
    table.appendChild(trSkip);

    equipPanel.appendChild(table);
}

function TargetingSubMenu(equipPanel, unit) {
    let table = document.createElement("table");
    table.id = "targetButtonsTable";
    table.style.width = "150px";

    let trDefend = document.createElement("tr");
    let tdDefend = document.createElement("td");
    tdDefend.style.alignContent = "center";
    let defendButton = document.createElement("input");
    defendButton.type = "button";
    defendButton.value = "Защита";
    defendButton.className = "button unitSubMenu";

    defendButton.onclick = function () {
        field.send(JSON.stringify({
            event: "Defend",
            x: Number(unit.x),
            y: Number(unit.y)
        }));
    };

    tdDefend.appendChild(defendButton);
    trDefend.appendChild(tdDefend);

    table.appendChild(trDefend);

    equipPanel.appendChild(table);
}

function EffectsPanel(unitSubMenu, unit) {
    let effectPanel = document.createElement("div");
    effectPanel.id = "effectPanel";

    let head = document.createElement("span");
    head.className = "Value";
    head.innerHTML = "Эффекты:";
    effectPanel.appendChild(head);

    let panel = document.createElement("table");
    panel.className = "panel Effect";

    let rowEffect;
    let count = 0;

    for (let j = 0; j < unit.effects.length; j++) {
        if (unit.effects[j].type !== "unit_always_animate") {
            if (count % 4 === 0) {
                rowEffect = document.createElement("tr");
                rowEffect.className = "row Effect";
            }

            let cellEffect = document.createElement("td");
            cellEffect.className = "cell Effect";
            cellEffect.style.backgroundImage = "url(/assets/effects/" + unit.effects[j].name + "_" + unit.effects[j].level + ".png)";
            cellEffect.effects = unit.effects[j];

            cellEffect.onmouseover = function () {
                TipEffectOn(this.effects);
            };

            cellEffect.onmouseout = function () {
                TipEffectOff();
            };

            rowEffect.appendChild(cellEffect);

            if (count % 4 === 0) {
                panel.appendChild(rowEffect);
            }

            count++;
        }
    }

    effectPanel.appendChild(panel);
    unitSubMenu.appendChild(effectPanel);
}