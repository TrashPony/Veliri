function CreateTerrain(coordinate, x, y, q, r) {

    let floorSprite = game.floorLayer.create(x, y, coordinate.texture_flore);
    floorSprite.anchor.setTo(0.5);
    floorSprite.scale.set(0.5);

    floorSprite.inputEnabled = true; // включаем ивенты на спрайт
    floorSprite.input.pixelPerfectOver = true;   // уберает ивенты наведения на пустую зону спрайта
    floorSprite.input.pixelPerfectClick = true;  // уберает ивенты кликов на пустую зону спрайта

    floorSprite.events.onInputOut.add(TipOff, floorSprite);
    floorSprite.z = 0;
    coordinate.sprite = floorSprite;

    if (game && game.typeService === "battle") {
        let fogSprite = game.fogOfWar.create(x, y, 'FogOfWar');
        fogSprite.anchor.setTo(0.5);
        fogSprite.scale.set(0.5);
        floorSprite.events.onInputDown.add(RemoveSelect);
        coordinate.fogSprite = fogSprite;
        coordinate.fogSprite = fogSprite;
    }

    coordinate.coordinateText = {};

    if (game && game.typeService !== "battle" && game.typeService !== "global") {
        let style = {font: "12px Arial", fill: "#606060", align: "center"};
        coordinate.coordinateText.qr = game.add.text(x - 10, y - 40, q + "," + r, style, game.redactorMetaText);

        if (metaAlpha === 0){
            coordinate.coordinateText.qr.alpha = metaAlpha;
        }

        let allow = {font: "12px Arial", fill: "#150bff", align: "center"};
        let noAllow = {font: "12px Arial", fill: "#ff2821", align: "center"};

        if (!(coordinate.move && coordinate.view && coordinate.attack)) {
            if (coordinate.move) {
                coordinate.coordinateText.move = game.add.text(x - 40, y - 25, 'm', allow, game.redactorMetaText);
            } else {
                coordinate.coordinateText.move = game.add.text(x - 40, y - 25, 'm', noAllow, game.redactorMetaText);
            }

            if (coordinate.view) {
                coordinate.coordinateText.view = game.add.text(x - 25, y - 25, 'w', allow, game.redactorMetaText);
            } else {
                coordinate.coordinateText.view = game.add.text(x - 25, y - 25, 'w', noAllow, game.redactorMetaText);
            }

            if (coordinate.attack) {
                coordinate.coordinateText.attack = game.add.text(x - 10, y - 25, 'a', allow, game.redactorMetaText);
            } else {
                coordinate.coordinateText.attack = game.add.text(x - 10, y - 25, 'a', noAllow, game.redactorMetaText);
            }
        }
    }

    if (coordinate.level === 0) {
        let style = {font: "36px Arial", fill: "#bbfff1", align: "center"};
        coordinate.coordinateText.height = game.add.text(x - 50, y - 15, coordinate.level, style, game.redactorMetaText);
    }

    if (coordinate.level === 1) {
        let style = {font: "36px Arial", fill: "#35daff", align: "center"};
        coordinate.coordinateText.height = game.add.text(x - 50, y - 15, coordinate.level, style, game.redactorMetaText);
    }

    if (coordinate.level === 3) {
        let style = {font: "36px Arial", fill: "#68ff59", align: "center"};
        coordinate.coordinateText.height = game.add.text(x - 50, y - 15, coordinate.level, style, game.redactorMetaText);
    }

    if (coordinate.level === 4) {
        let style = {font: "36px Arial", fill: "#fff523", align: "center"};
        coordinate.coordinateText.height = game.add.text(x - 50, y - 15, coordinate.level, style, game.redactorMetaText);
    }

    if (coordinate.level === 5) {
        let style = {font: "36px Arial", fill: "#ff2821", align: "center"};
        coordinate.coordinateText.height = game.add.text(x - 50, y - 15, coordinate.level, style, game.redactorMetaText);
    }

    game.bmdTerrain.draw(floorSprite, x, y);
    floorSprite.destroy();
}