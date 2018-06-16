function update() {
    UpdateRotateUnit(); // функция для повора юнитовский спрайтов
    GrabCamera(); // функцуия для перетаскивания карты мышкой /* Магия */
    UnitDestroy(); // создание уберажие тумана войны
    MoveUnit();

    game.floorObjectLayer.sort('y', Phaser.Group.SORT_ASCENDING);
}