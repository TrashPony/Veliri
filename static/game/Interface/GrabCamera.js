function GrabCamera() {
    if (game.input.activePointer.rightButton.isDown && game.input.activePointer.rightButton.duration > 100) { // ловит нажатие правой кнопки маши в игре
        if (game.origDragPoint) {
            game.camera.x += game.origDragPoint.x - game.input.activePointer.position.x; // перемещать камеру по сумме, перемещенную мышью с момента последнего обновления
            game.camera.y += game.origDragPoint.y - game.input.activePointer.position.y;

            if (game && game.typeService === "battle") {
                RemoveSelect();
            }

            CreateMiniMap();
            game.camera.target = null;
        }
        game.origDragPoint = game.input.activePointer.position.clone(); // установите новое начало перетаскивания в текущую позицию
    } else {
        game.origDragPoint = null;
    }
}