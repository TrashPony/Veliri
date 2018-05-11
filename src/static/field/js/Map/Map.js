function CreateMap(GameMap) {
    for (var x = 0; x < GameMap.XSize; x++) {
        for (var y = 0; y < GameMap.YSize; y++) {

            //game.add.bitmapText(x * tileWidth + tileWidth / 2, y * tileWidth + tileWidth / 2, 'carrier_command', x + ":" + y, 12);
            //console.log(GameMap.OneLayerMap[x][y].type + " " + x + ":" + y);

            if (GameMap.OneLayerMap[x][y].type === "" || GameMap.OneLayerMap[x][y].type === "respawn") { // пустая клетка или респаун
                var floorSprite = game.add.tileSprite(x * tileWidth, y * tileWidth, tileWidth, tileWidth, 'floor');
                floorSprite.tint = 0x757575;
                floorSprite.inputEnabled = true; // включаем ивенты на спрайт
                floorSprite.events.onInputDown.add(SelectTarget, this);
                floorSprite.events.onInputOut.add(mouse_out, this);
                floorSprite.z = 0;

                GameMap.OneLayerMap[x][y].sprite = floorSprite;
            }

            if (GameMap.OneLayerMap[x][y].type === "obstacle") { // препятсвие
                var obstacle = game.add.tileSprite(x * tileWidth, y * tileWidth, tileWidth, tileWidth, 'obstacle');
                obstacle.inputEnabled = true;
                obstacle.events.onInputOut.add(mouse_out);

                GameMap.OneLayerMap[x][y].sprite = obstacle;
            }
        }
    }
}