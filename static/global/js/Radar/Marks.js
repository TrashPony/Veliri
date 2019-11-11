function CreateMark(mark, x, y) {

    if (!game || !game.radar_marks) return;

    let oldMark = game.radar_marks[mark.uuid];
    if (oldMark) {
        RemoveMark(mark)
    }

    mark.sprite = game.flyObjectsLayer.create(x, y, mark.type + "_radar_icon");
    mark.sprite.anchor.setTo(0.5, 0.5);
    mark.sprite.scale.setTo(0.2);

    game.radar_marks[mark.uuid] = mark;

    return mark
}

function RemoveMark(mark) {
    if (!game || !game.radar_marks) return;
    mark = game.radar_marks[mark.uuid];
    if (mark) {
        mark.sprite.destroy();
        game.radar_marks[mark.uuid] = null;
    }
}

function HideMark(mark) {
    if (!game || !game.radar_marks) return;
    mark = game.radar_marks[mark.uuid];
    if (mark) {
        mark.sprite.visible = false;
    }
}

function UnhideMark(mark, x, y) {
    if (!game || !game.radar_marks) return;
    let oldMark = game.radar_marks[mark.uuid];
    if (oldMark) {
        oldMark.sprite.x = x;
        oldMark.sprite.y = y;
        oldMark.sprite.visible = true;
    } else {
        CreateMark(mark, x, y)
    }
}

function MoveMark(data) {
    if (!game || !game.radar_marks) return;

    let mark = game.radar_marks[data.radar_mark.uuid];
    if (!mark) {
        mark = CreateMark(data.radar_mark, data.path_unit.x, data.path_unit.y)
    }

    if (!mark.sprite.visible) {
        UnhideMark(mark, data.path_unit.x, data.path_unit.y);
        return;
    }

    game.add.tween(mark.sprite).to({
            x: data.path_unit.x,
            y: data.path_unit.y
        }, data.path_unit.millisecond, Phaser.Easing.Linear.None, true, 0
    );
    CreateMiniMap();
}