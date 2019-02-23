function LoadQueueUnits() {
    document.getElementById("queueLine").innerHTML = "";

    if (game.Phase === "move") {
        document.getElementById("queue").style.visibility = "visible";
    }

    let moveUnit = document.getElementById("moveUnit");
    moveUnit.innerHTML = "";

    let holdUnits = document.getElementsByClassName("boxUnit");
    for (let i = 0; i < holdUnits.length; i++) {
        holdUnits[i].className = "boxUnit";
    }

    let allActionUnits = [];

    for (let i in game.memoryHostileUnit) {
        if (game.memoryHostileUnit.hasOwnProperty(i) && game.memoryHostileUnit[i].action_points > 0) {
            allActionUnits.push(game.memoryHostileUnit[i]);
        }
    }

    for (let i in game.unitStorage) {
        if (game.unitStorage.hasOwnProperty(i) && game.unitStorage[i].action_points > 0) {
            allActionUnits.push(game.unitStorage[i]);
        }
    }

    for (let q in game.units) {
        if (game.units.hasOwnProperty(q)) {
            for (let r in game.units[q]) {
                if (game.units[q].hasOwnProperty(r) && game.units[q][r].action_points > 0) {
                    if (game.user.name === game.units[q][r].owner) {
                        allActionUnits.push(game.units[q][r]);
                    }
                }
            }
        }
    }

    function initiativeSort(a, b) {
        if (a.initiative < b.initiative) return 1;
        if (a.initiative > b.initiative) return -1;
    }

    allActionUnits.sort(initiativeSort);

    let move = false;
    let randomUnit = [];

    for (let i = 0; i < allActionUnits.length; i++) {

        if (allActionUnits[i].move) {
            createQueueBlock([allActionUnits[i]], allActionUnits[i].move);
            move = true;
            if (allActionUnits.length === i + 1) {
                createQueueBlock(randomUnit);
            }

            continue;
        }

        if ((allActionUnits.length > i + 1 && allActionUnits[i + 1] && allActionUnits[i].initiative !== allActionUnits[i + 1].initiative)
            || allActionUnits.length > i + 2 && allActionUnits[i + 1].move && allActionUnits[i].initiative !== allActionUnits[i + 2].initiative
            || allActionUnits.length === i + 1) {

            randomUnit.push(allActionUnits[i]);
            createQueueBlock(randomUnit);
            randomUnit = [];
        } else if (allActionUnits.length > i + 1 && allActionUnits[i].initiative === allActionUnits[i + 1].initiative) {
            randomUnit.push(allActionUnits[i]);
        }

    }

    if (!move) {
        moveUnit.style.backgroundImage = "url(/assets/unknown.png)";
        moveUnit.style.boxShadow = "inset 0px 0px 35px 1px rgba(255,0,0,0.8)";
        moveUnit.style.animation = "none";
        moveUnit.style.border = "3px solid rgba(108, 108, 108, 0.7)"
    }
}

