let selectUnits = [];
let start = false; // что бы не генерить прямоуголник всегда

let selectRect = null;
let rectOption = {x: 0, y: 0, height: 0, width: 0};
let selectOneUnit = false;

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
        CheckSelectUnits();
    }
}

function StopSelectableUnits(pointer) {
    //https://ru.stackoverflow.com/questions/758529/%D0%9F%D0%B5%D1%80%D0%B5%D1%81%D0%B5%D1%87%D0%B5%D0%BD%D0%B8%D0%B5-%D0%B4%D0%B2%D1%83%D1%85-%D0%BF%D1%80%D1%8F%D0%BC%D0%BE%D1%83%D0%B3%D0%BE%D0%BB%D1%8C%D0%BD%D0%B8%D0%BA%D0%BE%D0%B2-c
    if (game.input.activePointer.leftButton.isDown) {
        start = false;
        if (pointer.duration > 200) {
            // все союзные юниты которые внутри выделяющего квадрата, или имеет с ним касание попадют в выделеных юнитов
            CheckSelectUnits();
        }
        if (selectRect) selectRect.destroy();
    }
}

function CheckBoxInBox(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2) {
    // поправки увеличеной камеры
    ax1 = ax1 / game.camera.scale.x;
    ay1 = ay1 / game.camera.scale.y;
    ax2 = ax2 / game.camera.scale.x;
    ay2 = ay2 / game.camera.scale.y;

    return ((ax1 < bx2 && ax2 > bx1) || (ax1 > bx2 && ax2 < bx1)) && ((ay1 < by2 && ay2 > by1) || (ay1 > by2 && ay2 < by1))
}

function SelectOneUnit(unit, boxSprite, setFlag) {
    if (setFlag) {
        selectOneUnit = true;
    }

    if (unit.owner_id !== game.user_id) {
        return
    }

    for (let i in selectUnits) {
        selectUnits[i].sprite.frame = 0;
    }

    selectUnits = [];
    selectUnits.push(unit);
    setTimeout(function () {
        unitInfo(unit, boxSprite);
    }, 10)
}

function CheckSelectUnits() {
    // каждое выделение снимает выделение с других, но если юнит уже был выделен то снимать не надо
    selectUnits = [];

    for (let i in game.units) {
        if (game.units.hasOwnProperty(i)) {
            if (CheckBoxInBox(
                rectOption.x,
                rectOption.y,
                rectOption.x + -1 * rectOption.width,
                rectOption.y + -1 * rectOption.height,

                game.units[i].sprite.x - (game.units[i].sprite.width / 8),
                game.units[i].sprite.y - (game.units[i].sprite.height / 8),
                game.units[i].sprite.x + (game.units[i].sprite.width / 8),
                game.units[i].sprite.y + (game.units[i].sprite.height / 8),
            )) {

                if (game.units[i].owner_id === game.user_id) {
                    unitInfo(game.units[i], game.units[i].sprite, game.units[i].owner_id);
                    selectUnits.push(game.units[i])
                }
            } else {
                unitRemoveInfo(game.units[i], game.units[i].sprite)
            }
        }
    }
}

function GetSelectUnitByID(unitID) {
    for (let i in selectUnits) {
        if (selectUnits[i].id === unitID) {
            return selectUnits[i]
        }
    }
}

function UnSelectUnit() {
    if (game.input.activePointer.rightButton.isDown && game.input.activePointer.rightButton.duration <= 100) {
        UnselectAll();
    }
}

function initMove(pointer) {
    if (game.input.activePointer.leftButton.isDown && pointer.duration <= 200) {


        if (!selectOneUnit) {

            // todo если игрок передумал и тыкнул в другое место то не преследовать ящик
            // if (game.units[selectUnits[i].id] && game.units[selectUnits[i].id].toBox) {
            //     game.units[selectUnits[i].id].toBox.to = false
            // }


            global.send(JSON.stringify({
                event: "MoveTo",
                to_x: (game.input.mousePointer.x + game.camera.x) / game.camera.scale.x,
                to_y: (game.input.mousePointer.y + game.camera.y) / game.camera.scale.y,
                units_id: getIDsSelectUnits(),
            }));
        } else {
            selectOneUnit = false;
        }
    }
}

function getIDsSelectUnits() {
    let unitsID = [];
    for (let i in selectUnits) {
        unitsID.push(selectUnits[i].id);
    }

    return unitsID;
}

function StopUnit(id) {
    global.send(JSON.stringify({
        event: "StopMove",
        units_id: [Number(id)],
    }));
}