function CreateMap() {
    for (var x = 0; x < game.map.XSize; x++) {
        for (var y = 0; y < game.map.YSize; y++) {

            //game.add.bitmapText(x * tileWidth + tileWidth / 2, y * tileWidth + tileWidth / 2, 'carrier_command', x + ":" + y, 12);
            //console.log(GameMap.OneLayerMap[x][y].type + " " + x + ":" + y);

            if (game.map.OneLayerMap[x][y].type === "" || game.map.OneLayerMap[x][y].type === "respawn") { // пустая клетка или респаун
                var floorSprite = game.add.tileSprite(x * game.tileSize, y * game.tileSize, game.tileSize, game.tileSize, 'floor');
                floorSprite.tint = 0x757575;
                floorSprite.inputEnabled = true; // включаем ивенты на спрайт
                //floorSprite.events.onInputDown.add(SelectTarget, floorSprite);
                floorSprite.events.onInputOut.add(TipOff, floorSprite);
                floorSprite.z = 0;

                game.map.OneLayerMap[x][y].sprite = floorSprite;
            }

            if (game.map.OneLayerMap[x][y].type === "obstacle") { // препятсвие
                var obstacle = game.add.tileSprite(x * game.tileSize, y * game.tileSize, game.tileSize, game.tileSize, 'obstacle');
                obstacle.inputEnabled = true;
                obstacle.events.onInputOut.add(TipOff);

                game.map.OneLayerMap[x][y].sprite = obstacle;
            }
        }
    }
}