function createQueueBlock(units) {
    if (units.length === 1) {
        if (units[0].move) {
            let moveUnit = document.getElementById("moveUnit");
            moveUnit.style.backgroundImage = "url(/assets/units/body/" + units[0].body.name + ".png)";

            let weaponMoveUnit = document.getElementsByClassName("weaponMoveUnit");
            while (weaponMoveUnit.childElementCount > 0) {
                weaponMoveUnit.firstChild.remove();
            }

            let weapon = document.createElement("div");
            weapon.className = "weaponMoveUnit";
            for (let i in units[0].body.weapons) {
                if (units[0].body.weapons.hasOwnProperty(i) && units[0].body.weapons[i].weapon) {
                    weapon.style.backgroundImage = "url(/assets/units/weapon/" + units[0].body.weapons[i].weapon.name + ".png)";
                }
            }

            if (units[0].body.mother_ship) {
                weapon.style.backgroundSize = "50%";
                weapon.style.backgroundRepeat = "no-repeat";
                weapon.style.marginTop = "29px";
                weapon.style.marginLeft = "29px";
            }

            moveUnit.appendChild(weapon);

            if (game.user.name === units[0].owner) {
                moveUnit.style.animation = "queuePulse 2s infinite";
                moveUnit.style.border = "3px solid rgb(104, 255, 89)"
            } else {
                moveUnit.style.boxShadow = "none";
                moveUnit.style.border = "3px solid rgba(108, 108, 108, 0.7)"
            }

            if (units[0].on_map) {
                moveUnit.onclick = function () {
                    if (game.user.name === units[0].owner) {
                        SelectUnit.call(units[0]);
                    }
                }
            } else {
                let boxUnit = document.getElementById(units[0].id);
                if (boxUnit) {
                    boxUnit.className = "boxUnit Move";

                    moveUnit.onclick = function () {

                        CreateUnitSubMenu(units[0]);

                        field.send(JSON.stringify({
                            event: "SelectStorageUnit",
                            unit_id: Number(units[0].id)
                        }));
                    }
                }
            }

        } else {
            let queueLine = document.getElementById("queueLine");

            let unitBlock = document.createElement("div");
            unitBlock.id = units[0].id;
            unitBlock.className = "queueLineUnitBlock";
            unitBlock.style.backgroundImage = "url(/assets/units/body/" + units[0].body.name + ".png)";

            let weapon = document.createElement("div");
            weapon.className = "weaponQueueLineUnit";
            for (let i in units[0].body.weapons) {
                if (units[0].body.weapons.hasOwnProperty(i) && units[0].body.weapons[i].weapon) {
                    weapon.style.backgroundImage = "url(/assets/units/weapon/" + units[0].body.weapons[i].weapon.name + ".png)";
                }
            }

            if (units[0].body.mother_ship) {
                weapon.style.backgroundSize = "50%";
                weapon.style.backgroundRepeat = "no-repeat";
                weapon.style.marginTop = "25px";
                weapon.style.marginLeft = "25px";
            }

            unitBlock.appendChild(weapon);

            if (game.user.name === units[0].owner) {
                unitBlock.style.boxShadow = "inset 0px 0px 35px 1px rgba(0, 255, 64, 0.8)";
            } else {
                unitBlock.style.boxShadow = "inset 0px 0px 35px 1px rgba(255,0,0,0.8)";
            }

            queueLine.appendChild(unitBlock);
        }
    } else if (units.length > 1) {
        let queueLine = document.getElementById("queueLine");

        let unitBlock = document.createElement("div");
        unitBlock.className = "queueLineUnitBlock";
        unitBlock.style.backgroundImage = "url(/assets/dice.png)";
        unitBlock.style.boxShadow = "inset 0px 0px 35px 1px rgba(255, 255, 64, 0.8)";
        unitBlock.units = units;
        unitBlock.onmousemove = openDice;
        unitBlock.onmouseout = function () {
            document.getElementById("unitDice").remove();
        };

        queueLine.appendChild(unitBlock);
    }
}

function openDice() {
    let unitDice = document.getElementById("unitDice");
    if (!unitDice) {
        unitDice = document.createElement("div");
        unitDice.id = "unitDice";

        for (let i = 0; i < this.units.length; i++) {
            let unitBlock = document.createElement("div");
            unitBlock.className = "diceUnitBlock";
            unitBlock.style.backgroundImage = "url(/assets/units/body/" + this.units[i].body.name + ".png)";

            let weapon = document.createElement("div");
            weapon.className = "weaponDiceQueueUnit";
            for (let j in this.units[i].body.weapons) {
                if (this.units[i].body.weapons.hasOwnProperty(j) && this.units[i].body.weapons[j].weapon) {
                    weapon.style.backgroundImage = "url(/assets/units/weapon/" + this.units[i].body.weapons[j].weapon.name + ".png)";
                }
            }
            unitBlock.appendChild(weapon);

            if (game.user.name === this.units[i].owner) {
                unitBlock.style.boxShadow = "inset 0px 0px 35px 1px rgba(0, 255, 64, 0.8)";
            } else {
                unitBlock.style.boxShadow = "inset 0px 0px 35px 1px rgba(255,0,0,0.8)";
            }

            unitDice.appendChild(unitBlock);
        }
    }

    unitDice.style.top = stylePositionParams.top + "px";
    unitDice.style.left = stylePositionParams.left + "px";

    document.body.appendChild(unitDice);
}