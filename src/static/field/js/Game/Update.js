function update() {
    //MoveUnit();
    RotateUnit();
    GrabCamera(); // функцуия для перетаскивания карты мышкой /* Магия */

    game.floorObjectLayer.sort('y', Phaser.Group.SORT_ASCENDING);
}