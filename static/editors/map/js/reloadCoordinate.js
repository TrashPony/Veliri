function ReloadCoordinate(mapPoint) {
    if (mapPoint.coordinate.objectSprite) {
        if (mapPoint.coordinate.objectSprite.shadow) {
            mapPoint.coordinate.objectSprite.shadow.destroy();
        }
        mapPoint.coordinate.objectSprite.destroy();
    }

    if (mapPoint.coordinate.texture_object !== '') {
        CreateObject(mapPoint.coordinate, mapPoint.x, mapPoint.y);
    }

    if (mapPoint.coordinate.animate_sprite_sheets !== '') {
        CreateAnimate(mapPoint.coordinate, mapPoint.x, mapPoint.y);
    }

    CreateLabels(game.map.OneLayerMap[mapPoint.x][mapPoint.y], mapPoint.x, mapPoint.y)
}