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
    slideLeft.innerHTML = "&#8592";
    slideLeft.addEventListener('click', SliderMoveLeft);
    sliderNav.appendChild(slideLeft);

    var slideRight = document.createElement("div");
    slideRight.id = "slideRight";
    slideRight.innerHTML = "&#8594";
    slideRight.addEventListener('click', SliderMoveRight);
    sliderNav.appendChild(slideRight);

    return sliderNav
}

function SliderMoveLeft() {
    var sliderContent = document.getElementById("sliderContent");
    DeleteInfoSquad();

    var last = sliderContent.matherShips.pop();    // беру последний обьект
    sliderContent.matherShips.unshift(last);       // кладу его первым

    lobby.send(JSON.stringify({
        event: "SelectMatherShip",
        mather_ship_id: Number(sliderContent.matherShips[0].id)
    }));

    if (sliderContent.matherShips.length > 0) {
        ConfigurationMatherShip(sliderContent.matherShips[0]);
    }
}

function SliderMoveRight() {
    var sliderContent = document.getElementById("sliderContent");
    DeleteInfoSquad();

    var first = sliderContent.matherShips.shift(); // беру перый обьект
    sliderContent.matherShips.push(first);         // кладу его последним

    lobby.send(JSON.stringify({
        event: "SelectMatherShip",
        mather_ship_id: Number(sliderContent.matherShips[0].id)
    }));

    if (sliderContent.matherShips.length > 0) {
        ConfigurationMatherShip(sliderContent.matherShips[0]);
    }
}