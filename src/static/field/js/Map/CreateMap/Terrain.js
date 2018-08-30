function CreateTerrain(coordinate, q, r) {
    let hexagonWidth = 100;
    let hexagonHeight = 80;

    let hexagonX = hexagonWidth * q * 1.5 + (hexagonWidth / 4 * 3) * (r % 2);
    let hexagonY = hexagonHeight * r / 2;

    let floorSprite = game.floorLayer.create(hexagonX, hexagonY, "hexagon");
    let fogSprite = game.fogOfWar.create(hexagonX, hexagonY, 'FogOfWar');
    floorSprite.inputEnabled = true; // включаем ивенты на спрайт
    floorSprite.events.onInputOut.add(TipOff, floorSprite);
    floorSprite.events.onInputDown.add(RemoveSelect);
    floorSprite.z = 0;

    coordinate.sprite = floorSprite;
    coordinate.fogSprite = fogSprite;


    let label = game.add.text(20, 15, q + "," + r);
    floorSprite.addChild(label);
}