function Attack() {
    if (!game.squad.AttackLine) {
        // рисует радиус оружия // todo переделать на зону оружия
        game.squad.AttackLine = {
            graphics: game.add.graphics(0, 0),
            diameter: (game.squad.mather_ship.range_view * game.hexagonHeight) * 4,
            visible: false,
        };
        game.squad.sprite.addChild(game.squad.AttackLine.graphics);
        Attack();
    } else {
        game.squad.AttackLine.graphics.clear();

        // если игрок нажал и линия была видна, то значит игрок вЫключил режим атаки и убераем ивенты у всех мсов
        if (game.squad.AttackLine.visible) {
            game.squad.AttackLine.visible = false;
            for (let i = 0; i < game.otherUsers.length; i++) {
                // удаляем всем мсам ивент на клик
                game.otherUsers[i].sprite.unitBody.events.onInputDown.removeAll();
            }
        } else {
            // иначе игрок нажал атаку и вешаем всем мсам ивент для атаки, так же всем обьектам на карте

            // рисуем линию
            game.squad.AttackLine.visible = true;
            game.squad.AttackLine.graphics.lineStyle(3, 0xb74213, 0.2);
            game.squad.AttackLine.graphics.drawCircle(0, 0, game.squad.AttackLine.diameter);
            game.squad.AttackLine.graphics.lineStyle(1, 0xff0000, 1);
            game.squad.AttackLine.graphics.drawCircle(0, 0, game.squad.AttackLine.diameter);

            // TODO переназначить ивент на землю с движения на стрельбу, если игрок стреляет в землю отправлять ивент
            // todo если в игрока отсылать то что в игрока, если в обьект то отсылать что в обьект
            // todo но только 1 ивент
            // todo добавить ховер всем целям что бы было понятно что игрок атакует

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

            // todo ящики
            // todo обьекты на карте с хп
        }
    }
}