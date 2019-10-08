function FillSquadBlock(squad) {
    if (!squad) return;

    fillSquadUnit("MS", squad.mather_ship);
    fillEquipBlock(squad.mather_ship);

    for (let i in squad.mather_ship.units) {
        if (squad.mather_ship.units.hasOwnProperty(i) && squad.mather_ship.units[i].unit) {
            fillSquadUnit("unitSlot" + i, squad.mather_ship.units[i].unit);
            fillEquipBlock(squad.mather_ship.units[i].unit);
        }
    }

    fillFormation(Data.squad, scaleFormation);
}

let scaleFormation = 2;

function changeScaleFormation(change) {
    if (scaleFormation + change > 0) {
        scaleFormation += change;
        fillFormation(Data.squad, scaleFormation);
    }
}

function fillFormation(squad, scale) {
    // окно строя 150/150px центр 75 75
    // TODO подсвечивать придерживается юнит строя или нет

    let formationUnits = document.getElementById("formationUnits");
    let notFormationUnits = document.getElementById("notFormationUnits");
    formationUnits.innerHTML = ``;
    notFormationUnits.innerHTML = ``;

    let createNotPosUnit = function (unit, content) {
        notFormationUnits.innerHTML += `
            <div onclick="NewFormationPos(${unit.id}, 0, 0)">${content}</div>
        `
    };

    let crateBlock = function (unit, x, y, content, ms) {

        if (!ms && x === 0 && y === 0) {
            createNotPosUnit(unit, content)
        }

        let width = unit.body.height / scale, height = unit.body.width / scale;

        let unitBlock = document.createElement("div");
        unitBlock.className = "formationUnit";
        unitBlock.style.height = (height * 2) + "px";
        unitBlock.style.width = (width * 2) + "px";
        unitBlock.style.top = ((75 + y / scale) - height) + "px";
        unitBlock.style.left = ((75 + x / scale) - width) + "px";
        unitBlock.innerHTML = content;
        formationUnits.appendChild(unitBlock);

        if (!ms) {
            $(unitBlock).draggable({
                containment: "#formationUnits",
                stop: function (event, ui) {
                    let stopPos = $(this).position();

                    stopPos.x = (stopPos.left - 75 + width) * scale;
                    stopPos.y = (stopPos.top - 75 + height) * scale;

                    NewFormationPos(unit.id, stopPos.x, stopPos.y)
                }
            })
        }
    };


    crateBlock(squad.mather_ship, 0, 0, "↑", true);

    for (let i in squad.mather_ship.units) {
        if (squad.mather_ship.units.hasOwnProperty(i) && squad.mather_ship.units[i].unit) {
            let unit = squad.mather_ship.units[i].unit;
            if (unit.on_map) {
                if (unit.formation_pos) {
                    crateBlock(unit, unit.formation_pos.x, unit.formation_pos.y, squad.mather_ship.units[i].number_slot, false);
                } else {
                    createNotPosUnit(unit, squad.mather_ship.units[i].number_slot)
                }
            }
        }
    }
}

function NewFormationPos(unitID, x, y) {
    global.send(JSON.stringify({
        event: "NewFormationPos",
        unit_id: Number(unitID),
        x: Math.round(x),
        y: Math.round(y),
    }))
}

