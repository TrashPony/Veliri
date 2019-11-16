function UpdateObject(mark, object) {

    if (mark.type_object === "transport") {

    }

    if (mark.type_object === "box") {
    }

    if (mark.type_object === "unit") {
    }

    if (mark.type_object === "reservoir") {
    }

    if (mark.type_object === "dynamic_objects") {
        for (let i in game.objects) {
            if (game.objects[i] && game.objects[i].id === object.id) {

                game.objects[i].scale = object.scale;
                game.objects[i].hp = object.hp;
                game.objects[i].max_hp = object.max_hp;

                let setScale = function (sprite) {
                    game.add.tween(sprite.scale).to(
                        {x: (object.scale / 100) / 2, y: (object.scale / 100) / 2},
                        object.grow_time,
                        Phaser.Easing.Linear.None, true, 0
                    );
                };

                console.log()
                setScale(game.objects[i].objectSprite);
                if (game.objects[i].objectSprite.shadow) {
                    setScale(game.objects[i].objectSprite.shadow);
                    game.add.tween(game.objects[i].objectSprite.shadow).to(
                        {x: object.x + object.x_shadow_offset, y: object.y + object.y_shadow_offset},
                        object.grow_time,
                        Phaser.Easing.Linear.None, true, 0
                    );
                }
            }
        }
    }
}