function StartMining(jsonData) {

    let q = jsonData.q;
    let r = jsonData.r;

    let xy = GetXYCenterHex(q, r);

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

    game.floorObjectLayer.add(laserOut);
    game.floorObjectLayer.add(laserIn);

    setTimeout(function () {
        laserOut.destroy();
        laserIn.destroy();
    }, jsonData.seconds * 1000 - 3000);

    if (Number(jsonData.other_user.squad_id) === game.squad.id) {

        if (!game.squad.miningLaser) {
            game.squad.miningLaser = [];
        }

        let progressBar = document.getElementById("miningEquip" + jsonData.type_slot + jsonData.slot);
        progressBar.style.animation = "none";
        setTimeout(function () {
            progressBar.style.animation = "reload " + jsonData.seconds + "s linear 1";
        }, 10);

        game.squad.miningLaser.push({
            out: laserOut,
            in: laserIn,
            xy: xy,
            id: "miningEquip" + jsonData.type_slot + "" + jsonData.slot
        });
        let tween = game.add.tween(xy).to({
                x: xy.x - 8 + 15,
                y: xy.y - 8 + 15
            }, 1000, Phaser.Easing.Linear.None, true, 0
        ).loop(true);
        tween.yoyo(true, 0);
    } else {
        for (let i = 0; game.otherUsers && i < game.otherUsers.length; i++) {
            if (game.otherUsers[i].user_name === jsonData.other_user.user_name) {

                if (!game.otherUsers[i].miningLaser) {
                    game.otherUsers[i].miningLaser = [];
                }

                game.otherUsers[i].miningLaser.push({
                    out: laserOut, in: laserIn, xy: xy, id: game.otherUsers[i].user_name +
                        "miningEquip" + jsonData.type_slot + "" + jsonData.slot
                });
                let tween = game.add.tween(xy).to({
                        x: xy.x - 15 + 30,
                        y: xy.y - 15 + 30
                    }, 1000, Phaser.Easing.Linear.None, true, 0
                ).loop(true);
                tween.yoyo(true, 0);
            }
        }
    }
}

function InitMiningOre(equip, numberSlot, type) {

    function unselectMiningOre() {
        for (let q in game.map.reservoir) {
            for (let r in game.map.reservoir[q]) {
                game.map.reservoir[q][r].sprite.events.onInputDown.removeAll();
                game.map.reservoir[q][r].reservoirLine.destroy()
            }
        }
        game.squad.selectMiningLine.graphics.destroy();
        game.input.onDown.add(initMove, game);
    }

    let graphics = game.add.graphics(0, 0);
    game.squad.selectMiningLine = {graphics: graphics, radius: equip.radius * 200};
    game.floorObjectLayer.add(graphics);

    for (let q in game.map.reservoir) {
        for (let r in game.map.reservoir[q]) {
            let reservoir = game.map.reservoir[q][r];
            let reservoirLine = game.floorObjectSelectLineLayer.create(reservoir.sprite.x, reservoir.sprite.y, reservoir.name);
            reservoirLine.anchor.setTo(0.5);
            reservoirLine.scale.set(0.27);
            reservoirLine.tint = 0x0FFF00;
            reservoirLine.angle = reservoir.rotate;

            reservoir.sprite.input.priorityID = 1;
            reservoir.reservoirLine = reservoirLine;

            game.input.onDown.remove(initMove, game);
            game.input.onDown.add(function () {
                if (game.input.activePointer.rightButton.isDown) {
                    unselectMiningOre()
                }
            });
            reservoir.sprite.events.onInputDown.add(function () {
                global.send(JSON.stringify({
                    event: "startMining",
                    slot: Number(numberSlot),
                    q: reservoir.q,
                    r: reservoir.r,
                    type_slot: type
                }));
                unselectMiningOre()
            });
        }
    }
}

function InitDigger(equip, numberSlot, type) {
    global.send(JSON.stringify({
        event: "SelectDigger",
        slot: Number(numberSlot),
        type_slot: type
    }));
}