function CreateEmitter(x, y, minScale, maxScale, minSpeed, maxSpeed, ttl, width, height, color, frequency, minAlpha, maxAlpha, animateSpeed, animate, nameParticle, alphaLoopTime, yoyo) {
    let emitter = game.add.emitter(x, y, 200);
    game.effectsLayer.add(emitter);
    emitter.width = width;
    emitter.height = height;
    emitter.makeParticles(nameParticle);

    emitter.minParticleScale = minScale;
    emitter.maxParticleScale = maxScale;

    emitter.setXSpeed(minSpeed, maxSpeed);
    emitter.setYSpeed(minSpeed, maxSpeed);

    emitter.setRotation(0, 0);
    emitter.setAlpha(minAlpha, maxAlpha, alphaLoopTime, null, yoyo);
    emitter.gravity = 0;

    emitter.start(false, ttl, frequency);

    emitter.forEach(function (singleParticle) {
        if (animate) {
            singleParticle.animations.add('particleAnim');
            singleParticle.animations.play('particleAnim', animateSpeed, true);
        }
        singleParticle.tint = color
    });
}