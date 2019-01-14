function UseDigger(jsonData) {
    console.log(jsonData)

    game.floorObjectSelectLineLayer.forEach(function (sprite) {
        sprite.visible = false;
    });

    if (jsonData.anomaly_text) {
        let xy = GetXYCenterHex(jsonData.q, jsonData.r);
        let paket = game.floorObjectLayer.create(xy.x, xy.y, 'infoAnomaly');
        paket.scale.set(0.20);
        paket.anchor.setTo(0.5);

        paket.inputEnabled = true;
        paket.events.onInputDown.add(function () {
            anomalyText(jsonData.anomaly_text, jsonData.anomaly_text.Pages[0])
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
}

function anomalyText(text, locateText) {
    Alert(locateText.Text, "Неизвестная запись<br>", false, 0, false, "anomalyText");
    let anomalyTextBlock = $('#anomalyText');

    for (let i in locateText.Asc) {
        let asc = $('<div></div>');
        asc.addClass('Ask');
        asc.text(locateText.Asc[i].Text);
        asc.click(function () {
           if (locateText.Asc[i].ToPage === 0) {
               anomalyTextBlock.remove();
           } else {
               anomalyTextBlock.remove();
               anomalyText(text, text.Pages[locateText.Asc[i].ToPage-1])
           }
        });
        anomalyTextBlock.append(asc);
    }
}