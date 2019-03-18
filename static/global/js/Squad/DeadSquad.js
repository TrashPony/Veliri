function SquadDead(user) {

    let explosion = game.effectsLayer.create(0, 0, 'explosion_2');
    explosion.anchor.setTo(0.5);
    explosion.scale.set(1.5);


    if (game.squad && Number(user.squad_id) === game.squad.id) {
        explosion.x = game.squad.sprite.x;
        explosion.y = game.squad.sprite.y;
        game.squad.sprite.kill();
    } else {
        for (let i = 0; game.otherUsers && i < game.otherUsers.length; i++) {
            if (game.otherUsers[i].squad_id === user.squad_id) {
                explosion.x = game.otherUsers[i].sprite.x;
                explosion.y = game.otherUsers[i].sprite.y;
                game.otherUsers[i].kill();
            }
        }
    }

    explosion.animations.add('explosion_2');
    explosion.animations.play('explosion_2', 10, false, true);
}