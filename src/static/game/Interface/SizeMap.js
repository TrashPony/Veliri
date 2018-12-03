function SizeGameMap(size) {
    if (game) {
        let center = Phaser.Point.add(game.camera.position, new Phaser.Point(game.camera.view.halfWidth, game.camera.view.halfHeight));
        let oldCameraScale = game.camera.scale.clone();

        game.camera.scale.x += size;
        game.camera.scale.y += size;

        let cameraScaleRatio = Phaser.Point.divide(game.camera.scale, oldCameraScale);
        game.camera.focusOn(Phaser.Point.multiply(center, cameraScaleRatio));
    }
}