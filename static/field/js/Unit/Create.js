function CreateLocalUnit(unit, inVisible) {
    let xy = GetXYCenterHex(unit.q, unit.r);

    unit.user_id = unit.owner_id;
    unit.user_name = unit.owner;

    if (game.user.name === unit.owner) {
        unit = CreateUnit(unit, xy.x, xy.y, unit.rotate, unit.body_color_1, unit.body_color_2, unit.weapon_color_1, unit.weapon_color_2, unit.owner_id, 'MySelectUnit', true);
        if (unit.body.mother_ship) {
            game.camera.focusOnXY(xy.x * game.camera.scale.x, xy.y * game.camera.scale.y);
        }
    } else {
        unit = CreateUnit(unit, xy.x, xy.y, unit.rotate, unit.body_color_1, unit.body_color_2, unit.weapon_color_1, unit.weapon_color_2, unit.owner_id, 'HostileSelectUnit', true);
    }

    unit.sprite.unitBody.events.onInputDown.add(function () {
        SelectUnit(unit, false)
    });
    unit.sprite.unitBody.events.onInputOver.add(UnitMouseOver, unit); // обрабатываем наведение мышки
    unit.sprite.unitBody.events.onInputOut.add(UnitMouseOut, unit);   // TODO onInputOut работает плохо везде, его надо чемто заменить обрабатываем убирание мышки

    if (unit.effects !== null && unit.effects.length > 0) {
        CreateAnimateEffects(unit)
    }

    if (unit.action && game.user.name === unit.owner) {
        DeactivationUnit(unit);
    }

    if (inVisible) {
        unit.sprite.alpha = 0;
        unit.sprite.heal.alpha = 0;
    }

    addToGameUnit(unit);
    if (unit.reload) {
        ReloadMark({q: unit.q, r: unit.r})
    }

    return unit
}

function CreateAnimateEffects(unit) {
    for (let i in unit.effects) {
        if (unit.effects.hasOwnProperty(i) && unit.effects[i] != null && unit.effects[i].type === "unit_always_animate") {
            if (unit.effects[i].name === "animate_energy_shield") {
                energyShieldAnimate(unit);
            }
        }
    }
}