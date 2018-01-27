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

    var headTR = CreateHeadTable();
    choiceSquadTable.appendChild(headTR);

    var unitTR = document.createElement('tr');
    choiceSquadTable.appendChild(unitTR);

    var matherShipTD = CreateMatherShipTD();
    unitTR.appendChild(matherShipTD);

    var unitsTD = document.createElement('td');
    unitsTD.height = "50px";
    unitsTD.id = "unitsTD";
    unitsTD.style.textAlign = "center";
    unitsTD.style.verticalAlign = "middle";
    unitTR.appendChild(unitsTD);

    var paramsTR = document.createElement('tr');
    choiceSquadTable.appendChild(paramsTR);

    var paramsTD = document.createElement('td');
    paramsTD.id = "paramsTD";
    paramsTD.height = "120px";
    paramsTD.style.paddingLeft = "25px";
    paramsTR.appendChild(paramsTD);
}

function CreateHeadTable() {
    var headTR = document.createElement("tr");
    var headTD = document.createElement("td");
    headTD.colSpan = 2;
    headTD.style.textAlign = "center";
    var headSpan = document.createElement("span");
    headSpan.className = "Value";
    headSpan.innerHTML = "Выберите отряд";
    headTD.appendChild(headSpan);
    headTR.appendChild(headTD);

    return headTR;
}

function CreateMatherShipTD() {
    var matherShipTD = document.createElement('td');
    matherShipTD.style.textAlign = "center";
    matherShipTD.height = "155px";
    matherShipTD.width = "130px";
    matherShipTD.rowSpan = 2;
    matherShipTD.style.position = "relative";

    var slider = CreateSliderMatherShip();
    matherShipTD.appendChild(slider);

    var matherShipInfo = CreateMatherShipInfo();
    matherShipTD.appendChild(matherShipInfo);

    return matherShipTD;
}

function CreateMatherShipInfo() {
    var matherShipInfo = document.createElement("div");
    matherShipInfo.id = "infoMatherShip";

    var Type = document.createElement("span");
    Type.id = "MatherShipType";
    matherShipInfo.appendChild(Type);

    matherShipInfo.appendChild(document.createElement("br"));

    var SlotSize = document.createElement("span");
    SlotSize.id = "MatherShipSlotSize";
    matherShipInfo.appendChild(SlotSize);

    matherShipInfo.appendChild(document.createElement("p"));

    return matherShipInfo;
}