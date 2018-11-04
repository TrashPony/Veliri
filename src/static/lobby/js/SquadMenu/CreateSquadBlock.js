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
        matherShipTD.style.backgroundImage = "url(/assets/units/body/" + squad.mather_ship.body.name + ".png)";

        if (document.getElementById("lobbyMSWeapon")) {
            document.getElementById("lobbyMSWeapon").remove();
        }

        let weaponIcon = document.createElement("div");
        weaponIcon.id = "lobbyMSWeapon";

        for (let i in  squad.mather_ship.body.weapons) {
            if (squad.mather_ship.body.weapons.hasOwnProperty(i) && squad.mather_ship.body.weapons[i].weapon) {
                weaponIcon.style.backgroundImage = "url(/assets/units/weapon/" + squad.mather_ship.body.weapons[i].weapon.name + ".png)"
            }
        }
        matherShipTD.appendChild(weaponIcon);


        if (squad.mather_ship.units) {
            let unitsTD = document.getElementById("unitsTD");

            for (let i in squad.mather_ship.units) {
                if (squad.mather_ship.units.hasOwnProperty(i)) {
                    let unit = squad.mather_ship.units[i].unit;
                    if (unit) {
                        let unitDiv = document.createElement("div");
                        unitDiv.className = "lobbyUnitBox";
                        unitDiv.style.backgroundImage = "url(/assets/units/body/" + unit.body.name + ".png)";

                        if (document.getElementById(unit.id + "weapon")) {
                            document.getElementById(unit.id + "weapon").remove();
                        }

                        let weaponIcon = document.createElement("div");
                        weaponIcon.className = "lobbyUnitWeapon";
                        weaponIcon.id = unit.id + "weapon";

                        for (let i in  unit.body.weapons) {
                            if (unit.body.weapons.hasOwnProperty(i) && unit.body.weapons[i].weapon) {
                                weaponIcon.style.backgroundImage = "url(/assets/units/weapon/" + unit.body.weapons[i].weapon.name + ".png)"
                            }
                        }
                        unitDiv.appendChild(weaponIcon);

                        unitsTD.appendChild(unitDiv);
                    }
                }
            }
        }
    }
}