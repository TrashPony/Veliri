function Attack() {

    game.targetCursorSprite = game.add.sprite(0, 0, 'selectTarget');
    game.targetCursorSprite.scale.setTo(0.25);
    game.targetCursorSprite.anchor.setTo(0.5);
    game.targetCursorSprite.animations.add('select');
    game.targetCursorSprite.animations.play('select', 5, true);

    document.getElementById("GameCanvas").style.cursor = "none";

    let cursorUpdate = setInterval(function () {
        if (game.targetCursorSprite) {
            game.targetCursorSprite.x = ((game.input.mousePointer.x + game.camera.x) / game.camera.scale.x);
            game.targetCursorSprite.y = ((game.input.mousePointer.y + game.camera.y) / game.camera.scale.y);
        } else {
            clearInterval(cursorUpdate);
        }
    }, 10);

    // даем всем динамическим обьектам с хп ивент на атаку
    for (let i in game.objects) {
        if (game.objects[i] && game.objects[i].objectSprite && game.objects[i].hp > -2) {
            game.objects[i].objectSprite.events.onInputDown.add(function (sprite, pointer) {
                // TODO анимация на земле как подтверждение что действие совершилось

                UnselectAttack();
                if (pointer.isMouse && pointer.button === 0) {
                    global.send(JSON.stringify({
                        event: "Attack",
                        type: "object",
                        x: Math.round(game.objects[i].objectSprite.x),
                        y: Math.round(game.objects[i].objectSprite.y),
                        units_id: getIDsSelectUnits(),
                    }));
                }
            })
        }
    }

    // ящики
    for (let i in game.boxes) {
        if (game.boxes[i] && game.boxes[i].sprite) {
            game.boxes[i].sprite.events.onInputDown.add(function (sprite, pointer) {
                // TODO анимация на земле как подтверждение что действие совершилось
                UnselectAttack();
                if (pointer.isMouse && pointer.button === 0) {
                    global.send(JSON.stringify({
                        event: "Attack",
                        type: "box",
                        box_id: game.boxes[i].id,
                        units_id: getIDsSelectUnits(),
                    }));
                }
            })
        }
    }

    // юниты
    for (let i in game.units) {
        let unit = game.units[i];
        if (unit && unit.sprite && unit.sprite.unitBody) {
            unit.sprite.unitBody.events.onInputDown.add(function (sprite, pointer) {
                // TODO анимация на земле как подтверждение что действие совершилось

                UnselectAttack();
                if (pointer.isMouse && pointer.button === 0) {
                    global.send(JSON.stringify({
                        event: "Attack",
                        type: "unit",
                        unit_id: unit.id,
                        units_id: getIDsSelectUnits(),
                    }));
                }
            })
        }
    }

    // атака просто в карту
    game.bmdTerrain.sprite.events.onInputDown.add(function (sprite, pointer) {
        // TODO анимация на земле как подтверждение что действие совершилось

        UnselectAttack();
        dontMove = true; // TODO костыль из за которого надо делать еще 1 нажатие

        if (pointer.isMouse && pointer.button === 0) {
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