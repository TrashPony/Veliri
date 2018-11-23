function CreateTerrain(coordinate, x, y, q, r) {

    let floorSprite = game.floorLayer.create(x, y, "hexagon");
    floorSprite.scale.set(0.5);
    floorSprite.inputEnabled = true; // включаем ивенты на спрайт
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

    //let label = game.add.text(20, 15, q + "," + r);
    //floorSprite.addChild(label);


    if (coordinate.level === 3) {
        let style = { font: "24px Arial", fill: "#ffa92b", align: "center" };
        let label = game.add.text(20, 50, coordinate.level, style);
        floorSprite.addChild(label);
    }

    if (coordinate.level === 4) {
        let style = { font: "24px Arial", fill: "#ff3f41", align: "center" };
        let label = game.add.text(20, 50, coordinate.level, style);
        floorSprite.addChild(label);
    }
}