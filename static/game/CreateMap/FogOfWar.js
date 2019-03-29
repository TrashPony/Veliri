function CreateFowOfWar(coordinate, x, y) {
    // let floorSprite = game.floorLayer.create(x, y, coordinate.texture_flore);
    // floorSprite.anchor.setTo(0.5);
    // floorSprite.scale.set(0.5/2);

    // floorSprite.inputEnabled = true; // включаем ивенты на спрайт
    // floorSprite.input.pixelPerfectOver = true;   // уберает ивенты наведения на пустую зону спрайта
    // floorSprite.input.pixelPerfectClick = true;  // уберает ивенты кликов на пустую зону спрайта

    // floorSprite.events.onInputOut.add(TipOff, floorSprite);
    // floorSprite.z = 0;
    // coordinate.sprite = floorSprite;
    //
    if (game && game.typeService === "battle") {

        let fogSprite = game.fogOfWar.create(x, y, 'FogOfWar');
        fogSprite.anchor.setTo(0.5);
        fogSprite.scale.set(0.25);
        coordinate.fogSprite = fogSprite;
        coordinate.fogSprite = fogSprite;

        game.bmdFogOfWar.draw(fogSprite, x, y);
        fogSprite.destroy();
    }
}