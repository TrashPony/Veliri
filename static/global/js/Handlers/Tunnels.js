function OpenTunnel(jsonData) {
    if (!game || !game.map) return;

    let coordinatTunnel = game.map.OneLayerMap[jsonData.q][jsonData.r];
    coordinatTunnel.objectSprite.frame = 0;
}

function CloseTunnel(jsonData) {
    if (!game || !game.map) return;

    let coordinatTunnel = game.map.OneLayerMap[jsonData.q][jsonData.r];
    coordinatTunnel.objectSprite.frame = 2;
}