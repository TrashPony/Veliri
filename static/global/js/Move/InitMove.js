let selectUnits = [];
let start = false; // что бы не генерить прямоуголник всегда

let selectRect = null;
let rectOption = {x: 0, y: 0, height: 0, width: 0};

function StartSelectableUnits() {
    // создаем прямоуголник который тянется от точки куда тыкнул юзер изначально до текущего положения курсора
    if (game.input.activePointer.leftButton.isDown && !start) {
        start = true;
        rectOption.x = game.input.mousePointer.x + game.camera.x;
        rectOption.y = game.input.mousePointer.y + game.camera.y;
    }

    if (start && game.input.activePointer.leftButton.duration > 200) {
        if (selectRect) selectRect.destroy();

        rectOption.height = rectOption.y - (game.input.mousePointer.y + game.camera.y);
        rectOption.width = rectOption.x - (game.input.mousePointer.x + game.camera.x);

        selectRect = game.add.graphics(rectOption.x / game.camera.scale.x, rectOption.y / game.camera.scale.y);
        selectRect.lineStyle(1, 0xcce3ff, 0.8);
        selectRect.beginFill(0xcce3ff, 0.1);
        selectRect.drawRect(
            0,
            0,
            (rectOption.width / game.camera.scale.x) * -1,
            (rectOption.height / game.camera.scale.y) * -1,
        );
        selectRect.endFill();
    }
}

function StopSelectableUnits(pointer) {
    if (game.input.activePointer.leftButton.isDown) {
        start = false;
        if (pointer.duration > 200) {
            // все союзные юниты которые внутри выделяющего квадрата, или имеет с ним касание попадют в выделеных юнитов
        }
        if (selectRect) selectRect.destroy();
    }
}

function initMove(e) {
    if (game.input.activePointer.leftButton.isDown) {

        if (game.squad.toBox) {
            game.squad.toBox.to = false
        }

        global.send(JSON.stringify({
            event: "MoveTo",
            to_x: e.worldX / game.camera.scale.x,
            to_y: e.worldY / game.camera.scale.y,
            //units: selectUnits,
        }));
    }
}