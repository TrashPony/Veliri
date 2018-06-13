function MarkEquipSelect(markCode, clickFunction, myUnit, hostileUnit, inMap, equip) {
    if (inMap) {
        //todo эквипы которые работаю на територии
    }

    if (!inMap) {
        for (var x in game.units) {
            if (game.units.hasOwnProperty(x)) {
                for (var y in game.units[x]) {
                    if (game.units[x].hasOwnProperty(y) && game.units[x][y].sprite) {

                        var unit = game.units[x][y];

                        if (myUnit && hostileUnit) {
                            unit.sprite.frame = markCode;
                            unit.sprite.events.onInputDown.add(function () {
                                clickFunction(unit, equip)
                            });
                            unit.sprite.input.priorityID = 1; // утсанавливает повышеный приоритет среди спрайтов на которых мышь
                        } else {
                            if (myUnit && game.user.name === unit.owner) {
                                unit.sprite.frame = markCode;
                                unit.sprite.events.onInputDown.add(function () {
                                    clickFunction(unit, equip)
                                });
                                unit.sprite.input.priorityID = 1;
                            }

                            if (hostileUnit && game.user.name !== unit.owner) {
                                unit.sprite.frame = markCode;
                                unit.sprite.events.onInputDown.add(function () {
                                    clickFunction(unit, equip)
                                });
                                unit.sprite.input.priorityID = 1;
                            }
                        }
                    }
                }
            }
        }
    }
}