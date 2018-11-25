function update() {
    if (game && game.typeService === "battle") {
        UpdateRotateUnit(); // функция для повора юнитовский спрайтов
        MoveUnit();
        FlightBullet(); // ослеживает все летящие спрайты пуль
    }

    // Отдалить камеру
    //game.camera.scale.x -= 0.005;
    //game.camera.scale.y -= 0.005;
    // приблизить камеру
    //game.camera.scale.x += 0.005;
    //game.camera.scale.y += 0.005;

    GrabCamera(); // функцуия для перетаскивания карты мышкой /* Магия */
    game.floorObjectLayer.sort('y', Phaser.Group.SORT_DESCENDING);
}