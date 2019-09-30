// метод снятия всех опционных ивентов со всех игровых обьектов
function UnselectAll() {
    UnselectResource();
    UnselectUnits();
}

function UnselectResource() {
    // ресурсы
    for (let x in game.map.reservoir) {
        for (let y in game.map.reservoir[x]) {
            game.map.reservoir[x][y].sprite.events.onInputDown.removeAll();
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

function UnselectUnits() {
    // юниты

    for (let i in selectUnits) {
        selectUnits[i].sprite.frame = 0;
    }

    selectUnits = [];
}