function GetXYCenterHex(q, r) {
    let x, y;

    let verticalOffset = game.hexagonHeight * 3 / 4;
    let horizontalOffset = game.hexagonWidth;

    if (r % 2 !== 0) {
        x = game.hexagonWidth + (horizontalOffset * q)
    } else {
        x = game.hexagonWidth / 2 + (horizontalOffset * q)
    }
    y = game.hexagonHeight / 2 + (r * verticalOffset);
    return {x: x, y: y}
}