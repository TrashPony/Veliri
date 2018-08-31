function CreateTerrain(coordinate, q, r) {






    floorSprite.inputEnabled = true; // включаем ивенты на спрайт
    floorSprite.events.onInputOut.add(TipOff, floorSprite);
    floorSprite.events.onInputDown.add(RemoveSelect);
    floorSprite.z = 0;

    coordinate.sprite = floorSprite;
    coordinate.fogSprite = fogSprite;


    let label = game.add.text(20, 15, q + "," + r);
    floorSprite.addChild(label);
}