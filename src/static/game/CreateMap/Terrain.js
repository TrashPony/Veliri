function CreateTerrain(coordinate, x, y, q, r) {

    let floorSprite = game.floorLayer.create(x, y, coordinate.texture_flore);
    floorSprite.scale.set(0.5);

    floorSprite.inputEnabled = true; // включаем ивенты на спрайт
    floorSprite.input.pixelPerfectOver = true;   // уберает ивенты наведения на пустую зону спрайта
    floorSprite.input.pixelPerfectClick = true;  // уберает ивенты кликов на пустую зону спрайта

    floorSprite.events.onInputOut.add(TipOff, floorSprite);
    floorSprite.z = 0;
    coordinate.sprite = floorSprite;

    if (game && game.typeService === "battle") {
        let fogSprite = game.fogOfWar.create(x, y, 'FogOfWar');
        fogSprite.scale.set(0.5);
        floorSprite.events.onInputDown.add(RemoveSelect);
        coordinate.fogSprite = fogSprite;
        coordinate.fogSprite = fogSprite;
    }

    if (game && game.typeService !== "battle") {
        let style = {font: "12px Arial", fill: "#606060", align: "center"};
        game.add.text(x + 50, y + 20, q + "," + r, style, game.weaponEffectsLayer);

        let allow = {font: "12px Arial", fill: "#150bff", align: "center"};
        let noAllow = {font: "12px Arial", fill: "#ff2821", align: "center"};

        if (!(coordinate.move && coordinate.view && coordinate.attack)){
            if (coordinate.move) {
                game.add.text(x + 25, y + 45, 'm', allow, game.weaponEffectsLayer);
            } else {
                game.add.text(x + 25, y + 45, 'm', noAllow, game.weaponEffectsLayer);
            }

            if (coordinate.view) {
                game.add.text(x + 40, y + 45, 'w', allow, game.weaponEffectsLayer);
            } else {
                game.add.text(x + 40, y + 45, 'w', noAllow, game.weaponEffectsLayer);
            }

            if (coordinate.attack) {
                game.add.text(x + 55, y + 45, 'a', allow, game.weaponEffectsLayer);
            } else {
                game.add.text(x + 55, y + 45, 'a', noAllow, game.weaponEffectsLayer);
            }
        }
    }

    if (coordinate.level === 3) {
        let style = {font: "36px Arial", fill: "#68ff59", align: "center"};
        let label = game.add.text(50, 55, coordinate.level, style);
        floorSprite.addChild(label);
    }

    if (coordinate.level === 4) {
        let style = {font: "36px Arial", fill: "#fff523", align: "center"};
        let label = game.add.text(50, 55, coordinate.level, style);
        floorSprite.addChild(label);
    }

    if (coordinate.level === 5) {
        let style = {font: "36px Arial", fill: "#ff2821", align: "center"};
        let label = game.add.text(50, 55, coordinate.level, style);
        floorSprite.addChild(label);
    }
}