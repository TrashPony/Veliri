function SizeGameMap(size) {
    if (game) {
        CreateMiniMap(game.map);
        let center = Phaser.Point.add(game.camera.position, new Phaser.Point(game.camera.view.halfWidth, game.camera.view.halfHeight));
        let oldCameraScale = game.camera.scale.clone();

        game.camera.scale.x += size;
        game.camera.scale.y += size;

        if (game.camera.scale.x < 0.50) {
            game.camera.scale.x = 0.5;
            game.camera.scale.y = 0.5;
        } else if(game.camera.scale.x > 2) {
            game.camera.scale.x = 2;
            game.camera.scale.y = 2;
        }

        let cameraScaleRatio = Phaser.Point.divide(game.camera.scale, oldCameraScale);
        game.camera.focusOn(Phaser.Point.multiply(center, cameraScaleRatio));
    }
}