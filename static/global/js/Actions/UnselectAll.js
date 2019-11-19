// метод снятия всех опционных ивентов со всех игровых обьектов
function UnselectAll() {
    UnselectResource();
    UnselectUnits();
    UnselectDigger();
    UnselectAttack();

    dontMove = false;
    notOpen = false;
}

function UnselectResource() {
    // ресурсы
    for (let x in game.map.reservoir) {
        for (let y in game.map.reservoir[x]) {
            if (game.map.reservoir[x][y] && game.map.reservoir[x][y].sprite) {
                game.map.reservoir[x][y].sprite.events.onInputDown.removeAll();
            }
        }
    }

    for (let i in game.units) {
        let unit = game.units[i];
        if (unit.selectMiningLine) {
            unit.selectMiningLine.graphics.destroy();
            unit.selectMiningLine = null;
        }
    }
}

function UnselectDigger() {
    game.input.onDown.removeAll();
    for (let i in game.units) {
        let unit = game.units[i];
        if (unit.selectDiggerLine) {
            unit.selectDiggerLine.graphics.destroy();
            unit.selectDiggerLine = null;
        }
    }
}

function UnselectAttack() {
    document.getElementById("GameCanvas").style.cursor = "unset";
    game.bmdTerrain.sprite.events.onInputDown.removeAll();
    notOpen = false;

    if (game.targetCursorSprite) {
        game.targetCursorSprite.destroy();
        game.targetCursorSprite = null;
    }

    for (let i in game.objects) {
        if (game.objects[i] && game.objects[i].objectSprite) {
            game.objects[i].objectSprite.events.onInputDown.removeAll();
        }
    }

    for (let i in game.boxes) {
        if (game.boxes[i] && game.boxes[i].sprite) {
            game.boxes[i].sprite.events.onInputDown.removeAll();
        }
    }

    for (let i in game.units) {
        let unit = game.units[i];
        if (unit && unit.sprite && unit.sprite.unitBody) {
            unit.sprite.unitBody.events.onInputDown.removeAll();
        }
    }
}

function UnselectUnits() {
    // юниты

    for (let i in selectUnits) {
        selectUnits[i].sprite.frame = 0;
    }

    selectUnits = [];
}