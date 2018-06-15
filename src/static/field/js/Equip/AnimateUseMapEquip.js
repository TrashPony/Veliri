function AnimateUseMapEquip(jsonMessage) {
    console.log(jsonMessage);

    var equipBox = document.getElementById(JSON.parse(jsonMessage).applied_equip.id + ":equip"); // id:equip
    RemoveEquipCell(equipBox);

    var xUse = JSON.parse(jsonMessage).x_use;
    var yUse = JSON.parse(jsonMessage).y_use;
    var zone_effect = JSON.parse(jsonMessage).zone_effect;
    var equip = JSON.parse(jsonMessage).applied_equip; // id:equip

    var coordinateUse = game.map.OneLayerMap[xUse][yUse];

    for (var x in zone_effect) {
        if (zone_effect.hasOwnProperty(x)) {
            for (var y in zone_effect[x]) {
                if (zone_effect[x].hasOwnProperty(y)) {
                    var coordinate = game.map.OneLayerMap[x][y];
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