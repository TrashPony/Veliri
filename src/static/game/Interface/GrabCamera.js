function GrabCamera() {
    if (game.input.activePointer.rightButton.isDown) { // ловит нажатие правой кнопки маши в игре

        game.camera.target = null;

        if (game && game.typeService === "battle") {
            RemoveSelect();
        } else if (game && game.typeService === "global" || game.typeService === "mapEditor") {
            CreateMiniMap(game.map)
        }

        if (game.origDragPoint) {
            game.camera.x += game.origDragPoint.x - game.input.activePointer.position.x; // перемещать камеру по сумме, перемещенную мышью с момента последнего обновления
            game.camera.y += game.origDragPoint.y - game.input.activePointer.position.y;
        }
        game.origDragPoint = game.input.activePointer.position.clone(); // установите новое начало перетаскивания в текущую позицию
    } else {
        game.origDragPoint = null;
    }
}