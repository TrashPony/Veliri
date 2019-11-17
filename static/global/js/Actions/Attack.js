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

    for (let i in game.objects) {
        if (game.objects[i] && game.objects[i].objectSprite && game.objects[i].hp > -2) {
            game.objects[i].objectSprite.events.onInputDown.add(function () {
                console.log(1)
            })
        }
    }

    game.input.onDown.add(function () {
        console.log(2)
        // TODO анимация на земле как подтверждение что действие совершилось

        targetCursorSprite.destroy();
        UnselectAttack();

        if (game.input.activePointer.leftButton.isDown) {
            global.send(JSON.stringify({
                event: "Attack",
                type: "map",
                x: Math.round((game.input.mousePointer.x + game.camera.x) / game.camera.scale.x),
                y: Math.round((game.input.mousePointer.y + game.camera.y) / game.camera.scale.y),
                units_id: getIDsSelectUnits(),
            }));
        }
    });
}