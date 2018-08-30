function AnimateUseMapEquip(jsonMessage) {
    console.log(jsonMessage);

    let xUse = JSON.parse(jsonMessage).x_use;
    let yUse = JSON.parse(jsonMessage).y_use;
    let zone_effect = JSON.parse(jsonMessage).zone_effect;
    let equip = JSON.parse(jsonMessage).applied_equip; // id:equip

    let coordinateUse = game.map.OneLayerMap[xUse][yUse];

    for (let x in zone_effect) {
        if (zone_effect.hasOwnProperty(x)) {
            for (let y in zone_effect[x]) {
                if (zone_effect[x].hasOwnProperty(y)) {
                    let coordinate = game.map.OneLayerMap[x][y];
                    coordinate.effects = zone_effect.effects;
                }
            }
        }
    }

    if (equip.type === "small_bomb") {
        smallBombAnimate(coordinateUse);
    }
}

function smallBombAnimate(coordinateUse) {
    console.log(coordinateUse)
}