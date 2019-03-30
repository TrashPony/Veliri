function CreateFowOfWar(coordinate, x, y) {
    if (game && game.typeService === "battle") {

        if (!game.cache.getBitmapData('FogOfWar')) {
            // заранее кешируем обьект тумана войны
            let bmd = game.add.bitmapData(256, 256);
            bmd.draw(game.make.sprite(0, 0, 'FogOfWar'), 0, 0);
            game.cache.addBitmapData('FogOfWar', bmd);
        }

        let fogSprite = game.make.sprite(0, 0, game.cache.getBitmapData('FogOfWar'));
        fogSprite.anchor.setTo(0.5);
        fogSprite.scale.set(0.25);

        game.bmdFogOfWar.draw(fogSprite, x, y);
        fogSprite.destroy();
    }
}