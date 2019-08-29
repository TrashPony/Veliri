function FocusMS() {
    game.camera.focusOnXY(game.my_squad_sprite.x*game.camera.scale.x, game.my_squad_sprite.y*game.camera.scale.y);
    game.camera.follow(game.my_squad_sprite);
    CreateMiniMap(game.map);
}