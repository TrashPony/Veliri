function UseDigger(jsonData) {
    console.log(jsonData)

    game.floorObjectSelectLineLayer.forEach(function (sprite) {
        sprite.visible = false;
    });

    if (jsonData.dynamic_object && jsonData.dynamic_object.texture_object !== '') {
        let xy = GetXYCenterHex(jsonData.q, jsonData.r);
        let paket = game.floorObjectLayer.create(xy.x, xy.y, jsonData.dynamic_object.texture_object);
        paket.scale.set(0.20);
        paket.anchor.setTo(0.5);

        paket.inputEnabled = true;
        paket.events.onInputDown.add(function () {
            anomalyText(jsonData.dynamic_object.dialog, jsonData.dynamic_object.dialog.pages[0])
        });

        let paketLine;
        paket.events.onInputOver.add(function () {
            paketLine = game.floorObjectSelectLineLayer.create(xy.x, xy.y, 'infoAnomaly');
            paketLine.anchor.setTo(0.5);
            paketLine.scale.set(0.22);
            paketLine.tint = 0x00FF00;
        });
        paket.events.onInputOut.add(function () {
            paketLine.destroy();
        });
        paket.input.priorityID = 1;
    }

    if (jsonData.box) {
        CreateBox(jsonData.box);
    }

    if (jsonData.reservoir) {
        CreateReservoir(jsonData.reservoir, jsonData.reservoir.Q, jsonData.reservoir.R);
    }
}

function anomalyText(text, locateText) {
    Alert(locateText.text, "Неизвестная запись<br>", false, 0, false, "anomalyText");
    let anomalyTextBlock = $('#anomalyText');

    for (let i in locateText.asc) {
        let asc = $('<div></div>');
        asc.addClass('Ask');
        asc.text(locateText.asc[i].text);
        asc.click(function () {
            if (locateText.asc[i].to_page === 0) {
                anomalyTextBlock.remove();
            } else {
                anomalyTextBlock.remove();
                anomalyText(text, text.pages[locateText.asc[i].to_page - 1])
            }
        });
        anomalyTextBlock.append(asc);
    }
}