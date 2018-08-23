function CreateSquadBlock(parentElem) {

    let choiceSquad = document.createElement("div");
    choiceSquad.id = "choiceSquad";
    parentElem.appendChild(choiceSquad);

    let choiceSquadTable = document.createElement('table');
    choiceSquadTable.className = "table";
    choiceSquadTable.id = "choiceSquadTable";
    choiceSquad.appendChild(choiceSquadTable);

    let unitTR = document.createElement('tr');
    choiceSquadTable.appendChild(unitTR);

    let matherShipTD = document.createElement('td');
    matherShipTD.id = "matherShipTD";
    unitTR.appendChild(matherShipTD);

    let unitsTD = document.createElement('td');
    unitsTD.id = "unitsTD";
    unitTR.appendChild(unitsTD);
}

function FillSquadBlock(jsonMessage) {

    while (document.getElementsByClassName("lobbyUnitBox").length > 0) {
        document.getElementsByClassName("lobbyUnitBox")[0].remove();
    }

    let squad = JSON.parse(jsonMessage).squad;
    let matherShipTD = document.getElementById("matherShipTD");

    if (squad.mather_ship != null && squad.mather_ship.body) {
        matherShipTD.style.backgroundImage = "url(/assets/" + squad.mather_ship.body.name + ".png)";

        if (squad.mather_ship.units) {
            let unitsTD = document.getElementById("unitsTD");

            for (let i in squad.mather_ship.units) {
                if (squad.mather_ship.units.hasOwnProperty(i)) {
                    let unit = squad.mather_ship.units[i].unit;
                    if (unit) {
                        let unitDiv = document.createElement("div");
                        unitDiv.className = "lobbyUnitBox";
                        unitDiv.style.backgroundImage = "url(/assets/" + unit.body.name + ".png)";

                        unitsTD.appendChild(unitDiv);
                    }
                }
            }
        }
    }
}