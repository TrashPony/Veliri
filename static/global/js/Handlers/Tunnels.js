function OpenTunnel(jsonData) {
    if (!game || !game.map) return;

    let coordinatTunnel = game.map.OneLayerMap[jsonData.x][jsonData.y];
    coordinatTunnel.objectSprite.frame = 0;
}

function CloseTunnel(jsonData) {
    if (!game || !game.map) return;

    let coordinatTunnel = game.map.OneLayerMap[jsonData.x][jsonData.y];
    coordinatTunnel.objectSprite.frame = 2;
}