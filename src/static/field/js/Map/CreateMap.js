function CreateMap() {
    for (var x = 0; x < game.map.XSize; x++) {
        for (var y = 0; y < game.map.YSize; y++) {

            var floorSprite = game.floorLayer.create(x * game.tileSize, y * game.tileSize, 'floor');
            floorSprite.tint = 0x757575;
            floorSprite.inputEnabled = true; // включаем ивенты на спрайт
            floorSprite.events.onInputOut.add(TipOff, floorSprite);
            floorSprite.z = 0;

            game.map.OneLayerMap[x][y].sprite = floorSprite;

            if (game.map.OneLayerMap[x][y].texture_object !== "") {

                var shadow;
                var object;

                if (game.map.OneLayerMap[x][y].texture_object === "terrain_1") {
                    object = gameObjectCreate(x, y, game.map.OneLayerMap[x][y].texture_object, -0.33, 0.70);
                }

                if (game.map.OneLayerMap[x][y].texture_object === "terrain_2") {
                    object = gameObjectCreate(x, y, game.map.OneLayerMap[x][y].texture_object, -0.42, 0.75);
                }

                if (game.map.OneLayerMap[x][y].texture_object === "wall") {
                    shadow = game.floorObjectLayer.create((x * game.tileSize), (y * game.tileSize) + game.shadowYOffset, game.map.OneLayerMap[x][y].texture_object);
                    shadow.anchor.setTo(0, 0);
                    shadow.tint = 0x000000;
                    shadow.alpha = 0.6;

                    object = game.floorObjectLayer.create(x * game.tileSize, y * game.tileSize, game.map.OneLayerMap[x][y].texture_object);
                    object.inputEnabled = true;
                    object.events.onInputOut.add(TipOff);
                    //object.anchor.setTo(0, 0);
                }


                game.map.OneLayerMap[x][y].objectSprite = object;
            }
        }
    }
}

function gameObjectCreate(x, y, texture, ShadowOffsetX, ShadowOffsetY) {
    var shadow = game.floorObjectLayer.create((x * game.tileSize), (y * game.tileSize) + game.shadowYOffset, texture);
    shadow.anchor.setTo(ShadowOffsetX, ShadowOffsetY);
    shadow.tint = 0x000000;
    shadow.alpha = 0.6;
    shadow.angle = 45;

    var object = game.floorObjectLayer.create(x * game.tileSize, y * game.tileSize, texture);
    object.inputEnabled = true;
    object.events.onInputOut.add(TipOff);
    object.anchor.setTo(0, 0.2);

    object.shadow = shadow;

    return object
}