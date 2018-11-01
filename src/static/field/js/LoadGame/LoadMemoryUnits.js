function LoadQueueUnits() {

    // TODO console.log(game.unitStorage);

    let allActionUnits = [];

    for (let i in game.memoryHostileUnit) {
        if (game.memoryHostileUnit.hasOwnProperty(i)) { // && game.memoryHostileUnit[i].action_points > 0
            allActionUnits.push(game.memoryHostileUnit[i]);
        }
    }

    for (let q in game.units) {
        if (game.units.hasOwnProperty(q)) {
            for (let r in game.units[q]) {
                if (game.units[q].hasOwnProperty(r)) { //  && game.units[q][r].action_points > 0
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

    for (let i = 0; i < allActionUnits.length; i++){
        createQueueBlock(allActionUnits[i]);
    }
}

function createQueueBlock(unit) {
    let queueLine = document.getElementById("queueLine");

    let unitBlock = document.createElement("div");
    unitBlock.id = unit.id;
    unitBlock.className = "queueLineUnitBlock";
    unitBlock.style.backgroundImage = "url(/assets/" + unit.body.name + ".png)";

    if (game.user.name === unit.owner) {
        unitBlock.style.boxShadow = "inset 0px 0px 35px 1px rgba(0, 255, 64, 0.8)";
    } else {
        unitBlock.style.boxShadow = "inset 0px 0px 35px 1px rgba(255,0,0,0.8)";
    }

    queueLine.appendChild(unitBlock);
}