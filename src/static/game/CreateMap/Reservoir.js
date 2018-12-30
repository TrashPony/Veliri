function CreateReservoir() {
    console.log(game.map.reservoir)

    for (let q in game.map.reservoir) {
        for (let r in game.map.reservoir[q]) {

            let reservoir = game.map.reservoir[q][r];
            let xy = GetXYCenterHex(q, r);

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
                tip.id = "reservoirTip";
                tip.style.left = stylePositionParams.left + "px";
                tip.style.top = stylePositionParams.top + "px";
                document.body.appendChild(tip)

                // TODO заполнение типа
            });

            game.map.reservoir[q][r].sprite.events.onInputOut.add(function () {
                tip.remove();
                reservoirLine.destroy()
            });

            // TODO ивент при тыке
        }
    }
}