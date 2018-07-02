function CreateChoiceSquadBlock(parentElem) {

    var choiceSquad = document.createElement("div");
    choiceSquad.id = "choiceSquad";
    parentElem.appendChild(choiceSquad);

    var choiceSquadTable = document.createElement('table');
    choiceSquadTable.className = "table";
    choiceSquadTable.id = "choiceSquadTable";
    choiceSquad.appendChild(choiceSquadTable);

    var unitTR = document.createElement('tr');
    choiceSquadTable.appendChild(unitTR);

    var matherShipTD = document.createElement('td');
    matherShipTD.id = "matherShipTD";
    unitTR.appendChild(matherShipTD);

    var unitsTD = document.createElement('td');
    unitsTD.id = "unitsTD";
    unitTR.appendChild(unitsTD);
}