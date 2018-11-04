function CreateTerrain(coordinate, x, y, q, r) {
    let floorSprite = game.floorLayer.create(x, y, "hexagon");
    coordinate.sprite = floorSprite;

    //let label = game.add.text(20, 15, q + "," + r);
    //floorSprite.addChild(label);
    coordinate.sprite = floorSprite;

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