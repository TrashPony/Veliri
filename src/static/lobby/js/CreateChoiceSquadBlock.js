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
    paramsTD.height = "120px";
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

    var button = document.createElement("input");
    button.type = "button";
    button.value = "Настроить";
    button.className = "button";
    //button.onclick = ;
    matherShipInfo.appendChild(button);

    return matherShipInfo;
}

function CreateSliderMatherShip() {

    var slider = document.createElement("div");
    var sliderContent = document.createElement("div");
    sliderContent.id = "sliderContent";
    slider.appendChild(sliderContent);

    lobby.send(JSON.stringify({
        event: "GetMatherShips"
    }));

    var sliderNav = CreateNavigationSlider();
    slider.appendChild(sliderNav);

    return slider;
}

function CreateNavigationSlider() {
    var sliderNav = document.createElement("div");

    var slideLeft = document.createElement("div");
    slideLeft.id = "slideLeft";
    slideLeft.addEventListener('click', SliderMoveLeft);
    slideLeft.innerHTML = "&#8592";
    sliderNav.appendChild(slideLeft);

    var slideRight = document.createElement("div");
    slideRight.addEventListener('click', SliderMoveRight);
    slideRight.id = "slideRight";
    slideRight.innerHTML = "&#8594";
    sliderNav.appendChild(slideRight);

    return sliderNav
}

function SliderMoveLeft() {
    var sliderContent = document.getElementById("sliderContent");
    RemoveUniBox();

    var last = sliderContent.matherShips.pop();    // беру последний обьект
    sliderContent.matherShips.unshift(last);       // кладу его первым

    if (sliderContent.matherShips.length > 0) {
        NextSlider(sliderContent);
    }
}

function SliderMoveRight() {
    var sliderContent = document.getElementById("sliderContent");
    RemoveUniBox();

    var first = sliderContent.matherShips.shift(); // беру перый обьект
    sliderContent.matherShips.push(first);         // кладу его последним

    if (sliderContent.matherShips.length > 0) {
        NextSlider(sliderContent);
    }
}

function RemoveUniBox() {
    var unitBoxs = document.getElementsByClassName("boxUnit");
    while (unitBoxs.length > 0) {
        unitBoxs[0].remove();
    }
}