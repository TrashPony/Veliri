function ViewPatternCoordinate(typeCoordinates) {

    let coordinateBlock = document.getElementById("coordinateBlock");

    for (let i = 0; i < typeCoordinates.length; i++) {
        let typeBlock = document.createElement("div");
        typeBlock.className = "coordinateType";

        if (typeCoordinates[i].texture_object === "") {
            typeBlock.style.background = "url(/assets/map/terrain/" + typeCoordinates[i].texture_flore + ".png)  center center / 115% no-repeat";
        } else {
            typeBlock.style.background = "url(/assets/map/objects/" + typeCoordinates[i].texture_object + ".png)  center center / 90% no-repeat," +
                " url(/assets/map/terrain/" + typeCoordinates[i].texture_flore + ".png)  center center / 115% no-repeat";
        }

        typeBlock.coordinateType = typeCoordinates[i];

        typeBlock.onclick = function (){
            PlaceCoordinate("placeCoordinate");
        };
        typeBlock.onmousemove = tipTypeCoordinate;
        typeBlock.onmouseout = function () {
            if (document.getElementById("typeTip")) {
                document.getElementById("typeTip").remove()
            }
        };

        coordinateBlock.appendChild(typeBlock);
    }
}

function ViewTerrainCoordinate(typeCoordinates) {
    let coordinateBlock = document.getElementById("coordinateBlock");

    for (let i = 0; i < typeCoordinates.length; i++) {
        if (!document.getElementById(typeCoordinates[i].texture_flore)) {
            let typeBlock = document.createElement("div");
            typeBlock.className = "coordinateType";
            typeBlock.style.background = "url(/assets/map/terrain/" + typeCoordinates[i].texture_flore + ".png)  center center / 115% no-repeat";

            typeBlock.id = typeCoordinates[i].texture_flore;

            typeBlock.onclick = function (){
                PlaceCoordinate("placeTerrain");
            };
            coordinateBlock.appendChild(typeBlock);
        }
    }
}

function ViewObjectsCoordinate(typeCoordinates) {
    let coordinateBlock = document.getElementById("coordinateBlock");

    for (let i = 0; i < typeCoordinates.length; i++) {

        if (typeCoordinates[i].texture_object !== "" && !document.getElementById(typeCoordinates[i].texture_object)) {

            let typeBlock = document.createElement("div");
            typeBlock.className = "coordinateType";
            typeBlock.style.background = "url(/assets/map/objects/" + typeCoordinates[i].texture_object + ".png)  center center / 90% no-repeat";

            typeBlock.id = typeCoordinates[i].texture_object;

            typeBlock.onclick = function (){
                PlaceCoordinate("placeObjects");
            };
            coordinateBlock.appendChild(typeBlock);
        }
    }
}

function ViewAnimateObjectsCoordinate(typeCoordinates) {
    let coordinateBlock = document.getElementById("coordinateBlock");

    for (let i = 0; i < typeCoordinates.length; i++) {

        if (typeCoordinates[i].animate_sprite_sheets !== "" && !document.getElementById(typeCoordinates[i].animate_sprite_sheets)) {

            let typeBlock = document.createElement("div");
            typeBlock.className = "coordinateType";

            // TODO анимировать спрайты в беграунде
            // TODO https://medium.com/@vladimirmorulus/%D0%B2%D0%B5%D0%B1-%D0%B0%D0%BD%D0%B8%D0%BC%D0%B0%D1%86%D0%B8%D1%8F-%D0%BD%D0%B0-%D0%BE%D1%81%D0%BD%D0%BE%D0%B2%D0%B5-%D1%81%D0%BF%D1%80%D0%B0%D0%B9%D1%82%D0%BE%D0%B2-8786a9cce59b
            typeBlock.style.background = "url(/assets/map/objects/" + typeCoordinates[i].animate_sprite_sheets + ".png)  center center / 90% no-repeat";

            typeBlock.id = typeCoordinates[i].animate_sprite_sheets;

            typeBlock.onclick = function (){
                PlaceCoordinate("placeAnimate");
            };
            coordinateBlock.appendChild(typeBlock);
        }
    }
}

function tipTypeCoordinate() {
    if (document.getElementById("typeTip")) {
        document.getElementById("typeTip").style.top = stylePositionParams.top + "px";
        document.getElementById("typeTip").style.left = stylePositionParams.left + "px";
    } else {
        CreateTipType(this.coordinateType)
    }
}

function CreateTipType(type) {
    let tip = document.createElement("div");
    tip.id = "typeTip";
    tip.style.top = stylePositionParams.top + "px";
    tip.style.left = stylePositionParams.left + "px";

    let move = "#F00";
    let view = "#F00";
    let attack = "#F00";

    if (type.move) {
        move = "#0F0";
    }

    if (type.view) {
        view = "#0F0";
    }

    if (type.attack) {
        attack = "#0F0";
    }

    tip.innerHTML = "<div><span> Move </span><span style=color:" + move + ">" + type.move + "</span></div>" +
        "<div><span> Watch </span><span style=color:" + view + ">" + type.view + "</span></div>" +
        "<div><span> Attack </span><span style=color:" + attack + ">" + type.attack + "</span></div>";


    document.body.appendChild(tip);
}