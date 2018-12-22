function initMove(e) {
    if (game.input.activePointer.leftButton.isDown) {
        if (game.squad.toBase) {
            game.squad.toBase.into = false
        }

        global.send(JSON.stringify({
            event: "MoveTo",
            to_x: e.worldX / game.camera.scale.x,
            to_y: e.worldY / game.camera.scale.y
        }));
    }
}