
function CreateChoiceSquadBlock(parentElem) {

    var choiceSquad = document.createElement("div");
    choiceSquad.style.position = "absolute";
    choiceSquad.style.bottom = "50px";
    choiceSquad.id = "choiceSquad";
    parentElem.appendChild(choiceSquad);

    var choiceSquadTable = document.createElement('table');
    choiceSquadTable.width = "400px";
    choiceSquadTable.className = "table";
    choiceSquad.appendChild(choiceSquadTable);

    var headTR = document.createElement("tr");
    choiceSquadTable.appendChild(headTR);
    var headTD = document.createElement("td");
    headTD.colSpan = 2;
    headTD.style.textAlign = "center";
    var headSpan = document.createElement("span");
    headSpan.className = "Value";
    headSpan.innerHTML = "Выберите отряд";
    headTD.appendChild(headSpan);
    headTR.appendChild(headTD);

    var unitTR = document.createElement('tr');
    choiceSquadTable.appendChild(unitTR);

    var matherShipTD = document.createElement('td');
    matherShipTD.height = "155px";
    matherShipTD.width = "100px";
    matherShipTD.rowSpan = 2;
    unitTR.appendChild(matherShipTD);

    var unitTD = document.createElement('td');
    unitTD.height = "35px";
    unitTR.appendChild(unitTD);

    var paramsTR = document.createElement('tr');
    choiceSquadTable.appendChild(paramsTR);

    var paramsTD = document.createElement('td');
    paramsTD.height = "120px";
    paramsTR.appendChild(paramsTD);
}