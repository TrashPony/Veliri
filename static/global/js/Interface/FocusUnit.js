function FocusUnit(id) {
    let unit = game.units[id];

    if (unit && unit.sprite) {
        game.camera.focusOnXY(unit.sprite.x * game.camera.scale.x, unit.sprite.y * game.camera.scale.y);
        game.camera.follow(unit.sprite);
        CreateMiniMap();
        SelectOneUnit(unit, unit.sprite)
    }
}