function Attack() {
    if (!game.squad.AttackLine) {
        game.squad.AttackLine = {
            graphics: game.add.graphics(0, 0),
            diameter: (game.squad.mather_ship.range_view * game.hexagonHeight) * 4,
            visible: false,
        };
        game.squad.sprite.addChild(game.squad.AttackLine.graphics);
        Attack();
    } else {
        game.squad.AttackLine.graphics.clear();

        if (game.squad.AttackLine.visible) {
            game.squad.AttackLine.visible = false;
            for (let i = 0; i < game.otherUsers.length; i++) {
                // удаляем всем мсам ивент на клик
                game.otherUsers[i].sprite.unitBody.events.onInputDown.removeAll();
            }
        } else {

            game.squad.AttackLine.visible = true;
            game.squad.AttackLine.graphics.lineStyle(3, 0xb74213, 0.2);
            game.squad.AttackLine.graphics.drawCircle(0, 0, game.squad.AttackLine.diameter);
            game.squad.AttackLine.graphics.lineStyle(1, 0xff0000, 1);
            game.squad.AttackLine.graphics.drawCircle(0, 0, game.squad.AttackLine.diameter);

            for (let i = 0; i < game.otherUsers.length; i++) {
                if (!game.otherUsers[i].sprite) continue;
                console.log(game.otherUsers[i]);
                game.otherUsers[i].sprite.unitBody.inputEnabled = true;
                game.otherUsers[i].sprite.unitBody.input.pixelPerfectOver = true;
                game.otherUsers[i].sprite.unitBody.input.pixelPerfectClick = true;
                game.otherUsers[i].sprite.unitBody.events.onInputDown.add(function () {
                    global.send(JSON.stringify({
                        event: "Attack",
                        to_squad_id: game.otherUsers[i].squad_id,
                    }));
                })
            }
        }
    }
}