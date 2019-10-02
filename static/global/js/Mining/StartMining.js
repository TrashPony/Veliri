function StartMining(jsonData) {

    let attachPoint = GetSpriteEqip(jsonData.short_unit.id, jsonData.type_slot, jsonData.slot).attachPoint;
    let equipSprite = GetSpriteEqip(jsonData.short_unit.id, jsonData.type_slot, jsonData.slot).sprite;

    let xy = {x: jsonData.x, y: jsonData.y};
    let laserOut = game.add.graphics(0, 0);
    laserOut.lineStyle(6, 0xFFEDFF, 1);

    let laserIn = game.add.graphics(0, 0);
    laserIn.lineStyle(2, 0xFFFFFF, 1);

    let blurX = game.add.filter('BlurX');
    let blurY = game.add.filter('BlurY');
    blurX.blur = 2;
    blurY.blur = 2;
    laserOut.filters = [blurX, blurY];
    blurX.blur = 1;
    blurY.blur = 1;
    laserIn.filters = [blurX, blurY];

    attachPoint.addChild(laserOut);
    attachPoint.addChild(laserIn);

    setTimeout(function () {
        laserOut.destroy();
        laserIn.destroy();
    }, jsonData.seconds * 1000 - 3000);

    let unit = game.units[jsonData.short_unit.id];

    if (unit) {
        //"reloadEquip" + unit.id + type + i
        let id = "reloadEquip" + unit.id + jsonData.type_slot + jsonData.slot;

        if (!unit.miningLaser) {
            unit.miningLaser = [];
        }

        unit.miningLaser.push({
            out: laserOut,
            in: laserIn,
            xy: xy,
            id: id,
            attachPoint: attachPoint,
            equipSprite: equipSprite,
        });

        let tween = game.add.tween(xy).to({
                x: xy.x - 8 + 15,
                y: xy.y - 8 + 15
            }, 1000, Phaser.Easing.Linear.None, true, 0
        ).loop(true);
        tween.yoyo(true, 0);

        let progressBar = document.getElementById(id);
        if (progressBar) {
            progressBar.style.animation = "none";
            setTimeout(function () {
                progressBar.style.animation = "reload " + jsonData.seconds + "s linear 1";
            }, 10);
        }
    }
}

function InitMiningOre(unitID, numberSlot, type, equip) {
    UnselectResource();

    let unit = game.units[unitID];
    let graphics = game.add.graphics(0, 0);
    unit.selectMiningLine = {graphics: graphics, radius: equip.equip.radius};
    game.floorObjectLayer.add(graphics);

    for (let x in game.map.reservoir) {
        for (let y in game.map.reservoir[x]) {

            let reservoir = game.map.reservoir[x][y];

            reservoir.sprite.events.onInputDown.add(function () {
                global.send(JSON.stringify({
                    event: "startMining",
                    unit_id: unitID,
                    slot: Number(numberSlot),
                    type_slot: type,
                    x: reservoir.x,
                    y: reservoir.y,
                }));
                UnselectResource()
            });
        }
    }
}