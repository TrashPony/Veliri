function update() {
    //MoveUnit();
    RotateUnit(); // функция для повора юнитовский спрайтов
    GrabCamera(); // функцуия для перетаскивания карты мышкой /* Магия */
    AlphaSelect(); // анимация линий который обозначают зоны
    AnimateFog();

    game.floorObjectLayer.sort('y', Phaser.Group.SORT_ASCENDING);
}