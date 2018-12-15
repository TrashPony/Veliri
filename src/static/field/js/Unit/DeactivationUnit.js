function DeactivationUnit(unit) {
    //todo сделать так что бы юзеру было понятно что этот юнит уже ходил
    unit.body.tint = 0x757575;
}

function ActivationUnit(unit) {
    unit.body.tint = 0xFFFFFF;
}

function removeUnitInput() { // удаляем ивенты при наведение
    for (let x in game.units) {
        if (game.units.hasOwnProperty(x)) {
            for (let y in game.units[x]) {
                if (game.units[x].hasOwnProperty(y)) {
                    game.units[x][y].sprite.unitBody.events.onInputOver.removeAll();
                    game.units[x][y].sprite.unitBody.events.onInputOut.removeAll();
                }
            }
        }
    }
}

function activateUnitInput() {
    for (let x in game.units) {
        if (game.units.hasOwnProperty(x)) {
            for (let y in game.units[x]) {
                if (game.units[x].hasOwnProperty(y)) {
                    if (game.units[x][y].sprite) {
                        game.units[x][y].sprite.unitBody.events.onInputOver.add(UnitMouseOver, game.units[x][y]); // обрабатываем наведение мышки
                        game.units[x][y].sprite.unitBody.events.onInputOut.add(UnitMouseOut, game.units[x][y]);   // обрабатываем убирание мышки
                    }
                }
            }
        }
    }
}