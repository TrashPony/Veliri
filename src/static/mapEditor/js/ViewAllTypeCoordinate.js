function ViewAllTypeCoordinate(typeCoordinates) {

    let coordinateBlock = document.getElementById("coordinateBlock");

    for (let i = 0; i < typeCoordinates.length; i++) {
        let typeBlock = document.createElement("div");
        typeBlock.className = "coordinateType";

        if (typeCoordinates[i].texture_object === "") {
            typeBlock.style.background = "url(/assets/map/" + typeCoordinates[i].texture_flore + ".png)  center center / contain no-repeat";
        } else {
            typeBlock.style.background = "url(/assets/map/" + typeCoordinates[i].texture_object + ".png)  center center / contain no-repeat," +
                " url(/assets/map/" + typeCoordinates[i].texture_flore + ".png)  center center / contain no-repeat";
        }

        typeBlock.coordinateType = typeCoordinates[i];
        typeBlock.onmousemove = tipTypeCoordinate;
        typeBlock.onmouseout = function () {
            if (document.getElementById("typeTip")) {
                document.getElementById("typeTip").remove()
            }
        };

        coordinateBlock.appendChild(typeBlock);
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