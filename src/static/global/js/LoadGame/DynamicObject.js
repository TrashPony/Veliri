function CreateDynamicObject(jsonData) {
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