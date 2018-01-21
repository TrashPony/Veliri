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
        NextSlide(sliderContent);
    }
}

function SliderMoveRight() {
    var sliderContent = document.getElementById("sliderContent");
    RemoveUnitBox();

    var first = sliderContent.matherShips.shift(); // беру перый обьект
    sliderContent.matherShips.push(first);         // кладу его последним

    if (sliderContent.matherShips.length > 0) {
        NextSlide(sliderContent);
    }
}

function RemoveUnitBox() {
    var paramsTD = document.getElementById("paramsTD");

    while (paramsTD.childNodes.length > 0) {
        paramsTD.removeChild(paramsTD.childNodes[0]);
    }

    var unitBoxs = document.getElementsByClassName("boxUnit");

    while (unitBoxs.length > 0) {
        unitBoxs[0].remove();
    }
}

function NextSlide(sliderContent) {
    var type = document.getElementById("MatherShipType");
    var slotSize = document.getElementById("MatherShipSlotSize");
    var unitsTD = document.getElementById("unitsTD");

    sliderContent.style.backgroundImage = "url(/lobby/img/" + sliderContent.matherShips[0].type + ".png)";

    type.innerHTML = "<spen class='Value'>" + sliderContent.matherShips[0].type + "</spen>";
    slotSize.innerHTML = "Размер доков:" + "<spen class='Value'>" + sliderContent.matherShips[0].unit_slot_size + "</spen>";

    for (var i = 0; i < sliderContent.matherShips[0].unit_slots; i++) {
        var boxUnit = document.createElement("div");
        boxUnit.className = "boxUnit";
        boxUnit.innerHTML = "+";
        boxUnit.onclick = InitCreateUnit; // TODO создать метода добавления и создания юнитов
        unitsTD.appendChild(boxUnit);
    }
}