function AnimateUseMapEquip(jsonMessage) {

    let qUse = JSON.parse(jsonMessage).q_use;
    let rUse = JSON.parse(jsonMessage).r_use;
    let zone_effect = JSON.parse(jsonMessage).zone_effect;
    let equip = JSON.parse(jsonMessage).applied_equip; // id:equip
    let useUnit = GetGameUnitID(JSON.parse(jsonMessage).use_unit.id);

    let coordinateUse = game.map.OneLayerMap[qUse][rUse];

    UpdateUnit(useUnit);

    for (let q in zone_effect) {
        if (zone_effect.hasOwnProperty(q)) {
            for (let r in zone_effect[q]) {
                if (zone_effect[q].hasOwnProperty(r)) {
                    let coordinate = game.map.OneLayerMap[q][r];
                    coordinate.effects = zone_effect[q][r].effects;
                }
            }
        }
    }

    if (equip.name === "small_bomb") {
        smallBombAnimate(coordinateUse);
    }
}

function smallBombAnimate(coordinateUse) {
    console.log(coordinateUse)
}