function Explosion(x,y) {
    let explosion = game.effectsLayer.create(x, y, 'explosion_1');
    explosion.anchor.setTo(0.5);
    explosion.animations.add('explosion_1');
    explosion.animations.play('explosion_1', 20, false, true);
}