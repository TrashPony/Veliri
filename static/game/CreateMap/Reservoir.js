function CreateReservoirs() {
    for (let q in game.map.reservoir) {
        for (let r in game.map.reservoir[q]) {
            let reservoir = game.map.reservoir[q][r];
            CreateReservoir(reservoir, q, r);
        }
    }
}

function CreateReservoir(reservoir, q, r) {
    let xy = GetXYCenterHex(q, r);

    if (game.map.reservoir[q] === undefined || game.map.reservoir[q][r] === undefined) {
        if (game.map.reservoir.hasOwnProperty(q)) {
            game.map.reservoir[q][r] = reservoir;
        } else {
            game.map.reservoir[q] = {};
            game.map.reservoir[q][r] = reservoir;
        }
    }

    let reservoirTexture = reservoir.name;
    let shadow = false;
    let shadowXOffset = 0;
    let shadowYOffset = 0;
    let group = game.floorObjectLayer;

    let full = 100 / ((reservoir.max_count - reservoir.min_count) / (reservoir.count - reservoir.min_count));
    if (full < 34) {

        if (!reservoir.low_move) {
            shadow = true;
            group = game.floorOverObjectLayer;
        }

        reservoirTexture += "_low"
    } else if (full < 67) {

        if (!reservoir.middle_move) {
            shadow = true;
            group = game.floorOverObjectLayer;
            shadowXOffset = -3;
            shadowYOffset = -3;
        }

        reservoirTexture += "_middle"
    } else {

        if (!reservoir.full_move) {
            shadow = true;
            group = game.floorOverObjectLayer;
        }

        reservoirTexture += "_full"
    }

    game.map.reservoir[q][r].sprite = gameObjectCreate(xy.x, xy.y, reservoirTexture, 20, shadow, reservoir.rotate,
        0, 0, group, shadowXOffset, shadowYOffset, 40);

    game.map.reservoir[q][r].sprite.inputEnabled = true;
    game.map.reservoir[q][r].sprite.input.pixelPerfectOver = true;
    game.map.reservoir[q][r].sprite.input.pixelPerfectClick = true;

    let tip;
    let reservoirLine;

    game.map.reservoir[q][r].sprite.events.onInputOver.add(function () {
        reservoirLine = game.floorObjectSelectLineLayer.create(xy.x, xy.y, reservoirTexture);
        reservoirLine.anchor.setTo(0.5);
        reservoirLine.scale.set(0.11);
        reservoirLine.tint = 0x00FF00;
        reservoirLine.angle = reservoir.rotate;

        tip = document.createElement("div");
        tip.id = "reservoirTip" + q + "" + r;
        tip.className = "reservoirTip";
        tip.style.left = stylePositionParams.left + "px";
        tip.style.top = stylePositionParams.top + "px";
        tip.innerHTML = "<h3>" + reservoir.name + "</h3>";
        document.body.appendChild(tip);

        let wrapper = document.createElement("div");
        tip.appendChild(wrapper);

        let icon = document.createElement("div");
        icon.className = "iconOreTip";
        icon.style.background = "url(/assets/resource/" + reservoir.name + ".png)" +
            " center center / contain no-repeat";
        wrapper.appendChild(icon);

        let nameOre = document.createElement("div");
        nameOre.className = "nameOre";
        nameOre.innerHTML = reservoir.name;
        wrapper.appendChild(nameOre);

        let count = document.createElement("div");
        count.className = "countOre";
        count.id = "countOre" + q + "" + r;
        count.innerHTML = game.map.reservoir[q][r].count;
        wrapper.appendChild(count);
    });

    game.map.reservoir[q][r].sprite.events.onInputOut.add(function () {
        tip.remove();
        reservoirLine.destroy()
    });
}