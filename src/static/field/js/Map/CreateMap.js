function CreateMap() {
    for (var x = 0; x < game.map.XSize; x++) {
        for (var y = 0; y < game.map.YSize; y++) {
            // todo построение карты из существующих координат а не по размеру
            var floorSprite = game.floorLayer.create(x * game.tileSize, y * game.tileSize, 'floor');
            var fogSprite = game.fogOfWar.create(x * game.tileSize, y * game.tileSize, 'FogOfWar');

            floorSprite.inputEnabled = true; // включаем ивенты на спрайт
            floorSprite.events.onInputOut.add(TipOff, floorSprite);
            floorSprite.events.onInputDown.add(RemoveSelect);
            floorSprite.z = 0;

            game.map.OneLayerMap[x][y].sprite = floorSprite;
            game.map.OneLayerMap[x][y].fogSprite = fogSprite;

            if (game.map.OneLayerMap[x][y].texture_object !== "") {

                var object;

                if (game.map.OneLayerMap[x][y].texture_object === "terrain_1") {
                    object = gameObjectCreate(x, y, game.map.OneLayerMap[x][y].texture_object, 0, 0.2, -0.33, 0.70, 45);
                }

                if (game.map.OneLayerMap[x][y].texture_object === "terrain_2") {
                    object = gameObjectCreate(x, y, game.map.OneLayerMap[x][y].texture_object, 0, 0.2, -0.42, 0.75, 45);
                }

                if (game.map.OneLayerMap[x][y].texture_object === "terrain_3") {
                    object = gameObjectCreate(x, y, game.map.OneLayerMap[x][y].texture_object, 0, 0, -0.5, 0.8, 45);
                }

                if (game.map.OneLayerMap[x][y].texture_object === "wall") {
                    object = gameObjectCreate(x, y, game.map.OneLayerMap[x][y].texture_object, 0, 0.2, -0.1, 0.35, 10);
                }

                if (game.map.OneLayerMap[x][y].texture_object === "crater") {
                    object = game.floorObjectLayer.create(x * game.tileSize, y * game.tileSize, game.map.OneLayerMap[x][y].texture_object);
                    object.inputEnabled = true;
                    object.events.onInputOut.add(TipOff);
                    object.events.onInputDown.add(RemoveSelect);
                    object.input.pixelPerfectOver = true;
                    object.input.pixelPerfectClick = true;
                }

                game.map.OneLayerMap[x][y].objectSprite = object;
            }

            if (game.map.OneLayerMap[x][y].effects != null && game.map.OneLayerMap[x][y].effects.length > 0) {
                MarkZoneEffect(game.map.OneLayerMap[x][y])
            }
        }
    }
}

function gameObjectCreate(x, y, texture, spriteAnchorX, spriteAnchorY, ShadowOffsetX, ShadowOffsetY, angle) {


    var shadow = game.floorObjectLayer.create((x * game.tileSize), (y * game.tileSize) + game.shadowYOffset, texture);
    shadow.anchor.setTo(ShadowOffsetX, ShadowOffsetY);
    shadow.tint = 0x000000;
    shadow.alpha = 0.6;
    shadow.angle = angle;

    shadow.inputEnabled = true;             // включаем ивенты на спрайт
    shadow.input.pixelPerfectOver = true;
    shadow.input.pixelPerfectClick = true;

    var object = game.floorObjectLayer.create(x * game.tileSize, y * game.tileSize, texture);
    object.inputEnabled = true;
    object.events.onInputOut.add(TipOff);
    object.events.onInputDown.add(RemoveSelect);
    object.anchor.setTo(spriteAnchorX, spriteAnchorY); // todo подгружать спрайт от невидимой границы

    shadow.inputEnabled = true;
    object.input.pixelPerfectOver = true;
    object.input.pixelPerfectClick = true;

    object.shadow = shadow;

    return object
}