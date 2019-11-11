function CreateReservoirs() {
    for (let x in game.map.reservoir) {
        for (let y in game.map.reservoir[x]) {
            let reservoir = game.map.reservoir[x][y];
            CreateReservoir(reservoir, Number(x), Number(y));
        }
    }
}

function CreateReservoir(reservoir, x, y) {

    if (!game.map.reservoir[x] || !game.map.reservoir[x][y]) {
        if (game.map.reservoir.hasOwnProperty(x)) {
            game.map.reservoir[x][y] = reservoir;
        } else {
            game.map.reservoir[x] = {};
            game.map.reservoir[x][y] = reservoir;
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

    game.map.reservoir[x][y].sprite = gameObjectCreate(x, y, reservoirTexture, 20, shadow, reservoir.rotate,
        group, shadowXOffset, shadowYOffset, 40);

    game.map.reservoir[x][y].sprite.inputEnabled = true;
    game.map.reservoir[x][y].sprite.input.pixelPerfectOver = true;
    game.map.reservoir[x][y].sprite.input.pixelPerfectClick = true;
    game.map.reservoir[x][y].sprite.input.pixelPerfectAlpha = 1;

    let tip;
    let posInterval;
    game.map.reservoir[x][y].sprite.events.onInputOver.add(function () {

        if (!game.map.reservoir[x][y].border) {
            game.map.reservoir[x][y].border = CreateBorder(x, y, reservoirTexture, 20, reservoir.rotate, group);
            group.swap(game.map.reservoir[x][y].sprite, game.map.reservoir[x][y].border);
        } else {
            game.map.reservoir[x][y].border.visible = true;
        }

        tip = document.createElement("div");
        tip.id = "reservoirTip" + x + "" + y;
        tip.className = "reservoirTip";
        tip.style.left = stylePositionParams.left + "px";
        tip.style.top = stylePositionParams.top + "px";
        document.body.appendChild(tip);

        tip.innerHTML = `
            <h3>${reservoir.name}</h3>
            <div class="Description"> Залежи минерала ${reservoir.name}. Лежат тут, никто их не трогает...</div>
            <div class="reservoirInfo">
                <div class="iconOreTip" style="background: url('/assets/resource/${reservoir.name}.png') center center / contain no-repeat"></div>
                <div class="nameOre">${reservoir.name}</div>
                <div class="countOre" id="countOre${x}${y}">${game.map.reservoir[x][y].count}</div>
            </div>
        `;

        posInterval = setInterval(function () {
            tip.style.left = stylePositionParams.left + "px";
            tip.style.top = stylePositionParams.top + "px";
        }, 10)
    });

    game.map.reservoir[x][y].sprite.events.onInputOut.add(function () {
        if (game.map.reservoir[x][y].border) game.map.reservoir[x][y].border.visible = false;
        setInterval(posInterval);
        if (tip) tip.remove();
    });
}