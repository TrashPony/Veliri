function ViewPatternCoordinate(typeCoordinates) {

    let coordinateBlock = document.getElementById("coordinateBlock");

    for (let i = 0; i < typeCoordinates.length; i++) {
        let typeBlock = document.createElement("div");
        typeBlock.className = "coordinateType";

        if (typeCoordinates[i].animate_sprite_sheets !== "") {

            typeBlock.style.background = "url(/assets/map/animate/" + typeCoordinates[i].animate_sprite_sheets + ".png)  center center / contain no-repeat"

        } else if (typeCoordinates[i].texture_object.split('_').length > 0 && typeCoordinates[i].texture_object.split('_')[0] === "mountain") {

            typeBlock.style.background = "url(/assets/map/objects/mountains/" + typeCoordinates[i].texture_object + ".png)  center center / 90% no-repeat"

        } else if (typeCoordinates[i].texture_object.split('_').length > 0 && typeCoordinates[i].texture_object.split('_')[0] === "plant") {

            typeBlock.style.background = "url(/assets/map/objects/plants/" + typeCoordinates[i].texture_object + ".png)  center center / 90% no-repeat"

        } else if (typeCoordinates[i].texture_object.split('_').length > 0 && typeCoordinates[i].texture_object.split('_')[0] === "ravine") {

            typeBlock.style.background = "url(/assets/map/objects/ravines/" + typeCoordinates[i].texture_object + ".png)  center center / 90% no-repeat"

        } else if (typeCoordinates[i].texture_object.split('_').length > 0 && typeCoordinates[i].texture_object.split('_')[0] === "road") {

            typeBlock.style.background = "url(/assets/map/objects/roads/" + typeCoordinates[i].texture_object + ".png)  center center / 90% no-repeat"

        } else if (typeCoordinates[i].texture_object.split('_').length > 0 && typeCoordinates[i].texture_object.split('_')[1] === "base") {

            typeBlock.style.background = "url(/assets/map/objects/bases/" + typeCoordinates[i].texture_object + ".png)  center center / 90% no-repeat"

        } else {
            typeBlock.style.background = "url(/assets/map/objects/" + typeCoordinates[i].texture_object + ".png)  center center / 90% no-repeat"
        }

        if (typeCoordinates[i].type === "respawn") {
            typeBlock.innerText = "RESPAWN";
        }

        typeBlock.coordinateType = typeCoordinates[i];

        typeBlock.onclick = function () {
            PlaceCoordinate("placeCoordinate", typeCoordinates[i]);
        };

        let menuBlock = document.createElement("div");
        menuBlock.className = "menuButton";
        menuBlock.onclick = function () {
            CreateSubMenu(typeCoordinates[i])
        };

        typeBlock.appendChild(menuBlock);
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
            typeBlock.coordinateType = typeCoordinates[i];

            typeBlock.onclick = function () {
                PlaceCoordinate("placeTerrain", typeCoordinates[i]);
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
            typeBlock.coordinateType = typeCoordinates[i];

            typeBlock.onclick = function () {
                PlaceCoordinate("placeObjects", typeCoordinates[i]);
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
            typeBlock.style.background = "url(/assets/map/animate/" + typeCoordinates[i].animate_sprite_sheets + ".png)  center center / 90% no-repeat";

            typeBlock.id = typeCoordinates[i].animate_sprite_sheets;
            typeBlock.coordinateType = typeCoordinates[i];

            typeBlock.onclick = function () {
                PlaceCoordinate("placeAnimate", typeCoordinates[i]);
            };

            coordinateBlock.appendChild(typeBlock);
        }
    }
}