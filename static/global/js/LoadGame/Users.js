function CreateUnits(units) {
    for (let i in units) {
        if (units.hasOwnProperty(i)) {
            game.units[i] = CreateUnit(
                units[i],
                units[i].x,
                units[i].y,
                units[i].rotate,
                units[i].body_color_1,
                units[i].body_color_2,
                units[i].weapon_color_1,
                units[i].weapon_color_2,
                units[i].owner_id,
                'MySelectUnit',
                false,
            );

            if (game.units[i].owner_id === game.user_id && game.units[i].body.mother_ship) {
                FocusUnit(game.units[i].id);
            }
        }
    }
}

function CreateNewUnit(newUnit) {
    if (!game || !game.units || !newUnit || !newUnit.id) return;

    let unit = game.units[newUnit.id];
    if (!unit || !unit.sprite) {
        game.units[newUnit.id] = CreateUnit(
            newUnit,
            newUnit.x,
            newUnit.y,
            newUnit.rotate,
            newUnit.body_color_1,
            newUnit.body_color_2,
            newUnit.weapon_color_1,
            newUnit.weapon_color_2,
            newUnit.owner_id,
            'MySelectUnit',
            false,
        );
    }
}