function fillSquadUnit(id, unit) {
    let msBlock = document.getElementById(id);
    let unitEquip = document.getElementById('unitEquip' + unit.id);

    if (unitEquip) {
        // TODO ленивое обновление данных

        if (id !== 'MS') {

        } else {
            $('#reactorStatus').html(`
            <div id="countPower">${(unit.power / 100).toFixed(0)} / ${(unit.max_power / 100).toFixed(0)}</div>
            <div id="recoverPower">+${(unit.recovery_power / 100).toFixed(1)} <span>ед/сек.</span></div>`)
        }
    } else {
        if (id !== 'MS') {

            let backButton = "";
            if (unit.on_map) {
                backButton = "background: url(https://img.icons8.com/cute-clipart/64/000000/login-rounded-up.png) center center / contain no-repeat, rgba(0, 0, 0, 0.6);"
            }

            msBlock.style.background = "none";
            msBlock.innerHTML += `
                <div id="unitEquip${unit.id}">
                    <div class="unitButtonsMask">
                        <div class="selectUnit" title="Выделить и показать" onclick="FocusUnit(${unit.id})"></div>
                        <div class="outUnit" style="${backButton}" title="Вывести/Вернуть" onclick="PlaceUnit(${unit.id})"></div>
                        <div class="changeAmmo"  title="Сменить боеприпасы" onclick="ChangeAmmo(${unit.id})"></div>
                        <div class="openUnitHold" title="Открыть трюм" onclick="OpenInventoryUnit(${unit.id})"></div>
                    </div>
                    
                    <div class="unitEquip">
                        <div id="${unit.id}21"></div>
                        <div id="${unit.id}22"></div>
                        <div id="${unit.id}23"></div>
                        
                        <div id="${unit.id}31"></div>
                        <div id="${unit.id}32"></div>
                        <div id="${unit.id}33"></div>
                    </div>
                    
                    <div class="hpBarWrapper">
                        <div class="hpCount" id="hp${unit.id}"></div>
                    </div>
                    
                    <div class="energyBarWrapper">
                        <div class="energyCount" id="energy${unit.id}"></div>
                    </div>
                </div>
            `;
        } else {

            msBlock.innerHTML += ` 
                <div id="unitEquip${unit.id}">
                    <div class="unitButtonsMask">
                        <div class="openUnitHold" title="Открыть трюм" onclick="InitInventoryMenu(null, 'inventory')"></div>
                        <div class="selectUnit" title="Выделить и показать" onclick="FocusUnit(${unit.id})"></div>
                        <div class="changeAmmo" title="Сменить боеприпасы" onclick="ChangeAmmo(${unit.id})"></div>
                    </div>
                    
                    <div id="MSEquip">
                        <div id="${unit.id}21"></div>
                        <div id="${unit.id}22"></div>
                        <div id="${unit.id}23"></div>
                        <div id="${unit.id}24"></div>
                        <div id="${unit.id}25"></div>
                        
                        <div id="${unit.id}31"></div>
                        <div id="${unit.id}32"></div>
                        <div id="${unit.id}33"></div>
                        <div id="${unit.id}34"></div>
                        <div id="${unit.id}35"></div>
                    </div>
                    
                    <div class="hpBarWrapper">
                        <div class="hpCount" id="hp${unit.id}"></div>
                    </div>
                </div>
        `;
            $('#reactorStatus').html(`
        <div id="countPower">${(unit.power / 100).toFixed(0)} / ${(unit.max_power / 100).toFixed(0)}</div>
        <div id="recoverPower">+${(unit.recovery_power / 100).toFixed(1)} <span>ед/сек.</span></div>`)
        }
    }


    for (let i in msBlock.childNodes) {

        if (msBlock.childNodes[i].nodeName === "SPAN") msBlock.childNodes[i].style.visibility = "hidden";

        // что бы не грузит картинку каждый раз при обновление
        if (msBlock.childNodes[i].className === "body" && msBlock.childNodes[i].style.background.indexOf(unit.body.name) < 0) {
            msBlock.childNodes[i].style.background = "url(/assets/units/body/" + unit.body.name + ".png)" +
                " center center / contain no-repeat, url(/assets/units/body/" + unit.body.name + "_bottom.png) center center / contain no-repeat, rgba(76, 76, 76, 0.66)";
            msBlock.childNodes[i].style.visibility = "visible";
        }

        if (msBlock.childNodes[i].className === "weapon") {
            for (let j in unit.body.weapons) {
                if (unit.body.weapons.hasOwnProperty(j) && unit.body.weapons[j].weapon) {
                    // что бы не грузит картинку каждый раз при обновление
                    if (msBlock.childNodes[i].style.background.indexOf(unit.body.weapons[j].weapon.name) < 0) {
                        msBlock.childNodes[i].style.background = "url(/assets/units/weapon/" + unit.body.weapons[j].weapon.name
                            + ".png) center bottom / contain no-repeat, rgba(76, 76, 76, 0.66)";
                    }
                }
            }
            msBlock.childNodes[i].style.visibility = "visible";
        }

        if (msBlock.childNodes[i].className === "ammo") {
            for (let j in unit.body.weapons) {
                if (unit.body.weapons.hasOwnProperty(j) && unit.body.weapons[j].weapon && unit.body.weapons[j].ammo) {

                    // что бы не грузит картинку каждый раз при обновление
                    if (msBlock.childNodes[i].style.background.indexOf(unit.body.weapons[j].ammo.name) < 0) {
                        msBlock.childNodes[i].style.background = "url(/assets/units/ammo/" + unit.body.weapons[j].ammo.name
                            + ".png) center center / cover no-repeat, rgba(76, 76, 76, 0.66)";
                    }

                    let percentHP = 100 / (unit.body.weapons[j].ammo_capacity / unit.body.weapons[j].ammo_quantity);
                    msBlock.childNodes[i].innerHTML = `
                    <span class="ammoCount" style="color: ${GetColorDamage(percentHP)}">${unit.body.weapons[j].ammo_quantity}</span>`
                }
            }
            msBlock.childNodes[i].style.visibility = "visible";
        }
    }
}

