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

    if (reservoir.type === "oil") {
        game.map.reservoir[q][r].sprite = gameObjectCreate(xy.x, xy.y, reservoir.name, 50, false, reservoir.rotate,
            0, 0, game.floorOverObjectLayer)
    } else {
        game.map.reservoir[q][r].sprite = gameObjectCreate(xy.x, xy.y, reservoir.name, 50, true, reservoir.rotate,
            0, 0, game.floorOverObjectLayer);
    }

    game.map.reservoir[q][r].sprite.inputEnabled = true;
    game.map.reservoir[q][r].sprite.input.pixelPerfectOver = true;
    game.map.reservoir[q][r].sprite.input.pixelPerfectClick = true;

    let tip;
    let reservoirLine;

    game.map.reservoir[q][r].sprite.events.onInputOver.add(function () {
        reservoirLine = game.floorObjectSelectLineLayer.create(xy.x, xy.y, reservoir.name);
        reservoirLine.anchor.setTo(0.5);
        reservoirLine.scale.set(0.55);
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