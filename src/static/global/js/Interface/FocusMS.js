function FocusMS() {
    game.camera.focusOnXY(game.squad.sprite.x*game.camera.scale.x, game.squad.sprite.y*game.camera.scale.y);
    game.camera.follow(game.squad.sprite);
    CreateMiniMap(game.map);
}