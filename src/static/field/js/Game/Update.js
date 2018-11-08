function update() {
    UpdateRotateUnit(); // функция для повора юнитовский спрайтов
    GrabCamera(); // функцуия для перетаскивания карты мышкой /* Магия */
    MoveUnit();
    FlightBullet(); // ослеживает все летящие спрайты пуль

    game.floorObjectLayer.sort('y', Phaser.Group.SORT_DESCENDING);
}