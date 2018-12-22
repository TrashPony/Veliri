function FocusMS() {
    game.camera.focusOnXY(game.squad.sprite.x*game.camera.scale.x, game.squad.sprite.y*game.camera.scale.y);
    CreateMiniMap(game.map);
}