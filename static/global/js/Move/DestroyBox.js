function DestroyBox(id, explosionAnimate) {
    for (let i = 0; game.boxes && i < game.boxes.length; i++) {
        if (game.boxes[i].id === id && game.boxes[i].sprite) {

            if (explosionAnimate) {
                let explosion = game.effectsLayer.create(game.boxes[i].sprite.x, game.boxes[i].sprite.y, 'explosion_2');
                explosion.anchor.setTo(0.5);
                explosion.scale.set(0.5);
                explosion.animations.add('explosion_2');
                explosion.animations.play('explosion_2', 10, false, true);
            }

            game.boxes[i].sprite.destroy();
            if (game.boxes[i].shadow) {
                game.boxes[i].shadow.destroy();
            }

            game.boxes.splice(i, 1);
            CreateMiniMap()
        }
    }
}