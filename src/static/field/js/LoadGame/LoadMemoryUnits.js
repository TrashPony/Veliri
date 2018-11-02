function LoadQueueUnits() {

    // TODO game.unitStorage не учитывается
    document.getElementById("queueLine").innerHTML = "";
    document.getElementById("queue").style.visibility = "visible";

    let allActionUnits = [];

    for (let i in game.memoryHostileUnit) {
        if (game.memoryHostileUnit.hasOwnProperty(i) && game.memoryHostileUnit[i].action_points > 0) {
            allActionUnits.push(game.memoryHostileUnit[i]);
        }
    }

    for (let q in game.units) {
        if (game.units.hasOwnProperty(q)) {
            for (let r in game.units[q]) {
                if (game.units[q].hasOwnProperty(r) && game.units[q][r].action_points > 0) {
                    allActionUnits.push(game.units[q][r]);
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

            if (allActionUnits.length === i+1) {
                createQueueBlock(randomUnit);
            }

            continue;
        }
        if ((allActionUnits.length > i + 1 && allActionUnits[i].initiative !== allActionUnits[i + 1].initiative)
            || allActionUnits.length === i + 1) {
            randomUnit.push(allActionUnits[i]);
            createQueueBlock(randomUnit);
            randomUnit = [];
        } else if (allActionUnits.length > i + 1 && allActionUnits[i].initiative === allActionUnits[i + 1].initiative) {
            randomUnit.push(allActionUnits[i]);
        }

    }

    if (!move) {
        let moveUnit = document.getElementById("moveUnit");
        moveUnit.style.backgroundImage = "url(/assets/unknown.png)";
        moveUnit.style.boxShadow = "inset 0px 0px 35px 1px rgba(255,0,0,0.8)";
    }
}

function createQueueBlock(units) {
    if (units.length === 1) {
        if (units[0].move) {
            let moveUnit = document.getElementById("moveUnit");
            moveUnit.style.backgroundImage = "url(/assets/" + units[0].body.name + ".png)";
            if (game.user.name === units[0].owner) {
                moveUnit.style.boxShadow = "inset 0px 0px 35px 1px rgba(0, 255, 64, 0.8)";
            } else {
                moveUnit.style.boxShadow = "inset 0px 0px 35px 1px rgba(255,0,0,0.8)";
            }
        } else {
            let queueLine = document.getElementById("queueLine");

            let unitBlock = document.createElement("div");
            unitBlock.id = units[0].id;
            unitBlock.className = "queueLineUnitBlock";
            unitBlock.style.backgroundImage = "url(/assets/" + units[0].body.name + ".png)";

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
            unitBlock.style.backgroundImage = "url(/assets/" + this.units[i].body.name + ".png)";
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