function GetColorDamage(percentHP) {
    if (percentHP === 100) {
        return "#00ff0f"
    } else if (percentHP < 100 && percentHP >= 75) {
        return "#fff326"
    } else if (percentHP < 75 && percentHP >= 50) {
        return "#fac227"
    } else if (percentHP < 50 && percentHP >= 25) {
        return "#fa7b31"
    } else if (percentHP < 25) {
        return "#ff2615"
    }
}

function fillEquipBlock(unit) {
    function checkEquip(equips, type) {

        for (let i in equips) {
            if (equips.hasOwnProperty(i)) {

                let equip = equips[i].equip;

                if (equip) {
                    let equipSlot = document.getElementById(unit.id + "" + type + "" + i);

                    if (equipSlot) {
                        // TODO ленивое обновление данных

                        if (equips[i].hp < equip.max_hp) {
                            let percentHP = 100 / (equip.max_hp / equips[i].hp);
                            equipSlot.style.boxShadow = " inset 0 0 2px 2px " + GetColorDamage(percentHP)
                        }

                        let back = "url(/assets/units/equip/icon/" + equip.name + ".png)" +
                            " center center / contain no-repeat, rgba(76, 76, 76, 0.66)";

                        // что бы не грузит картинку каждый раз при обновление
                        if (equipSlot.style.background.indexOf(equip.name) < 0) {
                            equipSlot.style.background = back;

                            let progressBar = document.createElement("div");
                            progressBar.id = "reloadEquip" + unit.id + type + i;
                            progressBar.className = "reloadEquip";
                            progressBar.style.animation = "reload " + equip.current_reload + "s linear 1";
                            equipSlot.appendChild(progressBar);
                        }

                        equipSlot.onclick = function () {
                            UnselectAll();
                            global.send(JSON.stringify({
                                event: "SelectEquip",
                                unit_id: unit.id,
                                type_slot: Number(type),
                                slot: Number(i),
                            }));
                        };
                    }
                }
            }
        }
    }

    checkEquip(unit.body.equippingI, 1);
    checkEquip(unit.body.equippingII, 2);
    checkEquip(unit.body.equippingIII, 3);
    checkEquip(unit.body.equippingIV, 4);
    checkEquip(unit.body.equippingV, 5);


    let hpBar = document.getElementById('hp' + unit.id);
    let percentHP = 100 / (unit.body.max_hp / unit.hp);
    hpBar.style.background = GetColorDamage(percentHP);
    hpBar.style.width = percentHP + "%";

    let energyBar = document.getElementById('energy' + unit.id);
    if (energyBar) {
        let percent = 100 / (unit.body.max_power / unit.power);
        energyBar.style.width = percent + "%";
    }
}

function PlaceUnit(unitID) {
    global.send(JSON.stringify({
        event: "PlaceUnit",
        unit_id: Number(unitID),
    }))
}