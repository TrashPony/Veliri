function CreateMyMatherShip() {
    var matherShip = game.add.sprite(0, 0, game.matherShip.type);
    game.physics.arcade.enable(matherShip);
    matherShip.inputEnabled = true;             // включаем ивенты на спрайт

    matherShip.gridX = game.matherShip.x;
    matherShip.gridY = game.matherShip.y;
    matherShip.info = game.matherShip;

    matherShip.events.onInputOver.add(MatherShipTip, matherShip.info); // обрабатываем наведение мышки
    matherShip.events.onInputOut.add(TipOff);   // обрабатываем убирание мышки

    //matherShip.scale.set(1);                  // устанавливаем размер спрайта от оригинала

    var style = {font: "16px Arial", fill: "#00ffff"};

    var label_score = game.add.text(0, 0, game.user.name, style);
    matherShip.addChild(label_score);

    game.map.OneLayerMap[game.matherShip.x][game.matherShip.y].sprite.addChild(matherShip);
}

function CreateHostileMatherShips() {
    for (var x in game.hostileMatherShips) {
        if (game.hostileMatherShips.hasOwnProperty(x)) {
            for (var y in game.hostileMatherShips[x]) {
                if (game.hostileMatherShips[x].hasOwnProperty(y)) {
                    CreateHostileMatherShip(game.hostileMatherShips[x][y]);
                }
            }
        }
    }
}

function CreateHostileMatherShip(matherShipStat) {
    var matherShip = game.add.sprite(0, 0, matherShipStat.type);
    game.physics.arcade.enable(matherShip);
    matherShip.inputEnabled = true;             // включаем ивенты на спрайт

    matherShip.gridX = matherShipStat.x;
    matherShip.gridY = matherShipStat.y;
    matherShip.info = matherShipStat;

    matherShip.events.onInputOver.add(MatherShipTip, matherShip.info); // обрабатываем наведение мышки
    matherShip.events.onInputOut.add(TipOff);   // обрабатываем убирание мышки

    var style = {font: "16px Arial", fill: "#bf0001"};

    var label_score = game.add.text(0, 0, "Hostile", style); // todo matherShipStat.owner ?
    matherShip.addChild(label_score);

    game.map.OneLayerMap[matherShipStat.x][matherShipStat.y].sprite.addChild(matherShip);
}

