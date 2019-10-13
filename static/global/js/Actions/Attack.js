function Attack() {

    let targetCursorSprite = game.add.sprite(0, 0, 'selectTarget');
    targetCursorSprite.scale.setTo(0.25);
    targetCursorSprite.anchor.setTo(0.5);
    targetCursorSprite.animations.add('select');
    targetCursorSprite.animations.play('select', 5, true);

    document.getElementById("GameCanvas").style.cursor = "none";

    setInterval(function () {
        targetCursorSprite.x = ((game.input.mousePointer.x + game.camera.x) / game.camera.scale.x);
        targetCursorSprite.y = ((game.input.mousePointer.y + game.camera.y) / game.camera.scale.y);
    }, 10);

    game.input.onDown.add(function () {

        dontMove = true;
        document.getElementById("GameCanvas").style.cursor = "unset";
        targetCursorSprite.destroy();

        // TODO анимация на земле как подтверждение что действие совершилось

        let x = (game.input.mousePointer.x + game.camera.x) / game.camera.scale.x;
        let y = (game.input.mousePointer.y + game.camera.y) / game.camera.scale.y;

        UnselectAttack();

        global.send(JSON.stringify({
            event: "Attack",
            type: "map",
            x: Math.round(x),
            y: Math.round(y),
            units_id: getIDsSelectUnits(),
        }));
    });
}