function SizeGameMap(size) {
    if (game) {
        game.camera.scale.x += size;
        game.camera.scale.y += size;
    }
}